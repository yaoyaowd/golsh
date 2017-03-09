package golsh

import (
	"os"
	"strings"
	"bufio"
	"strconv"
	"math/rand"
)

type Point struct {
	Id string
	Features []float32
}

type Dataset struct {
	Rows int
	VecLen int
	Points []*Point
}

func NewDataset(filename string) (*Dataset, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d := Dataset{}
	d.Rows = 0
	d.VecLen = 0
	d.Points = []*Point{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		keyValue := strings.Split(scanner.Text(), "\t")
		id := keyValue[0]
		featureStrs := strings.Split(keyValue[1], ",")
		features := make([]float32, len(featureStrs))
		for i, vs := range featureStrs {
			v, _ := strconv.ParseFloat(vs, 32)
			features[i] = float32(v)
		}
		p := Point{id, features}

		d.VecLen = len(features)
		d.Rows += 1
		d.Points = append(d.Points, &p)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &d, nil
}

func NewRandomData(rows, vecLen int) (*Dataset) {
	d := Dataset{}
	d.Rows = rows
	d.VecLen = vecLen
	d.Points = []*Point{}
	for i := 0; i < rows; i++ {
		p := Point{}
		p.Id = strconv.Itoa(i)
		p.Features = make([]float32, vecLen)
		for j := 0; j < vecLen; j++ {
			p.Features[j] = rand.Float32()
		}
		d.Points = append(d.Points, &p)
	}
	return &d
}
