package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Message struct {
	noopVisitor
	types *TypeDictionary

	*Options
	MessageName string
	OneOf       bool
	Fields      []*Field
	Messages    []*Message
}

func (m *Message) DeclaresMessageType(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (m *Message) DeclaresEnumType(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (m *Message) GetName() string {
	return m.MessageName
}

func (m *Message) GetType() ElementType {
	return MessageElementType
}

func (m *Message) VisitMessage(message *parser.Message) (next bool) {
	// If the message name is already set, we are looking at a nested message
	if m.MessageName != "" {
		// Create a new message
		newMsg := NewMessage(m.types)

		// Visit & append to nested messages
		message.Accept(newMsg)
		m.Messages = append(m.Messages, newMsg)

		return false
	}

	// Set the message name
	m.MessageName = message.MessageName

	// Declare that we are starting a new scope with this message
	m.types.startScope(m)

	// Register ourself as a type
	m.types.RegisterMessage(m)

	// Process definitions first
	for _, body := range message.MessageBody {
		switch body.(type) {
		case *parser.Message:
			body.Accept(m)
		case *parser.Enum:
			body.Accept(m)
		}
	}

	// Process the rest
	for _, body := range message.MessageBody {
		switch body.(type) {
		case *parser.Message:
			continue
		case *parser.Enum:
			continue
		}
		body.Accept(m)
	}

	// End our scope
	m.types.endScope(m)

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
	m.types.startScope(oneOfMsg)
	m.types.RegisterMessage(oneOfMsg)
	m.types.endScope(oneOfMsg)
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
	if m.Options == nil {
		m.Options = NewOptions(m, m.types)
	}
	return m.Options.VisitOption(option)
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
