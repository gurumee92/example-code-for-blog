package main

import (
	"testing"
)

func TestWriteMapMutex(t *testing.T) {
	m := map[int]int{}
	num := 100
	WriteMapMutex(m, num)

	for i := 0; i < num; i++ {
		if m[i] != i {
			t.Errorf("m[%d] != %d, want %d\n", i, m[i], i)
		}
	}
}

func TestWriteMapChan(t *testing.T) {
	m := map[int]int{}
	num := 100
	WriteMapChan(m, num)

	for i := 0; i < num; i++ {
		if m[i] != i {
			t.Errorf("m[%d] != %d, want %d\n", i, m[i], i)
		}
	}
}
