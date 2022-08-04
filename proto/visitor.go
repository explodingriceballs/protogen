package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type noopVisitor struct {
}

func (n noopVisitor) VisitComment(comment *parser.Comment) {}

func (n noopVisitor) VisitEmptyStatement(statement *parser.EmptyStatement) (next bool) {
	return false
}

func (n noopVisitor) VisitEnum(enum *parser.Enum) (next bool) {
	return false
}

func (n noopVisitor) VisitEnumField(field *parser.EnumField) (next bool) {
	return false
}

func (n noopVisitor) VisitExtend(extend *parser.Extend) (next bool) {
	return false
}

func (n noopVisitor) VisitExtensions(extensions *parser.Extensions) (next bool) {
	return false
}

func (n noopVisitor) VisitField(field *parser.Field) (next bool) {
	return false
}

func (n noopVisitor) VisitGroupField(field *parser.GroupField) (next bool) {
	return false
}

func (n noopVisitor) VisitImport(i *parser.Import) (next bool) {
	return false
}

func (n noopVisitor) VisitMapField(field *parser.MapField) (next bool) {
	return false
}

func (n noopVisitor) VisitMessage(message *parser.Message) (next bool) {
	return false
}

func (n noopVisitor) VisitOneof(oneof *parser.Oneof) (next bool) {
	return false
}

func (n noopVisitor) VisitOneofField(field *parser.OneofField) (next bool) {
	return false
}

func (n noopVisitor) VisitOption(option *parser.Option) (next bool) {
	return false
}

func (n noopVisitor) VisitPackage(p *parser.Package) (next bool) {
	return false
}

func (n noopVisitor) VisitReserved(reserved *parser.Reserved) (next bool) {
	return false
}

func (n noopVisitor) VisitRPC(rpc *parser.RPC) (next bool) {
	return false
}

func (n noopVisitor) VisitService(service *parser.Service) (next bool) {
	return false
}

func (n noopVisitor) VisitSyntax(syntax *parser.Syntax) (next bool) {
	return false
}
