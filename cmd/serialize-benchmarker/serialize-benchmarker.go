package main

import (
	"flag"
	"fmt"
	"serialize-benchmarker/internal/env"
	"serialize-benchmarker/internal/format"
	"serialize-benchmarker/internal/request"
	"serialize-benchmarker/internal/scorer"
	"serialize-benchmarker/internal/udp"
	"serialize-benchmarker/schema"
	"sync"
	"time"
)

func GetScorer(f format.Format, iter int) scorer.Scorer {
	testType := schema.InitTestType()
	if f != format.ProtoBuf {
		return scorer.NewScorer(f, testType, iter)
	} else {
		protoTestType := schema.TestTypeToProto(&testType)
		return scorer.NewProtoBufScorer(protoTestType, iter)
	}
}

func FormResponse(f format.Format, size int, st time.Duration, dt time.Duration) string {
	return fmt.Sprintf("%s-%d-%s-%s\n", f.Name(), size, st, dt)
}

func TestFormat(f format.Format, iter int) (string, error) {
	s := GetScorer(f, iter)
	size, st, dt := s.Score()
	return FormResponse(f, size, st, dt), nil
}

type Handler struct {
	f    format.Format
	iter int
}

func (h Handler) Handle(req string) string {
	if req != request.GetResult {
		return request.UnknownSimpleRequest(req)
	}
	res, _ := TestFormat(h.f, h.iter)
	return res
}

func ParseArgs() (port int, iter int, f format.Format, maddr string) {
	var formatStr string
	flag.StringVar(&formatStr, "format", env.GetStringEnv("FORMAT", ""), format.AllEnumToString())
	flag.IntVar(&port, "port", env.GetIntEnv("PORT", 2000), "port")
	flag.IntVar(&iter, "iter", 1000, "an int")
	flag.StringVar(&maddr, "maddr", env.GetStringEnv("MADDR", ""), "multicast address")
	flag.Parse()
	f = format.ParseFormatOrPanic(formatStr)
	return
}

func StartUDPServer(wg *sync.WaitGroup, f format.Format, iter int, port int) {
	wg.Add(1)
	go func() {
		udpServer := udp.InitUDPServer(port)
		defer udpServer.Close()
		udp.Serve(udpServer, Handler{f, iter})
		wg.Done()
	}()
}

func StartUDPMulticastServer(wg *sync.WaitGroup, f format.Format, iter int, maddr string, port int) {
	wg.Add(1)
	go func() {
		udpServer := udp.InitUDPMulticastServer(maddr, port)
		defer udpServer.Close()
		udp.Serve(udpServer, Handler{f, iter})
		wg.Done()
	}()
}

func main() {
	port, iter, f, maddr := ParseArgs()
	fmt.Printf("Start (Port: %d, Format: %s, Iter: %d, multicast addr: \"%s\")\n", port, f.String(), iter, maddr)

	var serversWaitGroup sync.WaitGroup
	StartUDPServer(&serversWaitGroup, f, iter, port)
	if maddr != "" {
		StartUDPMulticastServer(&serversWaitGroup, f, iter, maddr, port+10)
	}
	serversWaitGroup.Wait()
}
