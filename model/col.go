package model

import (
	"code.byted.org/personal/zhangwei.1024/util"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"strings"
)

type Col struct {
	Idx        int
	Cells      []*Cell
	Candidates mapset.Set
	Values     mapset.Set
}

func (c *Col) Init() {
	c.Idx = 0
	c.Cells = make([]*Cell, 9)
	c.Candidates = mapset.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	c.Values = mapset.NewSet()
}

func (c *Col) CellSet(value int) {
	c.Candidates.Remove(value)
	c.Values.Add(value)
}

func (c *Col) Print() {
	fmt.Println(c)
}

func (c *Col) GetCells() []*Cell {
	return c.Cells
}

func (c *Col) ScanUniq() int {
	res := 0
	ms := c.Candidates.Clone()
	for can := range ms.Iter() {
		cnt := 0
		var cell *Cell
		for i := 0; i < 9; i++ {
			if c.Cells[i].IsSet() {
				continue
			}
			if c.Cells[i].Candidates.Contains(can) {
				if cnt == 0 {
					cnt++
					cell = c.Cells[i]
				} else {
					cnt++
					break
				}
			}
		}

		if cnt == 1 {
			cell.SetValue(can.(int))
			fmt.Println(fmt.Sprintf("Rule Uniq(Col), Cell:%d,%d, Value is set %d!", cell.Row, cell.Col, can.(int)))
			res++
		}

	}
	return res
}

func (c *Col) UnionFind() int {
	res := 0
	for subSetIter := range c.Candidates.PowerSet().Iter() {
		subSet := subSetIter.(mapset.Set)
		if subSet.Cardinality() <= 1 {
			continue
		}
		cnt := 0
		for _, cell := range c.GetCells() {
			if cell.Candidates.Cardinality() > 0 && cell.Candidates.IsSubset(subSet) {
				cnt++
			}
		}
		if cnt == subSet.Cardinality() {
			for _, cell := range c.GetCells() {
				if !cell.IsSet() && !cell.Candidates.IsSubset(subSet) {
					precan := cell.Candidates.Clone()
					if cell.CutCandidates(subSet) {
						res++
						fmt.Println(fmt.Sprintf("Rule UnionFind(Col), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
							strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
							strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
					}
				}
			}
		}
	}
	return res
}

func (c *Col) CandidateInSingleSquare(candidate int) int {
	idx := -1
	canIdxs := make([]int, 0)
	for _, cell := range c.Cells {
		if !cell.IsSet() && cell.Candidates.Contains(candidate) {
			canIdxs = append(canIdxs, cell.Row)
		}
	}
	for _, canIdx := range canIdxs {
		if idx == -1 {
			idx = canIdx / 3
		} else if idx != canIdx/3 {
			return -1
		}
	}
	return idx
}

func (c *Col) GetCandidateIdx(candidate int) mapset.Set {
	res := mapset.NewSet()
	if c.Candidates.Contains(candidate) {
		for _, cell := range c.Cells {
			if cell.Candidates.Contains(candidate) {
				res.Add(cell.Row)
			}
		}
	}
	return res
}
