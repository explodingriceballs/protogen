package proto

import (
	"github.com/bbuck/go-lexer"
	"github.com/rs/zerolog/log"
)

const (
	CurlyLeftBracket   lexer.TokenType = iota // {
	CurlyRightBracket                         // }
	SquareLeftBracket                         // [
	SquareRightBracket                        // ]
	SemiColon                                 // ;
	Colon                                     // :
	Ident
	StringValue
	Value
)

func TokenToString(token lexer.TokenType) string {
	switch token {
	case CurlyLeftBracket:
		return "CurlyLeftBracket"
	case CurlyRightBracket:
		return "CurlyRightBracket"
	case SquareLeftBracket:
		return "SquareLeftBracket"
	case SquareRightBracket:
		return "SquareRightBracket"
	case SemiColon:
		return "SemiColon"
	case Colon:
		return "Colon"
	case Ident:
		return "Ident"
	case StringValue:
		return "StringValue"
	case Value:
		return "Value"
	}
	return "unknown"
}

func ReadArrayEnd(l *lexer.L) lexer.StateFunc {
	l.Next()
	if l.Current() != "]" {
		l.Error("expected curly bracket but got: '" + l.Current() + "'")
		return nil
	}
	l.Emit(SquareRightBracket)
	return ReadMessageEnd
}

func ReadArrayStart(l *lexer.L) lexer.StateFunc {
	l.Next()
	if l.Current() != "[" {
		l.Error("expected curly bracket but got: '" + l.Current() + "'")
		return nil
	}
	l.Emit(SquareLeftBracket)

	if l.Peek() == '{' {
		return ReadMessageStart
	}

	return ReadValue
}

func ReadValueEnd(l *lexer.L) lexer.StateFunc {
	// Read a comma or semicolon
	if l.Peek() == ',' || l.Peek() == ';' {
		l.Next()
		l.Ignore()
		if l.Peek() == '}' {
			return ReadValueEnd
		}
		return ReadIdentifier
	}

	if l.Peek() == '}' {
		return ReadMessageEnd
	}

	return nil
	//l.Next()
	//
	//peek := string(l.Peek())
	//fmt.Println(peek + " - " + l.Current())
	//
	//current := l.Current()
	//
	//// Check
	//if current == "}" {
	//	l.Emit(CurlyRightBracket)
	//	// Check if this is the last curly bracket
	//	if l.Peek() == lexer.EOFRune {
	//		return nil
	//	}
	//
	//	if l.Peek() == ',' || l.Peek() == ';' {
	//		return ReadValueEnd
	//	}
	//
	//	// Read another identifier
	//	return ReadIdentifier
	//}
	//if current == "," || current == ";" {
	//	l.Emit(SemiColon)
	//}
	//if l.Peek() == '{' {
	//	return ReadMessageStart
	//}
	//return ReadIdentifier
}

func ReadValue(l *lexer.L) lexer.StateFunc {
	l.Next()
	// String
	if l.Current() == "\"" {
		l.Ignore()
		for {
			peekOne := l.Next()
			peekTwo := l.Next()
			l.Rewind()
			l.Rewind()
			if peekOne != '\\' && peekTwo == '"' {
				// Advance by one
				l.Next()
				break
			}
			l.Next()
		}
		l.Emit(StringValue)
		l.Next()
		l.Ignore()
		return ReadValueEnd
	}

	// Others
	for {
		peek := l.Peek()
		if peek != ',' && peek != ';' && peek != '}' {
			l.Next()
		} else {
			break
		}
	}
	l.Emit(Value)

	return ReadValueEnd
}

func ReadColon(l *lexer.L) lexer.StateFunc {
	l.Next()
	if l.Current() != ":" {
		l.Error("expected curly bracket but got: '" + l.Current() + "'")
		return nil
	}
	l.Emit(Colon)

	if l.Peek() == '[' {
		return ReadArrayStart
	}

	return ReadValue
}

