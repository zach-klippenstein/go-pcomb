package pcomb

import (
	"fmt"
	"strings"
)

type ParseResult struct {
	Token     *Token
	Err       error
	NextInput Input
}

type ParserFunc func(Input) ParseResult

type Parser interface {
	Parse(Input) ParseResult
	String() string
}

var Noop = NamedParser{"noop", func(i Input) ParseResult {
	return ParseResult{
		NextInput: i,
	}
}}

type NamedParser struct {
	Name   string
	Parser ParserFunc
}

func (p NamedParser) Parse(i Input) ParseResult {
	return p.Parser(i)
}

func (p NamedParser) String() string {
	return p.Name
}

func String(value string) Parser {
	return NamedParser{fmt.Sprintf("%q", value), func(i Input) ParseResult {
		if strings.HasPrefix(i.Value(), value) {
			return ParseResult{
				Token:     NewValueToken(value),
				NextInput: i.Advance(len(value)),
			}
		}

		return ParseResult{
			Err:       fmt.Errorf("input doesn't match ‘%s’ at %v", value, i),
			NextInput: i,
		}
	}}
}

func Ignore(parser Parser) Parser {
	return NamedParser{fmt.Sprintf("ignore(%s)", parser), func(input Input) ParseResult {
		result := parser.Parse(input)

		if result.Err != nil {
			return result
		}

		return ParseResult{
			NextInput: result.NextInput,
		}
	}}
}

func Sequence(parsers ...Parser) Parser {
	return NamedParser{fmt.Sprintf("(%s)", parsers), func(input Input) ParseResult {
		tokens := make([]*Token, len(parsers))

		for i := 0; i < len(parsers); i++ {
			result := parsers[i].Parse(input)

			if result.Err != nil {
				return result
			}

			tokens[i] = result.Token
			input = result.NextInput
		}

		return ParseResult{
			Token:     NewTokensToken(tokens...),
			NextInput: input,
		}
	}}
}

func AnyOf(parsers ...Parser) Parser {
	return NamedParser{fmt.Sprintf("any(%s)", parsers), func(input Input) ParseResult {
		for _, parser := range parsers {
			result := parser.Parse(input)

			if result.Err == nil {
				return result
			}
		}

		return ParseResult{
			Err:       fmt.Errorf("input doesn't match any of %+v at %v", parsers, input),
			NextInput: input,
		}
	}}
}
