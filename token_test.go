package pcomb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueToken(t *testing.T) {
	token := NewValueToken("hi")

	assert.NotPanics(t, func() {
		token.Value()
	})

	assert.Panics(t, func() {
		token.Tokens()
	})

	value, ok := token.MaybeValue()
	assert.True(t, ok)
	assert.Equal(t, "hi", value)

	_, ok = token.MaybeTokens()
	assert.False(t, ok)
}

func TestTokensToken(t *testing.T) {
	token := NewTokensToken(
		NewValueToken("foo"),
		NewValueToken("bar"))

	assert.NotPanics(t, func() {
		token.Tokens()
	})

	assert.Panics(t, func() {
		token.Value()
	})

	tokens, ok := token.MaybeTokens()
	assert.True(t, ok)
	assert.Equal(t, 2, len(tokens))
	assert.Equal(t, "foo", tokens[0].Value())
	assert.Equal(t, "bar", tokens[1].Value())

	_, ok = token.MaybeValue()
	assert.False(t, ok)
}
