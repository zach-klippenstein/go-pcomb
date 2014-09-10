package pcomb

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNoopParser(t *testing.T) {
	log.Println(Noop)

	result := Noop.Parse(InputFromString("hello"))

	assert.NoError(t, result.Err)
	assert.Nil(t, result.Token)
	assert.Equal(t, result.NextInput.Value(), "hello")
	assert.Equal(t, result.NextInput.Pos(), 0)
}

func TestIgnoreParser(t *testing.T) {
	input := "foobar"
	ignoreFoo := Ignore(String("foo"))

	log.Println(ignoreFoo)

	result := ignoreFoo.Parse(InputFromString(input))

	assert.NoError(t, result.Err)
	assert.Nil(t, result.Token)
	assert.Equal(t, "bar", result.NextInput.Value())
}

func TestStringParser(t *testing.T) {
	input := "hello"
	parser := String("hello")

	log.Println(parser)

	result := parser.Parse(InputFromString(input))

	assert.Equal(t, NewValueToken("hello"), result.Token)
	assert.NoError(t, result.Err)
	assert.Empty(t, result.NextInput.Value())
}

func TestStringParserFail(t *testing.T) {
	input := "foobar"
	parser := String("bar")

	result := parser.Parse(InputFromString(input))

	assert.Error(t, result.Err)
	assert.Nil(t, result.Token)
	assert.Equal(t, "foobar", result.NextInput.Value())
}

func TestSequenceParser(t *testing.T) {
	input := InputFromString("helloworld")

	parser := Sequence(String("hello"), String("world"))
	log.Println(parser)
	result := parser.Parse(input)
	assert.NoError(t, result.Err)
	assert.Equal(t, NewTokensToken(NewValueToken("hello"), NewValueToken("world")), result.Token)
	assert.Empty(t, result.NextInput.Value())

	parser = Sequence(String("world"), String("hello"))
	result = parser.Parse(input)
	assert.Error(t, result.Err)
	assert.Nil(t, result.Token)
	assert.Equal(t, "helloworld", result.NextInput.Value())
}

func TestFirstOfParser(t *testing.T) {
	input := "foobar"
	parser := FirstOf(String("foo"), String("baz"))

	log.Println(parser)

	result := parser.Parse(InputFromString(input))

	assert.NoError(t, result.Err)
	assert.Equal(t, "foo", result.Token.Value())
	assert.Equal(t, "bar", result.NextInput.Value())

	result = parser.Parse(result.NextInput)

	assert.Error(t, result.Err)
	assert.Equal(t, "bar", result.NextInput.Value())
}

func TestRepeatParser(t *testing.T) {
	input := InputFromString("aaa")
	a := String("a")

	parser := Repeat(a, 1, 1)
	log.Println(parser)
	result := parser.Parse(input)
	assert.Equal(t, NewTokensToken(NewValueToken("a")), result.Token)
	assert.Equal(t, "aa", result.NextInput.Value())

	parser = Repeat(a, 0, 1)
	result = parser.Parse(input)
	assert.Equal(t, NewTokensToken(NewValueToken("a")), result.Token)
	assert.Equal(t, "aa", result.NextInput.Value())

	parser = Repeat(a, 0, 2)
	result = parser.Parse(input)
	assert.Equal(t, NewTokensToken(NewValueToken("a"), NewValueToken("a")), result.Token)
	assert.Equal(t, "a", result.NextInput.Value())

	parser = Repeat(a, 0, 3)
	result = parser.Parse(input)
	assert.Equal(t, NewTokensToken(NewValueToken("a"), NewValueToken("a"), NewValueToken("a")), result.Token)
	assert.Equal(t, "", result.NextInput.Value())

	parser = Repeat(a, 0, 0)
	log.Println(parser)
	result = parser.Parse(input)
	assert.Equal(t, NewTokensToken(NewValueToken("a"), NewValueToken("a"), NewValueToken("a")), result.Token)
	assert.Equal(t, "", result.NextInput.Value())

	b := String("b")
	parser = Repeat(b, 0, 1)
	result = parser.Parse(input)
	assert.Nil(t, result.Token)
	assert.Equal(t, "aaa", result.NextInput.Value())

	parser = Repeat(b, 1, 1)
	result = parser.Parse(input)
	assert.Nil(t, result.Token)
	assert.Equal(t, "aaa", result.NextInput.Value())

	input = InputFromString("abb")
	parser = Repeat(a, 1, 2)
	result = parser.Parse(input)
	assert.NoError(t, result.Err)
	assert.Equal(t, NewTokensToken(NewValueToken("a")), result.Token)
	assert.Equal(t, "bb", result.NextInput.Value())
}
