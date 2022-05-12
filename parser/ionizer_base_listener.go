// Code generated from /Users/bneuenfeldt/Development/source/protogen/Ionizer.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Ionizer
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIonizerListener is a complete listener for a parse tree produced by IonizerParser.
type BaseIonizerListener struct{}

var _ IonizerListener = &BaseIonizerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIonizerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIonizerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIonizerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIonizerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *BaseIonizerListener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BaseIonizerListener) ExitProto(ctx *ProtoContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BaseIonizerListener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseIonizerListener) ExitSyntax(ctx *SyntaxContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseIonizerListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseIonizerListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BaseIonizerListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BaseIonizerListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterOptionStatement is called when production optionStatement is entered.
func (s *BaseIonizerListener) EnterOptionStatement(ctx *OptionStatementContext) {}

// ExitOptionStatement is called when production optionStatement is exited.
func (s *BaseIonizerListener) ExitOptionStatement(ctx *OptionStatementContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BaseIonizerListener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BaseIonizerListener) ExitOptionName(ctx *OptionNameContext) {}

// EnterOptionBody is called when production optionBody is entered.
func (s *BaseIonizerListener) EnterOptionBody(ctx *OptionBodyContext) {}

// ExitOptionBody is called when production optionBody is exited.
func (s *BaseIonizerListener) ExitOptionBody(ctx *OptionBodyContext) {}

// EnterOptionElement is called when production optionElement is entered.
func (s *BaseIonizerListener) EnterOptionElement(ctx *OptionElementContext) {}

// ExitOptionElement is called when production optionElement is exited.
func (s *BaseIonizerListener) ExitOptionElement(ctx *OptionElementContext) {}

// EnterField is called when production field is entered.
func (s *BaseIonizerListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseIonizerListener) ExitField(ctx *FieldContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseIonizerListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseIonizerListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterFieldOption is called when production fieldOption is entered.
func (s *BaseIonizerListener) EnterFieldOption(ctx *FieldOptionContext) {}

// ExitFieldOption is called when production fieldOption is exited.
func (s *BaseIonizerListener) ExitFieldOption(ctx *FieldOptionContext) {}

// EnterFieldNumber is called when production fieldNumber is entered.
func (s *BaseIonizerListener) EnterFieldNumber(ctx *FieldNumberContext) {}

// ExitFieldNumber is called when production fieldNumber is exited.
func (s *BaseIonizerListener) ExitFieldNumber(ctx *FieldNumberContext) {}

// EnterOneof is called when production oneof is entered.
func (s *BaseIonizerListener) EnterOneof(ctx *OneofContext) {}

// ExitOneof is called when production oneof is exited.
func (s *BaseIonizerListener) ExitOneof(ctx *OneofContext) {}

// EnterOneofField is called when production oneofField is entered.
func (s *BaseIonizerListener) EnterOneofField(ctx *OneofFieldContext) {}

// ExitOneofField is called when production oneofField is exited.
func (s *BaseIonizerListener) ExitOneofField(ctx *OneofFieldContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BaseIonizerListener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BaseIonizerListener) ExitMapField(ctx *MapFieldContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BaseIonizerListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BaseIonizerListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseIonizerListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseIonizerListener) ExitType_(ctx *Type_Context) {}

// EnterReserved is called when production reserved is entered.
func (s *BaseIonizerListener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BaseIonizerListener) ExitReserved(ctx *ReservedContext) {}

// EnterExtensionStatement is called when production extensionStatement is entered.
func (s *BaseIonizerListener) EnterExtensionStatement(ctx *ExtensionStatementContext) {}

// ExitExtensionStatement is called when production extensionStatement is exited.
func (s *BaseIonizerListener) ExitExtensionStatement(ctx *ExtensionStatementContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BaseIonizerListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BaseIonizerListener) ExitRanges(ctx *RangesContext) {}

// EnterRange_ is called when production range_ is entered.
func (s *BaseIonizerListener) EnterRange_(ctx *Range_Context) {}

// ExitRange_ is called when production range_ is exited.
func (s *BaseIonizerListener) ExitRange_(ctx *Range_Context) {}

// EnterReservedFieldNames is called when production reservedFieldNames is entered.
func (s *BaseIonizerListener) EnterReservedFieldNames(ctx *ReservedFieldNamesContext) {}

// ExitReservedFieldNames is called when production reservedFieldNames is exited.
func (s *BaseIonizerListener) ExitReservedFieldNames(ctx *ReservedFieldNamesContext) {}

// EnterTopLevelDef is called when production topLevelDef is entered.
func (s *BaseIonizerListener) EnterTopLevelDef(ctx *TopLevelDefContext) {}

// ExitTopLevelDef is called when production topLevelDef is exited.
func (s *BaseIonizerListener) ExitTopLevelDef(ctx *TopLevelDefContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseIonizerListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseIonizerListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseIonizerListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseIonizerListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumElement is called when production enumElement is entered.
func (s *BaseIonizerListener) EnterEnumElement(ctx *EnumElementContext) {}

// ExitEnumElement is called when production enumElement is exited.
func (s *BaseIonizerListener) ExitEnumElement(ctx *EnumElementContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseIonizerListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseIonizerListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterEnumValueOptions is called when production enumValueOptions is entered.
func (s *BaseIonizerListener) EnterEnumValueOptions(ctx *EnumValueOptionsContext) {}

// ExitEnumValueOptions is called when production enumValueOptions is exited.
func (s *BaseIonizerListener) ExitEnumValueOptions(ctx *EnumValueOptionsContext) {}

// EnterEnumValueOption is called when production enumValueOption is entered.
func (s *BaseIonizerListener) EnterEnumValueOption(ctx *EnumValueOptionContext) {}

// ExitEnumValueOption is called when production enumValueOption is exited.
func (s *BaseIonizerListener) ExitEnumValueOption(ctx *EnumValueOptionContext) {}

// EnterExtendDef is called when production extendDef is entered.
func (s *BaseIonizerListener) EnterExtendDef(ctx *ExtendDefContext) {}

// ExitExtendDef is called when production extendDef is exited.
func (s *BaseIonizerListener) ExitExtendDef(ctx *ExtendDefContext) {}

// EnterExtendBody is called when production extendBody is entered.
func (s *BaseIonizerListener) EnterExtendBody(ctx *ExtendBodyContext) {}

// ExitExtendBody is called when production extendBody is exited.
func (s *BaseIonizerListener) ExitExtendBody(ctx *ExtendBodyContext) {}

// EnterExtendElement is called when production extendElement is entered.
func (s *BaseIonizerListener) EnterExtendElement(ctx *ExtendElementContext) {}

// ExitExtendElement is called when production extendElement is exited.
func (s *BaseIonizerListener) ExitExtendElement(ctx *ExtendElementContext) {}

// EnterMessageDef is called when production messageDef is entered.
func (s *BaseIonizerListener) EnterMessageDef(ctx *MessageDefContext) {}

// ExitMessageDef is called when production messageDef is exited.
func (s *BaseIonizerListener) ExitMessageDef(ctx *MessageDefContext) {}

// EnterMessageBody is called when production messageBody is entered.
func (s *BaseIonizerListener) EnterMessageBody(ctx *MessageBodyContext) {}

// ExitMessageBody is called when production messageBody is exited.
func (s *BaseIonizerListener) ExitMessageBody(ctx *MessageBodyContext) {}

// EnterMessageElement is called when production messageElement is entered.
func (s *BaseIonizerListener) EnterMessageElement(ctx *MessageElementContext) {}

// ExitMessageElement is called when production messageElement is exited.
func (s *BaseIonizerListener) ExitMessageElement(ctx *MessageElementContext) {}

// EnterServiceDef is called when production serviceDef is entered.
func (s *BaseIonizerListener) EnterServiceDef(ctx *ServiceDefContext) {}

// ExitServiceDef is called when production serviceDef is exited.
func (s *BaseIonizerListener) ExitServiceDef(ctx *ServiceDefContext) {}

// EnterServiceElement is called when production serviceElement is entered.
func (s *BaseIonizerListener) EnterServiceElement(ctx *ServiceElementContext) {}

// ExitServiceElement is called when production serviceElement is exited.
func (s *BaseIonizerListener) ExitServiceElement(ctx *ServiceElementContext) {}

// EnterRpc is called when production rpc is entered.
func (s *BaseIonizerListener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *BaseIonizerListener) ExitRpc(ctx *RpcContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseIonizerListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseIonizerListener) ExitConstant(ctx *ConstantContext) {}

// EnterBlockLit is called when production blockLit is entered.
func (s *BaseIonizerListener) EnterBlockLit(ctx *BlockLitContext) {}

// ExitBlockLit is called when production blockLit is exited.
func (s *BaseIonizerListener) ExitBlockLit(ctx *BlockLitContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseIonizerListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseIonizerListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterIdent is called when production ident is entered.
func (s *BaseIonizerListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseIonizerListener) ExitIdent(ctx *IdentContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BaseIonizerListener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BaseIonizerListener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterMessageName is called when production messageName is entered.
func (s *BaseIonizerListener) EnterMessageName(ctx *MessageNameContext) {}

// ExitMessageName is called when production messageName is exited.
func (s *BaseIonizerListener) ExitMessageName(ctx *MessageNameContext) {}

// EnterExtendName is called when production extendName is entered.
func (s *BaseIonizerListener) EnterExtendName(ctx *ExtendNameContext) {}

// ExitExtendName is called when production extendName is exited.
func (s *BaseIonizerListener) ExitExtendName(ctx *ExtendNameContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BaseIonizerListener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BaseIonizerListener) ExitEnumName(ctx *EnumNameContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseIonizerListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseIonizerListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterOneofName is called when production oneofName is entered.
func (s *BaseIonizerListener) EnterOneofName(ctx *OneofNameContext) {}

// ExitOneofName is called when production oneofName is exited.
func (s *BaseIonizerListener) ExitOneofName(ctx *OneofNameContext) {}

// EnterMapName is called when production mapName is entered.
func (s *BaseIonizerListener) EnterMapName(ctx *MapNameContext) {}

// ExitMapName is called when production mapName is exited.
func (s *BaseIonizerListener) ExitMapName(ctx *MapNameContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseIonizerListener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseIonizerListener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterRpcName is called when production rpcName is entered.
func (s *BaseIonizerListener) EnterRpcName(ctx *RpcNameContext) {}

// ExitRpcName is called when production rpcName is exited.
func (s *BaseIonizerListener) ExitRpcName(ctx *RpcNameContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BaseIonizerListener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BaseIonizerListener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BaseIonizerListener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BaseIonizerListener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterIntLit is called when production intLit is entered.
func (s *BaseIonizerListener) EnterIntLit(ctx *IntLitContext) {}

// ExitIntLit is called when production intLit is exited.
func (s *BaseIonizerListener) ExitIntLit(ctx *IntLitContext) {}

// EnterStrLit is called when production strLit is entered.
func (s *BaseIonizerListener) EnterStrLit(ctx *StrLitContext) {}

// ExitStrLit is called when production strLit is exited.
func (s *BaseIonizerListener) ExitStrLit(ctx *StrLitContext) {}

// EnterBoolLit is called when production boolLit is entered.
func (s *BaseIonizerListener) EnterBoolLit(ctx *BoolLitContext) {}

// ExitBoolLit is called when production boolLit is exited.
func (s *BaseIonizerListener) ExitBoolLit(ctx *BoolLitContext) {}

// EnterFloatLit is called when production floatLit is entered.
func (s *BaseIonizerListener) EnterFloatLit(ctx *FloatLitContext) {}

// ExitFloatLit is called when production floatLit is exited.
func (s *BaseIonizerListener) ExitFloatLit(ctx *FloatLitContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseIonizerListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseIonizerListener) ExitKeywords(ctx *KeywordsContext) {}
