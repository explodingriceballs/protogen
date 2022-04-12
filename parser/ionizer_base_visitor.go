// Code generated from Ionizer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Ionizer

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseIonizerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseIonizerVisitor) VisitProto(ctx *ProtoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitSyntax(ctx *SyntaxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitPackageStatement(ctx *PackageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitOptionStatement(ctx *OptionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitOptionName(ctx *OptionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitFieldOptions(ctx *FieldOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitFieldOption(ctx *FieldOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitFieldNumber(ctx *FieldNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitOneof(ctx *OneofContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitOneofField(ctx *OneofFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitMapField(ctx *MapFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitKeyType(ctx *KeyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitReserved(ctx *ReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitExtensionStatement(ctx *ExtensionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitRanges(ctx *RangesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitRange_(ctx *Range_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitReservedFieldNames(ctx *ReservedFieldNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitTopLevelDef(ctx *TopLevelDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumDef(ctx *EnumDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumElement(ctx *EnumElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumField(ctx *EnumFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumValueOptions(ctx *EnumValueOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumValueOption(ctx *EnumValueOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitExtendDef(ctx *ExtendDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitExtendBody(ctx *ExtendBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitExtendElement(ctx *ExtendElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitMessageDef(ctx *MessageDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitMessageBody(ctx *MessageBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitMessageElement(ctx *MessageElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitServiceDef(ctx *ServiceDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitServiceElement(ctx *ServiceElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitRpc(ctx *RpcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitBlockLit(ctx *BlockLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitFullIdent(ctx *FullIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitMessageName(ctx *MessageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitExtendName(ctx *ExtendNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumName(ctx *EnumNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitOneofName(ctx *OneofNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitMapName(ctx *MapNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitServiceName(ctx *ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitRpcName(ctx *RpcNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitMessageType(ctx *MessageTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitEnumType(ctx *EnumTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitIntLit(ctx *IntLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitStrLit(ctx *StrLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitBoolLit(ctx *BoolLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitFloatLit(ctx *FloatLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIonizerVisitor) VisitKeywords(ctx *KeywordsContext) interface{} {
	return v.VisitChildren(ctx)
}
