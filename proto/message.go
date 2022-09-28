package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Message struct {
	noopVisitor
	types *TypeDictionary

	*Options
	msgName  string
	oneOf    bool
	fields   []*Field
	messages []*Message
	enums    []*Enum
	msgType  *Type
}

func (m *Message) Name() string {
	return m.msgName
}

func (m *Message) Type() *Type {
	return m.msgType
}

func (m *Message) Enums() []*Enum {
	return m.enums
}

func (m *Message) Enum(name string) *Enum {
	for _, enum := range m.enums {
		if enum.Name() == name {
			return enum
		}
	}
	return nil
}

func (m *Message) GetElementType() ElementType {
	return MessageElementType
}

func (m *Message) IsOneOf() bool {
	return m.oneOf
}

func (m *Message) VisitMessage(message *parser.Message) (next bool) {
	// If the message name is already set, we are looking at a nested message
	if m.msgName != "" {
		// New a new message
		newMsg := NewMessage(m.types)

		// Visit & append to nested messages
		message.Accept(newMsg)
		m.messages = append(m.messages, newMsg)

		return false
	}

	// Set the message name
	m.msgName = message.MessageName

	// Declare that we are starting a new scope with this message & register the internal messages
	m.types.startScope(m)
	m.types.declareTypes(message.MessageBody)

	// Register ourselves as a type
	m.msgType = m.types.RegisterMessage(m)

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

func (m *Message) VisitEnum(enum *parser.Enum) (next bool) {
	newEnum := NewEnum(m.types)
	enum.Accept(newEnum)
	//m.types.RegisterEnum(newEnum)
	m.enums = append(m.enums, newEnum)
	return false
}

func (m *Message) VisitOneof(oneof *parser.Oneof) (next bool) {
	// create a new field for the one of
	oneOfField := NewField(m.types)
	oneOfField.fieldName = oneof.OneofName

	// TODO: FIX ME!
	oneOfMsg := NewMessage(m.types)
	oneOfMsg.oneOf = true
	oneOfMsg.msgName = oneof.OneofName
	for _, oneOfField := range oneof.OneofFields {
		oneOfMsg.fields = append(oneOfMsg.fields, &Field{
			types:       m.types,
			fieldName:   oneOfField.FieldName,
			fieldNumber: oneOfField.FieldNumber,
			t:           m.types.findType(oneOfField.Type),
			Options:     &Options{},
		})
	}

	// Register & assign
	m.types.startScope(oneOfMsg)
	m.types.RegisterMessage(oneOfMsg)
	m.types.endScope(oneOfMsg)
	oneOfField.t = m.types.findType(oneOfMsg.msgName)
	m.fields = append(m.fields, oneOfField)
	return false
}

func (m *Message) VisitField(field *parser.Field) (next bool) {
	msgField := NewField(m.types)
	field.Accept(msgField)
	m.fields = append(m.fields, msgField)
	return false
}

func (m *Message) VisitOption(option *parser.Option) (next bool) {
	if m.Options == nil {
		m.Options = NewOptions(m, m.types)
	}
	return m.Options.VisitOption(option)
}

func (m *Message) Fields() []*Field {
	return m.fields
}

func (m *Message) Field(name string) *Field {
	for _, f := range m.fields {
		if f.fieldName == name {
			return f
		}
	}
	return nil
}

func (m *Message) Messages() []*Message {
	return m.messages
}

func (m *Message) Message(name string) *Message {
	for _, msg := range m.messages {
		if msg.msgName == name {
			return msg
		}
	}
	return nil
}

func NewMessage(types *TypeDictionary) *Message {
	m := &Message{
		types: types,
	}
	m.Options = NewOptions(m, types)
	return m
}
