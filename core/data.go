package golsh

import (
	"os"
	"strings"
	"bufio"
	"strconv"
	"math/rand"
	"log"
)

type Point struct {
	Id string
	Features []float32
}

type Dataset struct {
	Rows int
	VecLen int
	Points []*Point
	Index map[string]int
}

func NewDatasetWithHeader(filename string, idField string, colField string) (*Dataset, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	idIdx := 0
	colIdxes := []int{}
	{
		scanner.Scan()
		values := strings.Split(scanner.Text(), "\t")
		for i, v := range values {
			if v == idField {
				idIdx = i
			} else if strings.HasPrefix(v, colField) {
				colIdxes = append(colIdxes, i)
			}
		}
	}
	log.Printf("%v\n", colIdxes)

	d := Dataset{}
	d.Rows = 0
	d.VecLen = len(colIdxes)
	d.Points = []*Point{}
	d.Index = map[string]int{}
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "\t")
		id := values[idIdx]
		features := make([]float32, d.VecLen)
		for i, vs := range colIdxes {
			v, _ := strconv.ParseFloat(values[vs], 32)
			features[i] = float32(v)
		}
		p := Point{id, features}
		d.Index[id] = d.Rows
		d.Rows += 1
		if d.Rows % 10000 == 0 {
			log.Printf("load %d lines\n sample: %v\n", d.Rows, p)
		}
		d.Points = append(d.Points, &p)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &d, nil
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
	d.Index = map[string]int{}
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
		d.Index[id] = d.Rows
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
	d.Index = map[string]int{}
	for i := 0; i < rows; i++ {
		p := Point{}
		p.Id = strconv.Itoa(i)
		p.Features = make([]float32, vecLen)
		for j := 0; j < vecLen; j++ {
			p.Features[j] = rand.Float32()
		}
		d.Points = append(d.Points, &p)
		d.Index[p.Id] = i
	}
	return &d
}
