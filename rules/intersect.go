package rules

import (
	"github.com/zhangjiuwei/sudoku_resolver/model"
	"github.com/zhangjiuwei/sudoku_resolver/util"
	"fmt"
	"strings"
)

func Intersect(p *model.Puzzle) int {
	total := 0
	for {
		cnt := 0
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if p.Cells[i][j].IsSet() {
					continue
				}
				s := p.Cells[i][j].Candidates
				s = s.Intersect(p.Rows[i].Candidates)
				s = s.Intersect(p.Cols[j].Candidates)
				s = s.Intersect(p.Squares[i/3][j/3].Candidates)
				if s.Cardinality() == 1 {
					n := s.Pop().(int)
					p.Cells[i][j].SetValue(n)
					fmt.Println(fmt.Sprintf("Rule Intersect, Cell:%d,%d, Value is set %d!", i, j, n))
					cnt++
				} else {
					if s.IsProperSubset(p.Cells[i][j].Candidates) {
						precan := p.Cells[i][j].Candidates
						p.Cells[i][j].Candidates = s
						fmt.Println(fmt.Sprintf("Rule Intersect, Cell:%d,%d, candidates is cut from [%s] to [%s]!", p.Cells[i][j].Row, p.Cells[i][j].Col,
							strings.Join(util.SliceElementToString(precan.ToSlice()), ","),
							strings.Join(util.SliceElementToString(s.ToSlice()), ",")))
						cnt++
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
