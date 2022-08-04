package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Message struct {
	noopVisitor
	types *TypeDictionary

	MessageName string
	OneOf       bool
	Fields      []*Field
	Options     []*Option
	Messages    []*Message
}

func (m *Message) VisitMessage(message *parser.Message) (next bool) {
	if m.MessageName != "" {
		m.types.startScope(message.MessageName)
		newMsg := NewMessage(m.types)
		message.Accept(newMsg)
		m.Messages = append(m.Messages, newMsg)
		m.types.endScope(message.MessageName)
		return false
	}
	m.MessageName = message.MessageName
	m.types.RegisterMessage(m)

	// Process definitions first
	for _, body := range message.MessageBody {
		if _, ok := body.(*parser.Message); ok {
			body.Accept(m)
		}
		if _, ok := body.(*parser.Enum); ok {
			body.Accept(m)
		}
	}

	// Process the rest
	for _, body := range message.MessageBody {
		if _, ok := body.(*parser.Message); ok {
			continue
		}
		if _, ok := body.(*parser.Enum); ok {
			continue
		}
		body.Accept(m)
	}

	return false
}

func (m *Message) VisitOneof(oneof *parser.Oneof) (next bool) {
	// create a new field for the one of
	oneOfField := NewField(m.types)
	oneOfField.FieldName = oneof.OneofName

	// TODO: FIX ME!
	oneOfMsg := NewMessage(m.types)
	oneOfMsg.OneOf = true
	oneOfMsg.MessageName = oneof.OneofName
	for _, oneOfField := range oneof.OneofFields {
		oneOfMsg.Fields = append(oneOfMsg.Fields, &Field{
			types:       m.types,
			FieldName:   oneOfField.FieldName,
			FieldNumber: oneOfField.FieldNumber,
			Type:        m.types.findType(oneOfField.Type),
		})
	}

	// Register & assign
	m.types.startScope(oneOfMsg.MessageName)
	m.types.RegisterMessage(oneOfMsg)
	m.types.endScope(oneOfMsg.MessageName)
	oneOfField.Type = m.types.findType(oneOfMsg.MessageName)
	m.Fields = append(m.Fields, oneOfField)
	return false
}

func (m *Message) VisitField(field *parser.Field) (next bool) {
	msgField := NewField(m.types)
	field.Accept(msgField)
	m.Fields = append(m.Fields, msgField)
	return false
}

func (m *Message) VisitOption(option *parser.Option) (next bool) {
	// Create a new option
	newOption := NewOption()
	option.Accept(newOption)
	m.Options = append(m.Options, newOption)

	// Skip comments
	return false
}

func (m *Message) Field(name string) *Field {
	for _, f := range m.Fields {
		if f.FieldName == name {
			return f
		}
	}
	return nil
}

func (m *Message) Message(name string) *Message {
	for _, msg := range m.Messages {
		if msg.MessageName == name {
			return msg
		}
	}
	return nil
}

func NewMessage(types *TypeDictionary) *Message {
	m := &Message{
		types: types,
	}
	return m
}