func ReadIdentifier(l *lexer.L) lexer.StateFunc {
	l.Next()
	for l.Peek() != ':' {
		l.Next()
	}
	l.Emit(Ident)
	return ReadColon
}

func ReadMessageEnd(l *lexer.L) lexer.StateFunc {
	l.Next()
	l.Emit(CurlyRightBracket)

	switch l.Peek() {
	case ',':
		l.Emit(SemiColon)
		l.Next()
		l.Ignore()
		if l.Peek() == '{' {
			return ReadMessageStart
		}
	case '{':
		return ReadMessageStart
	case ']':
		return ReadArrayEnd
	case lexer.EOFRune:
		return nil
	}
	log.Info().Msg("Unsure peek: " + string(l.Peek()))
	return nil
}

func ReadMessageStart(l *lexer.L) lexer.StateFunc {
	// Advance
	l.Next()

	// If not a bracket, break
	if l.Current() != "{" {
		l.Error("expected curly bracket but got: '" + l.Current() + "'")
		return nil
	}
	l.Emit(CurlyLeftBracket)
	return ReadIdentifier
}

func ParseOption(value string, t *Type) interface{} {
	// Skip the easy ones
	if t.isNative {
		return toNativeValue(value, t)
	}
	if t.isEnumType {
		return toEnumValue(value, t)
	}

	if !t.isMessageType {
		panic("unknown type handed for option parsing: " + t.referenceType)
	}

	l := lexer.New(value, ReadMessageStart)
	l.Start()

	builder := newOptionBuilder(true, t)
	for {
		token, done := l.NextToken()
		if done {
			return builder.result()
		}
		//fmt.Printf("[%s]: %s\n", TokenToString(token.t), token.Value)
		builder.accept(token)
	}
}

type optionBuilder struct {
	values       map[string]interface{}
	arrayValues  []interface{}
	t            *Type
	ident        string
	firstBracket bool
	inArray      bool
	done         bool
	builder      *optionBuilder
}

func (b *optionBuilder) accept(token *lexer.Token) {
	if b.builder != nil {
		b.builder.accept(token)
		if b.builder.isDone() {
			if b.inArray {
				b.arrayValues = append(b.arrayValues, b.builder.result())
			} else {
				b.values[b.ident] = b.builder.result()
			}
			b.builder = nil
		}
		return
	}

	switch token.Type {
	case CurlyLeftBracket:
		// Read over the first bracket
		if b.firstBracket {
			b.firstBracket = false
			return
		}

		field := b.t.message.Field(b.ident)
		if b.inArray {
			if !field.isRepeated {
				log.Warn().Str("name", field.fieldName).Msg("field is not repeated but found array")
				return
			}
			b.builder = newOptionBuilder(false, field.t)
			return
		}
	case CurlyRightBracket:
		b.done = true
		return
	case SquareLeftBracket:
		b.inArray = true
		return
	case SquareRightBracket:
		b.inArray = false
		b.values[b.ident] = b.arrayValues
		return
	case Ident:
		b.ident = token.Value
		return
	case Colon:
		if b.ident == "" {
			log.Warn().Msg("option value without a key")
		}
		return
	case SemiColon:
		if !b.inArray {
			b.ident = ""
		}
		return
	case StringValue:
		b.values[b.ident] = token.Value
	case Value:
		field := b.t.message.Field(b.ident)
		if field == nil {
			log.Warn().Str("field", b.ident).Msg("field in options not found")
			return
		}
		if field.t.isNative {
			b.values[b.ident] = toNativeValue(token.Value, field.t)
			return
		}
		if field.t.isEnumType {
			b.values[b.ident] = toEnumValue(token.Value, field.t)
			return
		}
	}
	//log.Info().Msgf("%#v", token)
}

func (b *optionBuilder) isDone() bool {
	return b.done
}

func (b *optionBuilder) result() map[string]interface{} {
	return b.values
}

func newOptionBuilder(topLevel bool, t *Type) *optionBuilder {
	return &optionBuilder{
		values:       map[string]interface{}{},
		t:            t,
		firstBracket: topLevel,
	}
}
