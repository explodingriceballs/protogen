package main

import (
	"errors"
)

type state uint

const (
	stateUnknown = iota
	stateInPackage
	stateInService
	stateInRPC
	stateInMessage
	stateInOneof
	stateInExtension
	stateInField
	stateInEnum
	stateInEnumField
	stateInOption
)

func (s state) String() string {
	switch s {
	case stateUnknown:
		return "unknown"
	case stateInPackage:
		return "package"
	case stateInService:
		return "service"
	case stateInRPC:
		return "rpc"
	case stateInMessage:
		return "message"
	case stateInOneof:
		return "oneof"
	case stateInExtension:
		return "extension"
	case stateInField:
		return "field"
	case stateInEnum:
		return "enum"
	case stateInEnumField:
		return "field"
	case stateInOption:
		return "option"
	default:
		return "unkown"
	}
}

type stateRef struct {
	state state
	ref   interface{}
}

type ParserState struct {
	pkg   *Package
	state []stateRef
}

func NewParserState() *ParserState {
	return &ParserState{
		pkg:   nil,
		state: []stateRef{},
	}
}

func (p *ParserState) StartPackage(pkg *Package) {
	p.pushState(stateInPackage, pkg)
	p.pkg = pkg
}

// StartMessage sets the internal state & adds the message to the relevant reference
func (p *ParserState) StartMessage(msg *Message) error {
	switch p.curState() {
	// Global message
	case stateInPackage:
		pkg := p.curStateRef().(*Package)
		pkg.Messages = append(pkg.Messages, msg)
		p.pushState(stateInMessage, msg)
	case stateInMessage:
		curMsg := p.curStateRef().(*Message)
		curMsg.NestedMessages = append(curMsg.NestedMessages, msg)
		p.pushState(stateInMessage, msg)
	default:
		return InvalidStateError("start message", p.curState())
	}
	return nil
}

func (p *ParserState) StartOneOf(oneOf *OneOf) error {
	if p.curState() != stateInMessage {
		return InvalidStateError("finalize message", p.curState())
	}

	message := p.curStateRef().(*Message)
	message.OneOfs = append(message.OneOfs, oneOf)
	p.pushState(stateInOneof, oneOf)

	return nil
}

func (p *ParserState) AddOneOfField(field *Field) error {
	if p.curState() != stateInOneof {
		return InvalidStateError("finalize oneof", p.curState())
	}
	oneOfMsg := p.curStateRef().(*OneOf)
	oneOfMsg.Fields = append(oneOfMsg.Fields, field)
	return nil
}

func (p *ParserState) FinalizeOneOf() error {
	if p.curState() != stateInOneof {
		return InvalidStateError("finalize oneof", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) FinalizeMessage() error {
	if p.curState() != stateInMessage {
		return InvalidStateError("finalize message", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) StartExtension(message *Message) error {
	if p.curState() != stateInPackage {
		return InvalidStateError("start extension", p.curState())
	}
	pkg := p.curStateRef().(*Package)
	pkg.Extensions = append(pkg.Extensions, message)
	p.pushState(stateInExtension, message)
	return nil
}

func (p *ParserState) FinalizeExtension() error {
	if p.curState() != stateInExtension {
		return InvalidStateError("finalize extension", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) StartService(service *Service) error {
	switch p.curState() {
	case stateInPackage:
		pkg := p.curStateRef().(*Package)
		pkg.Services = append(pkg.Services, service)
		p.pushState(stateInService, service)
		return nil
	default:
		return InvalidStateError("start service", p.curState())
	}
}

func (p *ParserState) FinalizeService() error {
	if p.curState() != stateInService {
		return InvalidStateError("finalize service", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) StartRPC(rpc *RPC) error {
	if p.curState() != stateInService {
		return InvalidStateError("start rpc", p.curState())
	}
	service := p.curStateRef().(*Service)
	service.RPCs = append(service.RPCs, rpc)
	p.pushState(stateInRPC, rpc)
	return nil
}

func (p *ParserState) FinalizeRPC() error {
	if p.curState() != stateInRPC {
		return InvalidStateError("finalize rpc", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) StartField(field *Field) error {
	if p.curState() != stateInMessage && p.curState() != stateInExtension {
		return InvalidStateError("start field", p.curState())
	}
	msg := p.curStateRef().(*Message)
	msg.Fields = append(msg.Fields, field)
	p.pushState(stateInField, field)
	return nil
}

func (p *ParserState) FinalizeField() error {
	if p.curState() != stateInField {
		return InvalidStateError("finalize field", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) StartOption(option *Option) error {
	switch p.curState() {
	case stateInPackage:
		pkg := p.curStateRef().(*Package)
		pkg.Options = append(pkg.Options, option)
	case stateInField:
		field := p.curStateRef().(*Field)
		field.Options = append(field.Options, option)
	case stateInEnum:
		enum := p.curStateRef().(*Enum)
		enum.Options = append(enum.Options, option)
	case stateInEnumField:
		enumField := p.curStateRef().(*EnumValue)
		enumField.Options = append(enumField.Options, option)
	case stateInRPC:
		rpc := p.curStateRef().(*RPC)
		rpc.Options = append(rpc.Options, option)
	case stateInOption:
		stateOpt := p.curStateRef().(*Option)
		stateOpt.Properties[option.Key] = option
	case stateInMessage:
		message := p.curStateRef().(*Message)
		message.Options = append(message.Options, option)
	default:
		panic("NYI")
	}
	p.pushState(stateInOption, option)
	return nil
}

func (p *ParserState) SetOptionValue(s string) error {
	if p.curState() != stateInOption {
		return InvalidStateError("set option value", p.curState())
	}
	option := p.curStateRef().(*Option)
	option.Value = s

	return nil
}

func (p *ParserState) FinalizeOption() error {
	if p.curState() != stateInOption {
		return InvalidStateError("finalize option", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) StartEnum(enum *Enum) error {
	switch p.curState() {
	case stateInPackage:
		pkg := p.curStateRef().(*Package)
		pkg.Enums = append(pkg.Enums, enum)
	case stateInMessage:
		msg := p.curStateRef().(*Message)
		msg.Enums = append(msg.Enums, enum)
	default:
		return InvalidStateError("start enum", p.curState())
	}
	p.pushState(stateInEnum, enum)
	return nil
}

func (p *ParserState) FinalizeEnum() error {
	if p.curState() != stateInEnum {
		return InvalidStateError("finalize enum", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) StartEnumValue(enumValue *EnumValue) error {
	if p.curState() != stateInEnum {
		return InvalidStateError("start enum value", p.curState())
	}

	enum := p.curStateRef().(*Enum)
	enum.Values = append(enum.Values, enumValue)

	p.pushState(stateInEnumField, enumValue)
	return nil
}

func (p *ParserState) FinalizeEnumValue() error {
	if p.curState() != stateInEnumField {
		return InvalidStateError("finalize enum value", p.curState())
	}
	p.popState()
	return nil
}

func (p *ParserState) pushState(newState state, ref interface{}) {
	p.state = append([]stateRef{
		{
			state: newState,
			ref:   ref,
		},
	}, p.state...)
}

func (p *ParserState) popState() {
	p.state = p.state[1:len(p.state)]
}

func (p *ParserState) curState() state {
	return p.state[0].state
}

func (p *ParserState) curStateRef() interface{} {
	return p.state[0].ref
}

func InvalidStateError(s string, curState state) error {
	return errors.New("invalid state: " + s + " - " + curState.String())
}
