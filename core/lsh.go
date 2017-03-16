package golsh

import (
	"math"
	"math/rand"
	"time"
)

type Params struct {
	VecLen int
	TableSize int
	KeySize int
	SlotSize float32
	Weights [][]Point
	Bias [][]float32
}

type HashKey []int

func NewParams(vecLen, tableSize, keySize int, slotSize float32) *Params {
	weights := make([][]Point, tableSize)
	bias := make([][]float32, tableSize)
	random := rand.New(rand.NewSource(int64(time.Now().Second())))
	for i := range weights {
		weights[i] = make([]Point, keySize)
		bias[i] = make([]float32, keySize)
		for j := range weights[i] {
			weights[i][j] = Point{}
			weights[i][j].Features = make([]float32, vecLen)
			for t := 0; t < vecLen; t++ {
				weights[i][j].Features[t] = float32(random.NormFloat64())
			}
			bias[i][j] = random.Float32() * slotSize
		}
	}
	return &Params{
		vecLen, tableSize, keySize, slotSize,weights, bias,
	}
}

// Returns all hash values for all hash tables.
func (p *Params) Hash(point *Point) []HashKey {
	ret := make([]HashKey, p.TableSize)
	for i := range ret {
		s := make([]int, p.KeySize)
		for j := 0; j < p.KeySize; j++ {
			v := (point.Dot(&p.Weights[i][j]) + p.Bias[i][j]) / p.SlotSize
			s[j] = int(math.Floor(float64(v)))
		}
		ret[i] = s
	}
	return ret
}
