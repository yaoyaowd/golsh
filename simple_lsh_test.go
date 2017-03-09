package golsh

import (
	"math/rand"
	"testing"
	// "fmt"
)

const (
	TEST_INSTANCES = 100
	TEST_RESULTS = 100
	ROWS = 10000
	VECLEN = 120
	TABLESIZE = 10
	KEYSIZE = 10
	SLOT = 5.0
)

func Test_NewSimpleLSHFromData(t *testing.T) {
	data := NewRandomData(ROWS, VECLEN)
	params := NewParams(VECLEN, TABLESIZE, KEYSIZE, SLOT)
	simpleLsh:= NewSimpleLSHFromData(params, data)

	for _, table := range simpleLsh.Tables {
		if len(table) == 0 {
			t.Error("new simple lsh from data failed!")
		}
	}
}

func Test_Query(t *testing.T) {
	data := NewRandomData(ROWS, VECLEN)
	params := NewParams(VECLEN, TABLESIZE, KEYSIZE, SLOT)
	simpleLsh:= NewSimpleLSHFromData(params, data)

	for _, table := range simpleLsh.Tables {
		if len(table) == 0 {
			t.Error("new simple lsh from data failed!")
		}
	}

	idx := rand.Perm(ROWS)[:TEST_INSTANCES]
	validIdx := rand.Perm(ROWS)[:TEST_RESULTS]
	for _, id := range idx {
		out := make(chan int)
		defer close(out)

		p := data.Points[id]
		avgDist := 0.0
		count := 0
		simpleLsh.Query(p, TEST_RESULTS, out)
		for {
			key := <- out
			if key == -1 {
				break
			}
			avgDist += float64(p.L2(data.Points[key]))
			count += 1
		}
		if count == 0 {
			t.Error("could not find any neighbours")
		}
		avgDist /= float64(count)

		validAvgDist := 0.0
		for _, key := range validIdx {
			validAvgDist += float64(p.L2(data.Points[key]))
		}
		validAvgDist /= TEST_RESULTS
		if avgDist > validAvgDist {
			t.Error("neighbbours are further than random")
		} else {
			// fmt.Printf("num found %d, avg dist %f, valid dist %f\n", count, avgDist, validAvgDist)
		}
	}
}
