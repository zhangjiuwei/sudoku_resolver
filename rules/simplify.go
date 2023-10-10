package rules

import (
	"github.com/zhangjiuwei/sudoku_resolver/model"
	"github.com/zhangjiuwei/sudoku_resolver/util"
	"fmt"
	"strings"
)

func Simplify(p *model.Puzzle) int {
	total := 0
	for {
		cnt := 0
		for i := 0; i < 9; i++ {
			r := p.Rows[i]
			for canIter := range r.Candidates.Iter() {
				candidate := canIter.(int)
				if idx := r.CandidateInSingleSquare(candidate); idx >= 0 {
					//fmt.Println(fmt.Sprintf("Row:%d, candidate:%d, idx:%d", i, candidate, idx))
					for _, cell := range p.Squares[i/3][idx].GetCells() {
						if !cell.IsSet() && cell.Row != i {
							precan := cell.Candidates.Clone()
							if cell.CutCandidate(candidate) {
								cnt++
								fmt.Println(fmt.Sprintf("Rule Simplify(Row), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
									strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
									strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
							}
						}
					}
				}
			}
		}
		for j := 0; j < 9; j++ {
			c := p.Cols[j]
			for canIter := range c.Candidates.Iter() {
				candidate := canIter.(int)
				if idx := c.CandidateInSingleSquare(candidate); idx >= 0 {
					//fmt.Println(fmt.Sprintf("Col:%d, candidate:%d, idx:%d", j, candidate, idx))
					for _, cell := range p.Squares[idx][j/3].GetCells() {
						if !cell.IsSet() && cell.Col != j {
							precan := cell.Candidates.Clone()
							if cell.CutCandidate(candidate) {
								cnt++
								fmt.Println(fmt.Sprintf("Rule Simplify(Col), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
									strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
									strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
							}
						}
					}
				}
			}
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				s := p.Squares[i][j]
				for canIter := range s.Candidates.Iter() {
					candidate := canIter.(int)
					if idx := s.CandidateInSingleRow(candidate); idx >= 0 {
						for _, cell := range p.Rows[idx].GetCells() {
							if cell.Col/3 != j && !cell.IsSet() {
								precan := cell.Candidates.Clone()
								if cell.CutCandidate(candidate) {
									cnt++
									fmt.Println(fmt.Sprintf("Rule Simplify(Square), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
										strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
										strings.Join(util.SliceElementToString(cell.Candidates.ToSlice()), ",")))
								}
							}
						}
					} else if idx = s.CandidateInSingleCol(candidate); idx >= 0 {
						for _, cell := range p.Cols[idx].GetCells() {
							if cell.Row/3 != i && !cell.IsSet() {
								precan := cell.Candidates.Clone()
								if cell.CutCandidate(candidate) {
									cnt++
									fmt.Println(fmt.Sprintf("Rule Simplify(Square), Cell:%d,%d, candidates is cut from [%s] to [%s]!", cell.Row, cell.Col,
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
