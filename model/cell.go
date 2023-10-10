package model

import (
	"github.com/zhangjiuwei/sudoku_resolver/util"
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

type Cell struct {
	Row              int
	Col              int
	Defined          bool
	Value            int
	Candidates       mapset.Set
	valueSetTriggers []func(int)
}

func (c *Cell) Init() {
	c.Row = 0
	c.Col = 0
	c.Defined = false
	c.Value = 0
	c.Candidates = mapset.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	c.valueSetTriggers = make([]func(int), 0)
}

func (c *Cell) IsDefined() bool {
	return c.Defined
}

func (c *Cell) IsSet() bool {
	return c.Defined || c.Value > 0
}

func (c *Cell) RegistTrigger(callback func(int)) {
	c.valueSetTriggers = append(c.valueSetTriggers, callback)
}

func (c *Cell) Define(value int) {
	c.Value = value
	c.Candidates.Clear()
	c.Defined = true

	for _, f := range c.valueSetTriggers {
		f(value)
	}
}

func (c *Cell) SetValue(value int) {
	c.Value = value
	c.Candidates.Clear()

	for _, f := range c.valueSetTriggers {
		f(value)
	}
}

func (c *Cell) CutCandidates(cut mapset.Set) bool {
	if c.IsSet() {
		return false
	}
	newCan := c.Candidates.Difference(cut)
	if newCan.IsProperSubset(c.Candidates) {
		c.Candidates = newCan
		return true
	}
	return false
}

func (c *Cell) CutCandidate(cut int) bool {
	if c.IsSet() {
		return false
	}
	if c.Candidates.Contains(cut) {
		c.Candidates.Remove(cut)
		return true
	}
	return false
}

func (c *Cell) Print() string {
	if c.Value == 0 {
		return util.Yellow("_")
	} else if c.Defined {
		return util.Red(fmt.Sprintf("%d", c.Value))
	} else {
		return util.Green(fmt.Sprintf("%d", c.Value))
	}
}
