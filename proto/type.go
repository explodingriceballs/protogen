package proto

import (
	"github.com/rs/zerolog/log"
	"strings"
)

type ExtensionType int

const (
	UnknownOption ExtensionType = iota
	FileOption
	MessageOption
	FieldOption
	OneofOption
	EnumOption
	EnumValueOption
	ServiceOption
	MethodOption
)

func extensionTypeFromString(name string) ExtensionType {
	actualOptionName := name
	if strings.Contains(name, ".") {
		partsOfName := strings.Split(name, ".")
		actualOptionName = partsOfName[len(partsOfName)-1]
	}
	switch actualOptionName {
	case "FileOptions":
		return FileOption
	case "MessageOptions":
		return MessageOption
	case "FieldOptions":
		return FieldOption
	case "OneofOptions":
		return OneofOption
	case "EnumOptions":
		return EnumOption
	case "EnumValueOptions":
		return EnumValueOption
	case "ServiceOptions":
		return ServiceOption
	case "MethodOptions":
		return MethodOption
	default:
		log.Warn().Str("option", name).Msg("option unrecognized. not parsing.")
		return UnknownOption
	}
}

var nativeTypes = map[string]*Type{
	"double":   {isNative: true, nativeType: "double"},
	"float":    {isNative: true, nativeType: "float"},
	"int32":    {isNative: true, nativeType: "int32"},
	"int64":    {isNative: true, nativeType: "int64"},
	"uint32":   {isNative: true, nativeType: "uint32"},
	"uint64":   {isNative: true, nativeType: "uint64"},
	"sint32":   {isNative: true, nativeType: "sint32"},
	"sint64":   {isNative: true, nativeType: "sint64"},
	"fixed32":  {isNative: true, nativeType: "fixed32"},
	"fixed64":  {isNative: true, nativeType: "fixed64"},
	"sfixed32": {isNative: true, nativeType: "sfixed32"},
	"sfixed64": {isNative: true, nativeType: "sfixed64"},
	"bool":     {isNative: true, nativeType: "bool"},
	"string":   {isNative: true, nativeType: "string"},
	"bytes":    {isNative: true, nativeType: "bytes"},
}

type Type struct {
	isNative   bool
	nativeType string

	referenceType string
	isMessageType bool
	isEnumType    bool
	message       *Message
	enum          *Enum
}

type TypeDictionary struct {
	currentScope []string
	types        map[string]*Type
	extensions   map[ExtensionType]map[string][]*Message
}

func (d *TypeDictionary) startScope(name string) {
	d.currentScope = append(d.currentScope, name)
	log.Debug().Str("startScope", name).Str("scope", d.getScope()).Msg("Starting scope")
}

func (d *TypeDictionary) endScope(name string) {
	elementIdx := -1
	for idx, scopeName := range d.currentScope {
		if scopeName == name {
			elementIdx = idx
		}
	}
	if elementIdx == -1 {
		panic("unexpected scope end: " + name)
	}
	d.currentScope = append(d.currentScope[:elementIdx], d.currentScope[elementIdx+1:]...)
	log.Debug().Str("endScope", name).Str("scope", d.getScope()).Msg("Ending scope")

}

func (d *TypeDictionary) getScope() string {
	return strings.Join(d.currentScope, ".")
}

func (d *TypeDictionary) RegisterExtension(extends string, msg *Message) {
	extensionType := extensionTypeFromString(extends)
	if _, ok := d.extensions[extensionType]; !ok {
		d.extensions[extensionType] = map[string][]*Message{}
	}
	d.extensions[extensionType][d.currentScope[0]] = append(d.extensions[extensionType][d.currentScope[0]], msg)
}

func (d *TypeDictionary) RegisterMessage(m *Message) {
	if existingType, ok := d.types[d.getScope()]; ok {
		existingType.isNative = false
		existingType.nativeType = ""
		existingType.isMessageType = true
		existingType.isEnumType = false
		existingType.message = m
		existingType.enum = nil
		return
	}
	d.types[d.getScope()] = &Type{
		isNative:      false,
		nativeType:    "",
		referenceType: m.MessageName,
		isMessageType: true,
		isEnumType:    false,
		message:       m,
		enum:          nil,
	}
}

func (d *TypeDictionary) RegisterEnum(enum *Enum) {
	if existingType, ok := d.types[d.getScope()]; ok {
		existingType.isNative = false
		existingType.nativeType = ""
		existingType.isMessageType = false
		existingType.isEnumType = true
		existingType.message = nil
		existingType.enum = enum
		return
	}
	d.types[d.getScope()] = &Type{
		isNative:      false,
		nativeType:    "",
		referenceType: enum.EnumName,
		isMessageType: false,
		isEnumType:    true,
		message:       nil,
		enum:          enum,
	}
}

func (d *TypeDictionary) findType(t string) *Type {
	if nativeType, ok := nativeTypes[t]; ok {
		return nativeType
	}

	// Check for the full qualified identifier
	lookupType := strings.Join(append(d.currentScope, t), ".")
	if configuredType, ok := d.types[lookupType]; ok {
		return configuredType
	}

	lookupType = t
	// Scope to global if the identifier contains no dots
	if !strings.Contains(t, ".") && d.currentScope[0] == "" {
		lookupType = "." + lookupType
	}

	if configuredType, ok := d.types[lookupType]; ok {
		return configuredType
	}

	// Loop thru all scopes
	for idx := len(d.currentScope) - 1; idx >= 0; idx-- {
		if d.currentScope[idx] == t {
			return d.types[d.getScope()]
		}
		searchScopeKeys := make([]string, len(d.currentScope[0:idx+1]))
		copy(searchScopeKeys, d.currentScope[0:idx+1])
		searchScope := append(searchScopeKeys, t)
		searchScopeStr := strings.Join(searchScope, ".")
		if foundType, ok := d.types[searchScopeStr]; ok {
			return foundType
		}
	}

	// Self reference
	if d.currentScope[len(d.currentScope)-1] == t {
		return d.types[d.getScope()]
	}

	newType := &Type{
		referenceType: t,
		isMessageType: false,
		isEnumType:    false,
		message:       nil,
		enum:          nil,
	}
	d.types[lookupType] = newType

	return newType
}

func (d *TypeDictionary) newMsgTypeForTesting(msg *Message) *Type {
	return &Type{
		isNative:      false,
		nativeType:    "",
		referenceType: msg.MessageName,
		isMessageType: true,
		isEnumType:    false,
		message:       msg,
		enum:          nil,
	}
}

func (d *TypeDictionary) newEnumTypeForTesting(enum *Enum) *Type {
	return &Type{
		isNative:      false,
		nativeType:    "",
		referenceType: enum.EnumName,
		isMessageType: false,
		isEnumType:    true,
		message:       nil,
		enum:          enum,
	}
}

func NewTypeDictionary() *TypeDictionary {
	return &TypeDictionary{
		currentScope: []string{},
		types:        map[string]*Type{},
		extensions:   map[ExtensionType]map[string][]*Message{},
	}
}
