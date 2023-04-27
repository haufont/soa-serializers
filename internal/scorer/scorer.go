package scorer

import (
	"serialize-benchmarker/internal/format"
	"serialize-benchmarker/internal/serializer"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Scorer interface {
	Score() (size int, st time.Duration, dt time.Duration)
}

type ImplScorer[TestType any] struct {
	testType TestType
	s        serializer.Serializer[TestType]
	iter     int
}

func NewScorer[TestType any](f format.Format, testType TestType, iter int) Scorer {
	if f == format.ProtoBuf {
		panic("UseNewProtoBufScorer")
	}
	return &ImplScorer[TestType]{testType, serializer.NewSerializer[TestType](f), iter}
}

func NewProtoBufScorer[TestType protoreflect.ProtoMessage](testType TestType, iter int) Scorer {
	return &ImplScorer[TestType]{testType, serializer.NewProtoBufSerializer[TestType](), iter}
}

func (scorer *ImplScorer[TestType]) Score() (size int, st time.Duration, dt time.Duration) {
	bytes, err := scorer.s.Serialize(&scorer.testType)
	if err != nil {
		panic(err)
	}
	size = len(bytes)
	v_copy := scorer.testType
	err = scorer.s.Deserialize(bytes, &v_copy)
	if err != nil {
		panic(err)
	}

	st_start := time.Now()
	for i := 0; i < scorer.iter; i++ {
		scorer.s.Serialize(&scorer.testType)
	}
	st = time.Since(st_start) / time.Duration(scorer.iter)

	dt_start := time.Now()
	for i := 0; i < scorer.iter; i++ {
		scorer.s.Deserialize(bytes, &v_copy)
	}
	dt = time.Since(dt_start) / time.Duration(scorer.iter)
	return
}
