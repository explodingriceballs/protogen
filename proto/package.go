package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Package struct {
	noopVisitor
	types *TypeDictionary

	*Options
	Name     string
	messages []*Message
	enums    []*Enum
	services []*Service
}

func (p *Package) DeclaresEnumType(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (p *Package) GetName() string {
	return p.Name
}

func (p *Package) GetType() ElementType {
	return PackageElementType
}

func (p *Package) VisitComment(comment *parser.Comment) {
	//TODO implement me
	panic("implement me")
}

func (p *Package) VisitEnum(enum *parser.Enum) (next bool) {
	// Create a new enum
	newEnum := NewEnum(p.types)

	// Start the scope & process & end
	enum.Accept(newEnum)

	// Finally append to the list of enums
	p.enums = append(p.enums, newEnum)
	return false
}

func (p *Package) VisitExtend(extend *parser.Extend) (next bool) {
	extendMessage := NewMessage(p.types)
	for _, body := range extend.ExtendBody {
		if field, ok := body.(*parser.Field); ok {
			extendMessage.VisitField(field)
		}
	}
	p.types.RegisterExtension(extend.MessageType, extendMessage)
	return false
}

func (p *Package) VisitMessage(message *parser.Message) (next bool) {
	// Create a new message, pass down the type dictionary
	newMessage := NewMessage(p.types)

	// Visit the message
	message.Accept(newMessage)

	// Finally append to the list of messages
	p.messages = append(p.messages, newMessage)
	return false
}

func (p *Package) VisitOption(option *parser.Option) (next bool) {
	if p.Options == nil {
		p.Options = NewOptions(p, p.types)
	}
	return p.Options.VisitOption(option)
}

func (p *Package) VisitService(service *parser.Service) (next bool) {
	newService := NewService(p.types)
	service.Accept(newService)
	p.services = append(p.services, newService)
	return false
}

func (p *Package) Message(name string) *Message {
	for _, msg := range p.messages {
		if msg.MessageName == name {
			return msg
		}
	}
	return nil
}

func (p *Package) Enum(name string) *Enum {
	for _, enum := range p.enums {
		if enum.EnumName == name {
			return enum
		}
	}
	return nil
}

func (p *Package) Extension(extType ExtensionType, pkgName string) []*Message {
	if extension, ok := p.types.extensions[extType]; ok {
		if extensionPkg, ok := extension[pkgName]; ok {
			return extensionPkg
		}
	}
	return nil
}

func (p *Package) Messages() []*Message {
	return p.messages
}

func (p *Package) Services() []*Service {
	return p.services
}

func (p *Package) Enums() []*Enum {
	return p.enums
}

func (p *Package) Extensions() map[ExtensionType]map[string][]*Message {
	return p.types.extensions
}

func NewPackage(name string, types *TypeDictionary) *Package {
	return &Package{
		types: types,
		Name:  name,
	}
}
