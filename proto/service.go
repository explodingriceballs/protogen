package proto

import "github.com/yoheimuta/go-protoparser/v4/parser"

type Service struct {
	noopVisitor
	types *TypeDictionary

	*Options
	serviceName string
	rpcs        []*RPC
}

func (s *Service) GetElementType() ElementType {
	return ServiceElementType
}

func (s *Service) VisitService(service *parser.Service) (next bool) {
	s.serviceName = service.ServiceName
	return true
}

func (s *Service) VisitRPC(rpc *parser.RPC) (next bool) {
	serviceRpc := NewRPC(s.types)
	rpc.Accept(serviceRpc)
	s.rpcs = append(s.rpcs, serviceRpc)
	return false
}

func (s *Service) VisitOption(option *parser.Option) (next bool) {
	s.AddOption(option.OptionName, option.Constant)

	// Ignore comments
	return false
}

func (s *Service) Name() string {
	return s.serviceName
}

func (s *Service) RPC(name string) *RPC {
	for _, r := range s.rpcs {
		if r.RPCName == name {
			return r
		}
	}
	return nil
}

func (s *Service) RPCs() []*RPC {
	return s.rpcs
}

func NewService(types *TypeDictionary) *Service {
	s := &Service{
		types: types,
	}
	s.Options = NewOptions(s, types)
	return s
}
