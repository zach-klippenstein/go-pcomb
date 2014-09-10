package pcomb

import (
	"fmt"
)

// Input provides text to the parser.
// It is immutable, and mutators return new Inputs.
type Input interface {
	// Value returns the string value of the rest of the input
	Value() string

	// Pos returns the number of bytes that have already been consumed,
	// or the number of bytes into the source that this input represents.
	Pos() int

	// Advance returns an Input with a greater Pos.
	Advance(int) Input

	String() string
}

// StringInput is an Input that is backed by a string.
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
