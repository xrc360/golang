package gmlock_test

import (
	"testing"

	"github.com/xrcn/cg/os/gmlock"
)

var (
	lockKey = "This is the lock key for gmlock."
)

func Benchmark_GMLock_Lock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gmlock.Lock(lockKey)
		gmlock.Unlock(lockKey)
	}
}

func Benchmark_GMLock_RLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gmlock.RLock(lockKey)
		gmlock.RUnlock(lockKey)
	}
}

func Benchmark_GMLock_TryLock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gmlock.TryLock(lockKey) {
			gmlock.Unlock(lockKey)
		}
	}
}

func Benchmark_GMLock_TryRLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gmlock.TryRLock(lockKey) {
			gmlock.RUnlock(lockKey)
		}
	}
}
