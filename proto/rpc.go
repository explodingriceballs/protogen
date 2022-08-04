package proto

import "github.com/yoheimuta/go-protoparser/v4/parser"

type RPC struct {
	noopVisitor
	types *TypeDictionary

	RPCName             string
	RequestMessageType  *Type
	ResponseMessageType *Type
}

func (r *RPC) VisitRPC(rpc *parser.RPC) (next bool) {
	r.RPCName = rpc.RPCName
	r.RequestMessageType = r.types.findType(rpc.RPCRequest.MessageType)
	r.ResponseMessageType = r.types.findType(rpc.RPCResponse.MessageType)
	return true
}

func NewRPC(types *TypeDictionary) *RPC {
	return &RPC{types: types}
}
