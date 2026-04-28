package main

import (
	"testing"
)

func BenchmarkAllocateViaPoolSyncNoSleep100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocateViaPoolSyncNoSleep(100)
	}
}

func BenchmarkAllocateViaPoolSyncNoSleep1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocateViaPoolSyncNoSleep(1000)
	}
}

func BenchmarkAllocateViaPoolSyncNoSleep50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocateViaPoolSyncNoSleep(50000)
	}
}

func BenchmarkAllocateViaPoolSyncWithSleep100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocateViaPoolSyncWithSleep(100)
	}
}

func BenchmarkAllocateViaPoolSyncWithSleep1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocateViaPoolSyncWithSleep(1000)
	}
}

func BenchmarkAllocateViaPoolSyncWithSleep50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocateViaPoolSyncWithSleep(50000)
	}
}
