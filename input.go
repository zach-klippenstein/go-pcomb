package pcomb

import (
	"fmt"
)

type Input interface {
	Value() string
	Pos() int
	Advance(int) Input
	String() string
}

type StringInput struct {
	Whole string
	pos   int
}

var _ Input = StringInput{}

func InputFromString(str string) StringInput {
	return StringInput{
		Whole: str,
	}
}

func (i StringInput) Pos() int {
	return i.pos
}

func (i StringInput) Value() string {
	return i.Whole[i.pos:]
}

func (i StringInput) Advance(n int) Input {
	return StringInput{
		Whole: i.Whole,
		pos:   i.pos + n,
	}
}

func (i StringInput) String() string {
	value := i.Value()
	var excerpt string

	if len(value) > 10 {
		excerpt = value[:10] + "…"
	} else {
		excerpt = value
	}

	return fmt.Sprintf("StringInput@%d:'%s'", i.pos, excerpt)
}
