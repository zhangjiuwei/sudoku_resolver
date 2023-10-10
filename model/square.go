package model

import (
	"code.byted.org/personal/zhangwei.1024/util"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"strings"
)

type Square struct {
	Row        int
	Col        int
	Cells      [][]*Cell
	Candidates mapset.Set
	Values     mapset.Set
}

func (s *Square) Init() {
	s.Row = 0
	s.Col = 0
	s.Cells = make([][]*Cell, 3)
	for i := 0; i < 3; i++ {
		s.Cells[i] = make([]*Cell, 3)
	}
	s.Candidates = mapset.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	s.Values = mapset.NewSet()
}

func (s *Square) CellSet(value int) {
	s.Candidates.Remove(value)
	s.Values.Add(value)
}

func (s *Square) GetCells() []*Cell {
	res := make([]*Cell, 0)
	for i := 0; i < 3; i++ {
		res = append(res, s.Cells[i]...)
	}
	return res
}

func (s *Square) Print() {
	fmt.Println(s)
}

func (s *Square) ScanUniq() int {
	res := 0
	ms := s.Candidates.Clone()
	for c := range ms.Iter() {
		cnt := 0
		var cell *Cell
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if s.Cells[i][j].IsSet() {
					continue
				}
				if s.Cells[i][j].Candidates.Contains(c) {
					if cnt == 0 {
						cnt++
						cell = s.Cells[i][j]
					} else {
						cnt++
						break
					}
				}
			}

		}

		if cnt == 1 {
			cell.SetValue(c.(int))
			fmt.Println(fmt.Sprintf("Rule Uniq(Square), Cell:%d,%d, Value set is %d!", cell.Row, cell.Col, c.(int)))
			res++
		}

	}
	return res
}

func (s *Square) UnionFind() int {
	res := 0
	for subSetIter := range s.Candidates.PowerSet().Iter() {
		subSet := subSetIter.(mapset.Set)
		if subSet.Cardinality() == 1 {
			continue
		}
		cnt := 0
		for _, cell := range s.GetCells() {
			if cell.Candidates.Cardinality() > 0 && cell.Candidates.IsSubset(subSet) {
				cnt++
			}
		}
		if cnt == subSet.Cardinality() {
			for _, cell := range s.GetCells() {
				if !cell.IsSet() && !cell.Candidates.IsSubset(subSet) {
					precan := cell.Candidates.Clone()
					if cell.CutCandidates(subSet) {
						res++
						fmt.Println(fmt.Sprintf("Rule UnionFind(Square), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
							strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
							strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
					}
				}
			}
		}
	}
	return res
}

func (s *Square) CandidateInSingleRow(candidate int) int {
	idxSet := mapset.NewSet()
	for _, cell := range s.GetCells() {
		if !cell.IsSet() && cell.Candidates.Contains(candidate) {
			idxSet.Add(cell.Row)
		}
	}
	if idxSet.Cardinality() == 1 {
		return idxSet.Pop().(int)
	}
	return -1
}

func (s *Square) CandidateInSingleCol(candidate int) int {
	idxSet := mapset.NewSet()
	for _, cell := range s.GetCells() {
		if !cell.IsSet() && cell.Candidates.Contains(candidate) {
			idxSet.Add(cell.Col)
		}
	}
	if idxSet.Cardinality() == 1 {
		return idxSet.Pop().(int)
	}
	return -1
}
