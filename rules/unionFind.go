package rules

import "code.byted.org/personal/zhangwei.1024/model"

func UnionFind(p *model.Puzzle) int {
	total := 0
	for {
		cnt := 0
		for i := 0; i < 9; i++ {
			cnt += p.Rows[i].UnionFind()
		}
		for j := 0; j < 9; j++ {
			cnt += p.Cols[j].UnionFind()
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				cnt += p.Squares[i][j].UnionFind()
			}
		}
		total += cnt
		if cnt <= 0 {
			break
		}
	}
	return total
}
