package pcomb

import (
	"fmt"
)

// Token is a parsed token that can either be a single string value
// or a slice of tokens.
type Token struct {
	stringValue string
	tokens      []*Token
	isValue     bool
}

var (
	ErrTokenIsNotValue = fmt.Errorf("token is not value")
	ErrTokenIsValue    = fmt.Errorf("token is value")
)

func NewValueToken(value string) *Token {
	return &Token{
		stringValue: value,
		isValue:     true,
	}
}

func NewTokensToken(tokens ...*Token) *Token {
	return &Token{
		tokens:  tokens,
		isValue: false,
	}
}

func (t *Token) MaybeValue() (value string, ok bool) {
	if t.isValue {
		value = t.stringValue
		ok = true
	}
	return
}

func (t *Token) Value() string {
	if value, ok := t.MaybeValue(); ok {
		return value
	}

	panic(ErrTokenIsNotValue)
}

func (t *Token) MaybeTokens() (tokens []*Token, ok bool) {
	if !t.isValue {
		tokens = t.tokens
		ok = true
	}
	return
}

func (t *Token) Tokens() []*Token {
	if tokens, ok := t.MaybeTokens(); ok {
		return tokens
	}

	panic(ErrTokenIsValue)
}
