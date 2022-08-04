package main

//
//type Option struct {
//	options map[string]string
//}
//
//func (o *Option) AddOption(name string, constant string) {
//	o.options[name] = constant
//}
//
//type Package struct {
//	PackageName string
//	Option
//	Messages
//	state *State
//}
//
//func (p *Package) AddMessage(name string) *Message {
//	message := p.Messages.AddMessage(name)
//	p.state.addMessageToLookup(p.PackageName+"."+name, message)
//	return message
//}
//
//type Field struct {
//	Option
//	FieldName   string
//	FieldNumber string
//	IsRepeated  bool
//}
//
//type Message struct {
//	MessageName string
//	Option
//}
//
//type Messages struct {
//	messages map[string]*Message
//}
//
//func (m *Messages) AddMessage(name string) *Message {
//	msg := &Message{
//		MessageName: name,
//		Option:      Option{},
//	}
//	m.messages[name] = msg
//	return msg
//}
//
//type State struct {
//	Option
//	Messages
//	packages      map[string]*Package
//	messageLookup map[string]*Message
//}
//
//func (s *State) AddPackage(name string) *Package {
//	p := &Package{
//		state:       s,
//		PackageName: name,
//		Option: Option{
//			options: map[string]string{},
//		},
//		Messages: Messages{
//			messages: map[string]*Message{},
//		},
//	}
//	s.packages[name] = p
//	return p
//}
//
//func (s *State) AddMessage(name string) *Message {
//	message := s.Messages.AddMessage(name)
//	s.addMessageToLookup(name, message)
//	return message
//}
//
//func (s *State) addMessageToLookup(name string, msg *Message) {
//	s.messageLookup[name] = msg
//}
//
//func (s *State) findType(t string) {
//
//}
