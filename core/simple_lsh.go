package golsh

import "fmt"

type SimpleLSH struct {
	*Params
	Tables []map[string][]int
}

func (lsh *SimpleLSH) toHashTableKeys(keys []HashKey) []string {
	hashKeys := make([]string, lsh.Params.TableSize)
	for i, key := range keys {
		s := ""
		for _, hashVal := range key {
			s += fmt.Sprintf("%.16x", hashVal)
		}
		hashKeys[i] = s
	}
	return hashKeys
}

// Query finds n neighbours, it is not k-NN query.
func (lsh *SimpleLSH) Query(p *Point, n int, out chan int) {
	go func() {
		count := 0
		hvs := lsh.toHashTableKeys(lsh.Params.Hash(p))
		checked := make(map[int]bool)
		for i, table := range lsh.Tables {
			if posting, exist := table[hvs[i]]; exist {
				for _, id := range posting {
					if _, exist := checked[id]; !exist {
						checked[id] = true
						out <- id
						count += 1
						if count >= n {
							break
						}
					}
				}
			}
			if count >= n {
				break
			}
		}
		out <- -1
	}()
}

func NewSimpleLSH(params *Params) *SimpleLSH {
	tables := make([]map[string][]int, params.TableSize)
	for i := range tables {
		tables[i] = map[string][]int{}
	}
	return &SimpleLSH{
		params, tables,
	}
}

func NewSimpleLSHFromData(params *Params, dataset *Dataset) *SimpleLSH {
	lsh := NewSimpleLSH(params)
	for idx, p := range dataset.Points {
		hvs := lsh.toHashTableKeys(params.Hash(p))
		for i := range lsh.Tables {
			hv := hvs[i]
			table := lsh.Tables[i]
			if _, exist := table[hv]; !exist {
				table[hv] = []int{}
			}
			table[hv] = append(table[hv], idx)
		}
	}
	return lsh
}
