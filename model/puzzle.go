package model

import (
	"bufio"
	"github.com/zhangjiuwei/sudoku_resolver/util"
	"fmt"
	"github.com/spf13/cast"
	"os"
	"strings"
)

type Puzzle struct {
	Cells   [][]*Cell
	Rows    []*Row
	Cols    []*Col
	Squares [][]*Square
	rules   []func(puzzle *Puzzle) int
}

func (p *Puzzle) Init() {
	p.Cells = make([][]*Cell, 9)
	p.rules = make([]func(puzzle *Puzzle) int, 0)
	for i := 0; i < 9; i++ {
		p.Cells[i] = make([]*Cell, 9)
	}
	p.Rows = make([]*Row, 9)
	p.Cols = make([]*Col, 9)
	p.Squares = make([][]*Square, 3)
	for i := 0; i < 3; i++ {
		p.Squares[i] = make([]*Square, 3)
	}

	for i := 0; i < 9; i++ {
		p.Rows[i] = new(Row)
		p.Rows[i].Init()
		p.Rows[i].Idx = i
		for j := 0; j < 9; j++ {
			if i == 0 {
				p.Cols[j] = new(Col)
				p.Cols[j].Init()
				p.Cols[j].Idx = j
			}
			if i%3 == 0 && j%3 == 0 {
				p.Squares[i/3][j/3] = new(Square)
				p.Squares[i/3][j/3].Init()
				p.Squares[i/3][j/3].Row = i / 3
				p.Squares[i/3][j/3].Col = j / 3
			}

			c := new(Cell)
			c.Init()
			c.Row = i
			c.Col = j

			p.Rows[i].Cells[j] = c
			c.RegistTrigger(p.Rows[i].CellSet)
			p.Cols[j].Cells[i] = c
			c.RegistTrigger(p.Cols[j].CellSet)
			p.Squares[i/3][j/3].Cells[i%3][j%3] = c
			c.RegistTrigger(p.Squares[i/3][j/3].CellSet)
			p.Cells[i][j] = c
		}
	}
}

func (p *Puzzle) Print() {
	for i := 0; i < 9; i++ {
		if i > 0 && i%3 == 0 {
			fmt.Println(util.Purple("======+=======+======"))
		}
		var chars []string
		for j := 0; j < 9; j++ {
			if j > 0 && j%3 == 0 {
				chars = append(chars, util.Purple("‖"))
			}
			chars = append(chars, p.Cells[i][j].Print())
		}
		fmt.Println(strings.Join(chars, " "))
	}
}

func (p *Puzzle) PrintDetail() {
	for i := 0; i < 9; i++ {
		if i > 0 && i%3 == 0 {
			fmt.Println(util.Purple("===============+=================+==============="))
			//fmt.Println(util.Purple("               ‖                 ‖               "))
		}
		lines := make([][]string, 3)
		for j := 0; j < 9; j++ {
			if j > 0 && j%3 == 0 {
				for k := 0; k < 3; k++ {
					lines[k] = append(lines[k], util.Purple("‖"))
				}
			}
			if p.Cells[i][j].IsSet() {
				for k := 0; k < 3; k++ {
					lines[k] = append(lines[k], strings.Repeat(p.Cells[i][j].Print(), 3))
				}
			} else {
				for k := 0; k < 3; k++ {
					lines[k] = append(lines[k], fmt.Sprintf("%s%s%s",
						util.If(p.Cells[i][j].Candidates.Contains(k*3+1), util.Yellow(cast.ToString(k*3+1)), util.Yellow("·")),
						util.If(p.Cells[i][j].Candidates.Contains(k*3+2), util.Yellow(cast.ToString(k*3+2)), util.Yellow("·")),
						util.If(p.Cells[i][j].Candidates.Contains(k*3+3), util.Yellow(cast.ToString(k*3+3)), util.Yellow("·"))))
				}
			}
		}
		for k := 0; k < 3; k++ {
			fmt.Println(strings.Join(lines[k], "  "))
			//fmt.Println(strings.Join(lines[k], util.Purple("|")))
		}
		if i%3 != 2 {
			fmt.Println(util.Purple("               ‖                 ‖               "))
		}
	}
}

func (p *Puzzle) Load(file string) bool {
	cf, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("read file err", err)
		return false
	}
	defer cf.Close()
	buf := bufio.NewScanner(cf)
	for i := 0; i < 9; i++ {
		if !buf.Scan() {
			fmt.Println("read file err")
			return false
		}
		line := buf.Text()
		//fmt.Println(line)
		if len(line) != 9 {
			fmt.Println("line format err, line: ", line, len(line))
			return false
		}
		for j := 0; j < 9; j++ {
			n := cast.ToInt(fmt.Sprintf("%c", line[j]))
			if n >= 1 && n <= 9 {
				p.Cells[i][j].Define(n)
			}
		}
	}
	return true
}

func (p *Puzzle) Solve() bool {
	round := 0
	for {
		cnt := 0
		for _, rule := range p.rules {
			cnt += rule(p)
		}
		if cnt <= 0 {
			break
		} else {
			p.Print()
			fmt.Println("Round:", round)
			round++
		}
		//break
		//fmt.Println(p.Cells[0][8])
		//fmt.Println(p.Rows[0].Candidates)
	}
	//fmt.Println(p.Cells[0][3])
	//fmt.Println(p.Cells[0][4])

	return p.IsSolved()
}

func (p *Puzzle) RegistRule(rule func(puzzle *Puzzle) int) {
	p.rules = append(p.rules, rule)
}

func (p *Puzzle) IsSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !p.Cells[i][j].IsSet() {
				return false
			}
		}
	}
	return true
}
