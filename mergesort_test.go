package mergesort

import (
    "testing"
    "math/rand"
)

func TestBasicSort(t *testing.T){
    c := 10000
    f := make([]int, c)
    for i := 0; i < c; i++ {
        f[i] = rand.Int()
    }

    r := Sort(f)

    if len(r) != len(f) {
        t.Fatalf("array length changed")
    }

    for i := 1; i < c; i++ {
        if r[i] < r[i-1] {
            t.Fatalf("not sorted %v", f)
        }
    }
}