package main

type Import struct {
	File   string
	Weak   bool
	Public bool
}

func NewImport(file string, weak bool, public bool) *Import {
	return &Import{File: file, Weak: weak, Public: public}
}

type Package struct {
	Name       string
	Services   []*Service
	Messages   []*Message
	Enums      []*Enum
	Extensions []*Message
	Options    []*Option
}

func NewPackage(name string) *Package {
	return &Package{Name: name}
}

type Service struct {
	Name    string
	RPCs    []*RPC
	Options []*Option
}

func NewService(name string) *Service {
	return &Service{Name: name}
}

type RPC struct {
	Name                string
	RequestMsg          string
	ResponseMsg         string
	IsRequestStreaming  bool
	IsResponseStreaming bool
	Options             []*Option
}

func NewRPC(name string, requestMsg string, responseMsg string) *RPC {
	return &RPC{Name: name, RequestMsg: requestMsg, ResponseMsg: responseMsg}
}

type Option struct {
	Key        string
	Value      string
	Properties map[string]interface{}
}

func NewOption(key string) *Option {
	return &Option{Key: key, Properties: map[string]interface{}{}}
}

type Message struct {
	Name           string
	Options        []*Option
	Fields         []*Field
	Enums          []*Enum
	NestedMessages []*Message
	OneOfs         []*OneOf
}

func NewMessage(name string) *Message {
	return &Message{Name: name}
}

type OneOf struct {
	Name   string
	Fields []*Field
}

func NewOneOf(name string) *OneOf {
	return &OneOf{Name: name}
}

type Field struct {
	Name        string
	Type        string
	FieldNumber string
	IsRepeated  bool
	Options     []*Option
}

func NewField(fieldName string, type_ string) *Field {
	return &Field{
		Name: fieldName,
		Type: type_,
	}
}

type Enum struct {
	Name    string
	Options []*Option
	Values  []*EnumValue
}

func NewEnum(name string) *Enum {
	return &Enum{Name: name}
}

type EnumValue struct {
	Name    string
	Value   string
	Options []*Option
}

func NewEnumValue(name string, value string) *EnumValue {
	return &EnumValue{Name: name, Value: value}
}

type File struct {
	Name     string
	Contents string
}
