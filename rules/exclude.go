package rules

import (
	"code.byted.org/personal/zhangwei.1024/model"
	"code.byted.org/personal/zhangwei.1024/util"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"strings"
)

func Exclude(p *model.Puzzle) int {
	total := 0
	for {
		cnt := 0
		for i := 1; i <= 9; i++ {
			rowCanColIdxs := mapset.NewSet()
			rowCanIdxArr := make([]mapset.Set, 9)
			for _, row := range p.Rows {
				idxs := row.GetCandidateIdx(i)
				if idxs.Cardinality() > 0 {
					rowCanColIdxs = rowCanColIdxs.Union(idxs)
					rowCanIdxArr[row.Idx] = idxs
				}
			}
			if rowCanColIdxs.Cardinality() <= 1 {
				continue
			}
			for subSetIter := range rowCanColIdxs.PowerSet().Iter() {
				subSet := subSetIter.(mapset.Set)
				if subSet.Cardinality() <= 1 {
					continue
				}
				rowIdxs := mapset.NewSet()
				for _, row := range p.Rows {
					if rowCanIdxArr[row.Idx] != nil && rowCanIdxArr[row.Idx].Cardinality() > 1 && rowCanIdxArr[row.Idx].IsSubset(subSet) {
						rowIdxs.Add(row.Idx)
					}
				}
				if rowIdxs.Cardinality() == subSet.Cardinality() {
					for colIdxIter := range subSet.Iter() {
						colIdx := colIdxIter.(int)
						for _, cell := range p.Cols[colIdx].Cells {
							if !cell.IsSet() && !rowIdxs.Contains(cell.Row) {
								precan := cell.Candidates.Clone()
								if cell.CutCandidate(i) {
									cnt++
									fmt.Println(fmt.Sprintf("Rule Exclude(Row), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
										strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
										strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
								}
							}
						}
					}
				}
			}
		}

		for i := 1; i <= 9; i++ {
			colCanRowIdxs := mapset.NewSet()
			rowCanIdxArr := make([]mapset.Set, 9)
			for _, col := range p.Cols {
				idxs := col.GetCandidateIdx(i)
				if idxs.Cardinality() > 0 {
					colCanRowIdxs = colCanRowIdxs.Union(idxs)
					rowCanIdxArr[col.Idx] = idxs
				}
			}
			if colCanRowIdxs.Cardinality() <= 1 {
				continue
			}
			for subSetIter := range colCanRowIdxs.PowerSet().Iter() {
				subSet := subSetIter.(mapset.Set)
				if subSet.Cardinality() <= 1 {
					continue
				}
				colIdxs := mapset.NewSet()
				for _, col := range p.Cols {
					if rowCanIdxArr[col.Idx] != nil && rowCanIdxArr[col.Idx].Cardinality() > 1 && rowCanIdxArr[col.Idx].IsSubset(subSet) {
						colIdxs.Add(col.Idx)
					}
				}
				if colIdxs.Cardinality() == subSet.Cardinality() {
					for rowIdxIter := range subSet.Iter() {
						rowIdx := rowIdxIter.(int)
						for _, cell := range p.Rows[rowIdx].Cells {
							if !cell.IsSet() && !colIdxs.Contains(cell.Col) {
								precan := cell.Candidates.Clone()
								if cell.CutCandidate(i) {
									cnt++
									fmt.Println(fmt.Sprintf("Rule Exclude(Col), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
										strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
										strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
								}
							}
						}
					}
				}
			}
		}

		total += cnt
		if cnt <= 0 {
			break
		}
	}

	return total
}
