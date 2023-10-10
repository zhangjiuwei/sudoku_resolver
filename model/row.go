package model

import (
	"code.byted.org/personal/zhangwei.1024/util"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"strings"
)

type Row struct {
	Idx        int
	Cells      []*Cell
	Candidates mapset.Set
	Values     mapset.Set
}

func (r *Row) Init() {
	r.Idx = 0
	r.Cells = make([]*Cell, 9)
	r.Candidates = mapset.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	r.Values = mapset.NewSet()
}

func (r *Row) CellSet(value int) {
	r.Candidates.Remove(value)
	r.Values.Add(value)
}

func (r *Row) GetCells() []*Cell {
	return r.Cells
}

func (r *Row) Print() {
	var chars []string
	for j := 0; j < 9; j++ {
		if j > 0 {
			if j%3 == 0 {
				chars = append(chars, util.Purple("‖"))
			}
		}
		chars = append(chars, r.Cells[j].Print())
	}
	fmt.Println(strings.Join(chars, " "))
}

func (r *Row) ScanUniq() int {
	res := 0
	ms := r.Candidates.Clone()
	for c := range ms.Iter() {
		cnt := 0
		var cell *Cell
		for i := 0; i < 9; i++ {
			if r.Cells[i].IsSet() {
				continue
			}
			if r.Cells[i].Candidates.Contains(c) {
				if cnt == 0 {
					cnt++
					cell = r.Cells[i]
				} else {
					cnt++
					break
				}
			}
		}

		if cnt == 1 {
			cell.SetValue(c.(int))
			fmt.Println(fmt.Sprintf("Rule Uniq(Row), Cell:%d,%d, Value is set %d!", cell.Row, cell.Col, c.(int)))
			res++
		}

	}
	return res
}

func (r *Row) UnionFind() int {
	res := 0
	for subSetIter := range r.Candidates.PowerSet().Iter() {
		subSet := subSetIter.(mapset.Set)
		if subSet.Cardinality() <= 1 {
			continue
		}
		cnt := 0
		for _, cell := range r.GetCells() {
			if cell.Candidates.Cardinality() > 0 && cell.Candidates.IsSubset(subSet) {
				cnt++
			}
		}
		if cnt == subSet.Cardinality() {
			for _, cell := range r.GetCells() {
				if !cell.IsSet() && !cell.Candidates.IsSubset(subSet) {
					precan := cell.Candidates.Clone()
					if cell.CutCandidates(subSet) {
						res++
						fmt.Println(fmt.Sprintf("Rule UnionFind(Row), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
							strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
							strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
					}
				}
			}
		}
	}
	return res
}

func (r *Row) CandidateInSingleSquare(candidate int) int {
	idx := -1
	canIdxs := make([]int, 0)
	for _, cell := range r.Cells {
		if !cell.IsSet() && cell.Candidates.Contains(candidate) {
			canIdxs = append(canIdxs, cell.Col)
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

func (r *Row) GetCandidateIdx(candidate int) mapset.Set {
	res := mapset.NewSet()
	if r.Candidates.Contains(candidate) {
		for _, cell := range r.Cells {
			if cell.Candidates.Contains(candidate) {
				res.Add(cell.Col)
			}
		}
	}
	return res
}
