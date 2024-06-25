// go test *.go -bench=".*" -benchmem

package gqueue_test

import (
	"testing"

	"github.com/xrc360/golang/container/gqueue"
)

var bn = 20000000

var length = 1000000

var qstatic = gqueue.New(length)

var qdynamic = gqueue.New()

var cany = make(chan interface{}, length)

func Benchmark_Gqueue_StaticPushAndPop(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		qstatic.Push(i)
		qstatic.Pop()
	}
}

func Benchmark_Gqueue_DynamicPush(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		qdynamic.Push(i)
	}
}

func Benchmark_Gqueue_DynamicPop(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		qdynamic.Pop()
	}
}

func Benchmark_Channel_PushAndPop(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		cany <- i
		<-cany
	}
}
