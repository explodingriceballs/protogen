package proto

import "github.com/yoheimuta/go-protoparser/v4/parser"

type Service struct {
	noopVisitor
	types *TypeDictionary

	ServiceName string
	RPCs        []*RPC
}

func (s *Service) VisitService(service *parser.Service) (next bool) {
	s.ServiceName = service.ServiceName
	return true
}

func (s *Service) VisitRPC(rpc *parser.RPC) (next bool) {
	serviceRpc := NewRPC(s.types)
	rpc.Accept(serviceRpc)
	s.RPCs = append(s.RPCs, serviceRpc)
	return false
}

func NewService(types *TypeDictionary) *Service {
	return &Service{
		types: types,
	}
}
