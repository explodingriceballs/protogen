package proto

import "github.com/yoheimuta/go-protoparser/v4/parser"

type RPC struct {
	noopVisitor
	types *TypeDictionary

	*Options
	RPCName                  string
	RequestMessageType       *Type
	ResponseMessageType      *Type
	RequestMessageStreaming  bool
	ResponseMessageStreaming bool
}

func (r *RPC) GetElementType() ElementType {
	return RPCElementType
}

func (r *RPC) Name() string {
	return r.RPCName
}

func (r *RPC) VisitOption(option *parser.Option) (next bool) {
	r.AddOption(option.OptionName, option.Constant)

	// Ignore comments
	return false
}

func (r *RPC) VisitRPC(rpc *parser.RPC) (next bool) {
	r.RPCName = rpc.RPCName
	r.RequestMessageType = r.types.findType(rpc.RPCRequest.MessageType)
	r.RequestMessageStreaming = rpc.RPCRequest.IsStream
	r.ResponseMessageType = r.types.findType(rpc.RPCResponse.MessageType)
	r.ResponseMessageStreaming = rpc.RPCResponse.IsStream

	// Visit options
	for _, option := range rpc.Options {
		r.VisitOption(option)
	}

	return true
}

func NewRPC(types *TypeDictionary) *RPC {
	r := &RPC{types: types}
	r.Options = NewOptions(r, types)
	return r
}
