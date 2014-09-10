package pcomb

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNoopParser(t *testing.T) {
	log.Println(Noop)

	result := Noop.Parse(InputFromString("hello"))

	assert.Nil(t, result.Token)
	assert.Equal(t, result.NextInput.Value(), "hello")
	assert.Equal(t, result.NextInput.Pos(), 0)
	assert.Nil(t, result.Err)
}

func TestIgnoreParser(t *testing.T) {
	input := "foobar"
	ignoreFoo := Ignore(String("foo"))

	log.Println(ignoreFoo)

	result := ignoreFoo.Parse(InputFromString(input))

	assert.Nil(t, result.Err)
	assert.Nil(t, result.Token)
	assert.Equal(t, "bar", result.NextInput.Value())
}

func TestStringParser(t *testing.T) {
	input := "hello"
	parser := String("hello")

	log.Println(parser)

	result := parser.Parse(InputFromString(input))

	assert.Equal(t, NewValueToken("hello"), result.Token)
	assert.Nil(t, result.Err)
	assert.Empty(t, result.NextInput.Value())
}

func TestStringParserFail(t *testing.T) {
	input := "foobar"
	parser := String("bar")

	result := parser.Parse(InputFromString(input))

	assert.NotNil(t, result.Err)
	assert.Nil(t, result.Token)
	assert.Equal(t, "foobar", result.NextInput.Value())
}

func TestSequenceParser(t *testing.T) {
	input := "helloworld"
	parser := Sequence(String("hello"), String("world"))
	expectedOutput := NewTokensToken(
		NewValueToken("hello"),
		NewValueToken("world"))

	log.Println(parser)

	result := parser.Parse(InputFromString(input))

	assert.Nil(t, result.Err)
	assert.Equal(t, expectedOutput, result.Token)
	assert.Empty(t, result.NextInput.Value())
}

func TestAnyOfParser(t *testing.T) {
	input := "foobar"
	parser := AnyOf(String("foo"), String("baz"))

	log.Println(parser)

	result := parser.Parse(InputFromString(input))

	assert.Nil(t, result.Err)
	assert.Equal(t, "foo", result.Token.Value())
	assert.Equal(t, "bar", result.NextInput.Value())

	result = parser.Parse(result.NextInput)

	assert.NotNil(t, result.Err)
	assert.Equal(t, "bar", result.NextInput.Value())
}
