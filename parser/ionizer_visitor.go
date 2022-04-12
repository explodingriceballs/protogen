// Code generated from Ionizer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Ionizer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by IonizerParser.
type IonizerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by IonizerParser#proto.
	VisitProto(ctx *ProtoContext) interface{}

	// Visit a parse tree produced by IonizerParser#syntax.
	VisitSyntax(ctx *SyntaxContext) interface{}

	// Visit a parse tree produced by IonizerParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by IonizerParser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) interface{}

	// Visit a parse tree produced by IonizerParser#optionStatement.
	VisitOptionStatement(ctx *OptionStatementContext) interface{}

	// Visit a parse tree produced by IonizerParser#optionName.
	VisitOptionName(ctx *OptionNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by IonizerParser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by IonizerParser#fieldOption.
	VisitFieldOption(ctx *FieldOptionContext) interface{}

	// Visit a parse tree produced by IonizerParser#fieldNumber.
	VisitFieldNumber(ctx *FieldNumberContext) interface{}

	// Visit a parse tree produced by IonizerParser#oneof.
	VisitOneof(ctx *OneofContext) interface{}

	// Visit a parse tree produced by IonizerParser#oneofField.
	VisitOneofField(ctx *OneofFieldContext) interface{}

	// Visit a parse tree produced by IonizerParser#mapField.
	VisitMapField(ctx *MapFieldContext) interface{}

	// Visit a parse tree produced by IonizerParser#keyType.
	VisitKeyType(ctx *KeyTypeContext) interface{}

	// Visit a parse tree produced by IonizerParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by IonizerParser#reserved.
	VisitReserved(ctx *ReservedContext) interface{}

	// Visit a parse tree produced by IonizerParser#extensionStatement.
	VisitExtensionStatement(ctx *ExtensionStatementContext) interface{}

	// Visit a parse tree produced by IonizerParser#ranges.
	VisitRanges(ctx *RangesContext) interface{}

	// Visit a parse tree produced by IonizerParser#range_.
	VisitRange_(ctx *Range_Context) interface{}

	// Visit a parse tree produced by IonizerParser#reservedFieldNames.
	VisitReservedFieldNames(ctx *ReservedFieldNamesContext) interface{}

	// Visit a parse tree produced by IonizerParser#topLevelDef.
	VisitTopLevelDef(ctx *TopLevelDefContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumDef.
	VisitEnumDef(ctx *EnumDefContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumElement.
	VisitEnumElement(ctx *EnumElementContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumField.
	VisitEnumField(ctx *EnumFieldContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumValueOptions.
	VisitEnumValueOptions(ctx *EnumValueOptionsContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumValueOption.
	VisitEnumValueOption(ctx *EnumValueOptionContext) interface{}

	// Visit a parse tree produced by IonizerParser#extendDef.
	VisitExtendDef(ctx *ExtendDefContext) interface{}

	// Visit a parse tree produced by IonizerParser#extendBody.
	VisitExtendBody(ctx *ExtendBodyContext) interface{}

	// Visit a parse tree produced by IonizerParser#extendElement.
	VisitExtendElement(ctx *ExtendElementContext) interface{}

	// Visit a parse tree produced by IonizerParser#messageDef.
	VisitMessageDef(ctx *MessageDefContext) interface{}

	// Visit a parse tree produced by IonizerParser#messageBody.
	VisitMessageBody(ctx *MessageBodyContext) interface{}

	// Visit a parse tree produced by IonizerParser#messageElement.
	VisitMessageElement(ctx *MessageElementContext) interface{}

	// Visit a parse tree produced by IonizerParser#serviceDef.
	VisitServiceDef(ctx *ServiceDefContext) interface{}

	// Visit a parse tree produced by IonizerParser#serviceElement.
	VisitServiceElement(ctx *ServiceElementContext) interface{}

	// Visit a parse tree produced by IonizerParser#rpc.
	VisitRpc(ctx *RpcContext) interface{}

	// Visit a parse tree produced by IonizerParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by IonizerParser#blockLit.
	VisitBlockLit(ctx *BlockLitContext) interface{}

	// Visit a parse tree produced by IonizerParser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by IonizerParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by IonizerParser#fullIdent.
	VisitFullIdent(ctx *FullIdentContext) interface{}

	// Visit a parse tree produced by IonizerParser#messageName.
	VisitMessageName(ctx *MessageNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#extendName.
	VisitExtendName(ctx *ExtendNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumName.
	VisitEnumName(ctx *EnumNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#oneofName.
	VisitOneofName(ctx *OneofNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#mapName.
	VisitMapName(ctx *MapNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#rpcName.
	VisitRpcName(ctx *RpcNameContext) interface{}

	// Visit a parse tree produced by IonizerParser#messageType.
	VisitMessageType(ctx *MessageTypeContext) interface{}

	// Visit a parse tree produced by IonizerParser#enumType.
	VisitEnumType(ctx *EnumTypeContext) interface{}

	// Visit a parse tree produced by IonizerParser#intLit.
	VisitIntLit(ctx *IntLitContext) interface{}

	// Visit a parse tree produced by IonizerParser#strLit.
	VisitStrLit(ctx *StrLitContext) interface{}

	// Visit a parse tree produced by IonizerParser#boolLit.
	VisitBoolLit(ctx *BoolLitContext) interface{}

	// Visit a parse tree produced by IonizerParser#floatLit.
	VisitFloatLit(ctx *FloatLitContext) interface{}

	// Visit a parse tree produced by IonizerParser#keywords.
	VisitKeywords(ctx *KeywordsContext) interface{}
}
