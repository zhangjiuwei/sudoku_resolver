package main

import (
	"code.byted.org/personal/zhangwei.1024/model"
	"code.byted.org/personal/zhangwei.1024/rules"
	"fmt"
	"time"
)

func main() {
	p := new(model.Puzzle)
	p.Init()
	p.Load("puzzle")
	p.Print()

	p.RegistRule(rules.Intersect)
	p.RegistRule(rules.Uniq)
	p.RegistRule(rules.UnionFind)
	p.RegistRule(rules.Simplify)
	p.RegistRule(rules.Exclude)

	start := time.Now()
	if p.Solve() {
		fmt.Println("Succeed!!!")
	} else {
		fmt.Println("Failed!!!")
		p.PrintDetail()
	}
	end := time.Now()
	cost := end.Sub(start).Nanoseconds() / 1000
	fmt.Println("Time Cost:", cost)
}
