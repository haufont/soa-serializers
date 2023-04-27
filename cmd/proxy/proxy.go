package main

import (
	"flag"
	"fmt"
	"net"
	"serialize-benchmarker/internal/env"
	"serialize-benchmarker/internal/format"
	"serialize-benchmarker/internal/request"
	"serialize-benchmarker/internal/udp"
	"strings"
	"sync"
)

type Ports struct {
	Port            int
	NativePort      int
	XMLPort         int
	JSONPort        int
	ProtoBufPort    int
	AvroPort        int
	YAMLPort        int
	MessagePackPort int
}

func (p Ports) String() string {
	return strings.Join([]string{
		fmt.Sprintf("Port: %d", p.Port),
		fmt.Sprintf("NativePort: %d", p.NativePort),
		fmt.Sprintf("XMLPort: %d", p.XMLPort),
		fmt.Sprintf("JSONPort: %d", p.JSONPort),
		fmt.Sprintf("ProtoBufPort: %d", p.ProtoBufPort),
		fmt.Sprintf("AvroPort: %d", p.AvroPort),
		fmt.Sprintf("YAMLPort: %d", p.YAMLPort),
		fmt.Sprintf("MessagePackPort: %d", p.MessagePackPort),
	}, ", ")
}

func FlagPort(p *int, portName string, def int) {
	flag.IntVar(p, portName, env.GetIntEnv(strings.ToUpper(portName), def), "port")
}

func ParseArgs() (ports Ports, maddr string) {
	FlagPort(&ports.Port, "port", 2000)
	FlagPort(&ports.NativePort, "native_port", 2001)
	FlagPort(&ports.XMLPort, "xml_port", 2002)
	FlagPort(&ports.JSONPort, "json_port", 2003)
	FlagPort(&ports.ProtoBufPort, "protobuf_port", 2004)
	FlagPort(&ports.AvroPort, "avro_port", 2005)
	FlagPort(&ports.YAMLPort, "yaml_port", 2006)
	FlagPort(&ports.MessagePackPort, "message_pack_port", 2007)
	flag.StringVar(&maddr, "maddr", env.GetStringEnv("MADDR", ""), "multicast address")
	flag.Parse()
	return
}

type Clients map[format.Format]*net.UDPConn

func SetClient(clients *Clients, mclients *Clients, f format.Format, port int, maddr string, serverPrefix string) {
	if serverPrefix == "" {
		serverPrefix = f.String()
	}
	addr := fmt.Sprintf("%s-serialize-benchmarker", serverPrefix)
	(*clients)[f] = udp.InitUDPClient(addr, port)
	if maddr != "" {
		(*mclients)[f] = udp.InitUDPClient(maddr, port+10)
	}
}

func InitClients(ports Ports, maddr string) (Clients, Clients) {
	clients := make(Clients)
	mclients := make(Clients)
	SetClient(&clients, &mclients, format.Native, ports.NativePort, maddr, "")
	SetClient(&clients, &mclients, format.XML, ports.XMLPort, maddr, "")
	SetClient(&clients, &mclients, format.JSON, ports.JSONPort, maddr, "")
	SetClient(&clients, &mclients, format.ProtoBuf, ports.ProtoBufPort, maddr, "")
	SetClient(&clients, &mclients, format.ApacheAvro, ports.AvroPort, maddr, "")
	SetClient(&clients, &mclients, format.YAML, ports.YAMLPort, maddr, "")
	SetClient(&clients, &mclients, format.MessagePack, ports.MessagePackPort, maddr, "message-pack")
	return clients, mclients
}

type Handler struct {
	clients  Clients
	mclients Clients
}

func (h Handler) HandleClient(f format.Format, client *net.UDPConn) string {
	sz, err := client.Write([]byte(request.GetResult))
	if err != nil || sz != len(request.GetResult) {
		return fmt.Sprintf("Unavailable %s-serialize-benchmarker", f.String())
	}
	buf := make([]byte, udp.BufferSize)
	sz, err = client.Read(buf)
	if err != nil || sz == 0 {
		return fmt.Sprintf("Unavailable %s-serialize-benchmarker", f.String())
	}
	return string(buf[:sz])
}

func (h Handler) HandleOne(f format.Format) string {
	return h.HandleClient(f, h.clients[f])
}

func (h Handler) HandleAll(clients Clients) string {
	if len(clients) == 0 {
		return "get_result all isn't available because maddr empty"
	}
	responses := make([]string, 0)
	responsesChan := make(chan string)
	for f, client := range clients {
		go func(f format.Format, client *net.UDPConn) {
			responsesChan <- h.HandleClient(f, client)
		}(f, client)
	}
	for range clients {
		responses = append(responses, <-responsesChan)
	}
	return strings.Join(responses, "")
}

func (h Handler) Handle(req string) string {
	f, all, err := request.ParseRequestWithFormatOrAll(req)
	if err != nil {
		return err.Error()
	}

	if !all {
		return h.HandleOne(f)
	} else {
		return h.HandleAll(h.clients)
	}
}

func StartUDPServer(wg *sync.WaitGroup, ports Ports, maddr string) {
	wg.Add(1)
	go func() {
		server := udp.InitUDPServer(ports.Port)
		defer server.Close()
		clients, mclients := InitClients(ports, maddr)
		udp.Serve(server, Handler{clients, mclients})
		wg.Done()
	}()
}

func main() {
	ports, maddr := ParseArgs()
	fmt.Printf("Start (%s, multicast addr: \"%s\")", ports.String(), maddr)
	var serversWaitGroup sync.WaitGroup
	StartUDPServer(&serversWaitGroup, ports, maddr)
	serversWaitGroup.Wait()
}
