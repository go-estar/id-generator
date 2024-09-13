package idGenerator

import (
	"strconv"
)

type IGenerator interface {
	NextId() (int64, error)
}

type Generator struct {
	IGenerator
}

func (g *Generator) Int64() int64 {
	id, _ := g.NextId()
	return id
}

func (g *Generator) Int() int {
	id, _ := g.NextId()
	return int(id)
}

func (g *Generator) String() string {
	id, _ := g.NextId()
	return strconv.FormatInt(id, 10)
}
