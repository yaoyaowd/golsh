package golsh

import (
	"math/rand"
	"testing"
	"fmt"
)

func Test_Query_FOREST(t *testing.T) {
	data := NewRandomData(ROWS, VECLEN)
	params := NewParams(VECLEN, TABLESIZE, KEYSIZE, SLOT)
	forestLsh := NewForestLSHFromData(params, data)

	idx := rand.Perm(ROWS)[:TEST_INSTANCES]
	validIdx := rand.Perm(ROWS)[:TEST_RESULTS]
	for _, id := range idx {
		out := make(chan int)
		defer close(out)

		p := data.Points[id]
		avgDist := 0.0
		count := 0
		forestLsh.Query(p, TEST_RESULTS, out)
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
			fmt.Printf("num found %d, avg dist %f, valid dist %f\n", count, avgDist, validAvgDist)
		}
	}
}
