package golsh

import "log"

type Node struct {
	Level int
	Key int
	Ind []int
	Children map[int]*Node
	Father *Node
}

func (n *Node) debugInfo() {
	log.Printf("level %d, key %d, num ids %d", n.Level, n.Key, len(n.Ind))
	for _, child := range n.Children {
		child.debugInfo()
	}
}

func (n *Node) insert(level, id int, key *HashKey) {
	n.Ind = append(n.Ind, id)
	if level != len(*key) {
		var child *Node
		if nodeChild, ok := n.Children[(*key)[level]]; !ok {
			child = &Node{
				level + 1,
				(*key)[level],
				make([]int,0),
				make(map[int]*Node),
				n,
			}
			n.Children[(*key)[level]] = child
		} else {
			child = nodeChild
		}
		child.insert(level + 1, id, key)
	}
}

func (n *Node) lookup(hvs *HashKey) *Node {
	c := n
	for l := 0; l < len(*hvs); l++ {
		if child, ok := c.Children[(*hvs)[l]]; ok {
			c = child
		} else {
			break
		}
	}
	return c
}

type ForestLSH struct {
	*Params
	Roots []*Node
	Data *Dataset
}

func (lsh *ForestLSH) Query(p *Point, n int, out chan int) {
	go func() {
		hvs := lsh.Hash(p)
		candidates := []int{}
		nodes := []*Node{}
		seens := map[int]bool{}
		for l := lsh.KeySize; l >= 0; l-- {
			for i, root := range lsh.Roots {
				nodes = append(nodes, root.lookup(&hvs[i]))
			}

			for _, node := range nodes {
				if node == nil || node.Level < l {
					continue
				}
				for _, candidate := range node.Ind {
					if _, seen := seens[candidate]; !seen {
						seens[candidate] = true
						candidates = append(candidates, candidate)
					}
				}
				node = node.Father
			}

			if len(candidates) >= n {
				break
			}
		}
		for _, q := range candidates {
			out <- q
		}
		out <- -1
	}()
}

func NewForestLSH(params *Params) *ForestLSH {
	roots := make([]*Node, params.TableSize)
	for i, _ := range roots {
		roots[i] = &Node{
			0, 0, make([]int, 0), make(map[int]*Node), nil,
		}
	}
	return &ForestLSH{
		params, roots, nil,
	}
}

func NewForestLSHFromData(params *Params, dataset *Dataset) *ForestLSH {
	forest := NewForestLSH(params)
	for idx, point := range dataset.Points {
		hvs := forest.Hash(point)
		for i := range forest.Roots {
			hv := hvs[i]
			forest.Roots[i].insert(0, idx, &hv)
		}
	}
	for i := range forest.Roots {
		log.Printf("tree %d:\n", i)
		forest.Roots[i].debugInfo()
	}
	return forest
}
