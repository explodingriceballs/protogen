// Code generated from /Users/bneuenfeldt/Development/source/protogen/Ionizer.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Ionizer
import "github.com/antlr/antlr4/runtime/Go/antlr"

// IonizerListener is a complete listener for a parse tree produced by IonizerParser.
type IonizerListener interface {
	antlr.ParseTreeListener

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterPackageStatement is called when entering the packageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterOptionStatement is called when entering the optionStatement production.
	EnterOptionStatement(c *OptionStatementContext)

	// EnterOptionName is called when entering the optionName production.
	EnterOptionName(c *OptionNameContext)

	// EnterOptionBody is called when entering the optionBody production.
	EnterOptionBody(c *OptionBodyContext)

	// EnterOptionElement is called when entering the optionElement production.
	EnterOptionElement(c *OptionElementContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldOptions is called when entering the fieldOptions production.
	EnterFieldOptions(c *FieldOptionsContext)

	// EnterFieldOption is called when entering the fieldOption production.
	EnterFieldOption(c *FieldOptionContext)

	// EnterFieldNumber is called when entering the fieldNumber production.
	EnterFieldNumber(c *FieldNumberContext)

	// EnterOneof is called when entering the oneof production.
	EnterOneof(c *OneofContext)

	// EnterOneofField is called when entering the oneofField production.
	EnterOneofField(c *OneofFieldContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterKeyType is called when entering the keyType production.
	EnterKeyType(c *KeyTypeContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterReserved is called when entering the reserved production.
	EnterReserved(c *ReservedContext)

	// EnterExtensionStatement is called when entering the extensionStatement production.
	EnterExtensionStatement(c *ExtensionStatementContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

	// EnterRange_ is called when entering the range_ production.
	EnterRange_(c *Range_Context)

	// EnterReservedFieldNames is called when entering the reservedFieldNames production.
	EnterReservedFieldNames(c *ReservedFieldNamesContext)

	// EnterTopLevelDef is called when entering the topLevelDef production.
	EnterTopLevelDef(c *TopLevelDefContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumElement is called when entering the enumElement production.
	EnterEnumElement(c *EnumElementContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterEnumValueOptions is called when entering the enumValueOptions production.
	EnterEnumValueOptions(c *EnumValueOptionsContext)

	// EnterEnumValueOption is called when entering the enumValueOption production.
	EnterEnumValueOption(c *EnumValueOptionContext)

	// EnterExtendDef is called when entering the extendDef production.
	EnterExtendDef(c *ExtendDefContext)

	// EnterExtendBody is called when entering the extendBody production.
	EnterExtendBody(c *ExtendBodyContext)

	// EnterExtendElement is called when entering the extendElement production.
	EnterExtendElement(c *ExtendElementContext)

	// EnterMessageDef is called when entering the messageDef production.
	EnterMessageDef(c *MessageDefContext)

	// EnterMessageBody is called when entering the messageBody production.
	EnterMessageBody(c *MessageBodyContext)

	// EnterMessageElement is called when entering the messageElement production.
	EnterMessageElement(c *MessageElementContext)

	// EnterServiceDef is called when entering the serviceDef production.
	EnterServiceDef(c *ServiceDefContext)

	// EnterServiceElement is called when entering the serviceElement production.
	EnterServiceElement(c *ServiceElementContext)

	// EnterRpc is called when entering the rpc production.
	EnterRpc(c *RpcContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterBlockLit is called when entering the blockLit production.
	EnterBlockLit(c *BlockLitContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterFullIdent is called when entering the fullIdent production.
	EnterFullIdent(c *FullIdentContext)

	// EnterMessageName is called when entering the messageName production.
	EnterMessageName(c *MessageNameContext)

	// EnterExtendName is called when entering the extendName production.
	EnterExtendName(c *ExtendNameContext)

	// EnterEnumName is called when entering the enumName production.
	EnterEnumName(c *EnumNameContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterOneofName is called when entering the oneofName production.
	EnterOneofName(c *OneofNameContext)

	// EnterMapName is called when entering the mapName production.
	EnterMapName(c *MapNameContext)

	// EnterServiceName is called when entering the serviceName production.
	EnterServiceName(c *ServiceNameContext)

	// EnterRpcName is called when entering the rpcName production.
	EnterRpcName(c *RpcNameContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterEnumType is called when entering the enumType production.
	EnterEnumType(c *EnumTypeContext)

	// EnterIntLit is called when entering the intLit production.
	EnterIntLit(c *IntLitContext)

	// EnterStrLit is called when entering the strLit production.
	EnterStrLit(c *StrLitContext)

	// EnterBoolLit is called when entering the boolLit production.
	EnterBoolLit(c *BoolLitContext)

	// EnterFloatLit is called when entering the floatLit production.
	EnterFloatLit(c *FloatLitContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitPackageStatement is called when exiting the packageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitOptionStatement is called when exiting the optionStatement production.
	ExitOptionStatement(c *OptionStatementContext)

	// ExitOptionName is called when exiting the optionName production.
	ExitOptionName(c *OptionNameContext)

	// ExitOptionBody is called when exiting the optionBody production.
	ExitOptionBody(c *OptionBodyContext)

	// ExitOptionElement is called when exiting the optionElement production.
	ExitOptionElement(c *OptionElementContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldOptions is called when exiting the fieldOptions production.
	ExitFieldOptions(c *FieldOptionsContext)

	// ExitFieldOption is called when exiting the fieldOption production.
	ExitFieldOption(c *FieldOptionContext)

	// ExitFieldNumber is called when exiting the fieldNumber production.
	ExitFieldNumber(c *FieldNumberContext)

	// ExitOneof is called when exiting the oneof production.
	ExitOneof(c *OneofContext)

	// ExitOneofField is called when exiting the oneofField production.
	ExitOneofField(c *OneofFieldContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitKeyType is called when exiting the keyType production.
	ExitKeyType(c *KeyTypeContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitReserved is called when exiting the reserved production.
	ExitReserved(c *ReservedContext)

	// ExitExtensionStatement is called when exiting the extensionStatement production.
	ExitExtensionStatement(c *ExtensionStatementContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)

	// ExitRange_ is called when exiting the range_ production.
	ExitRange_(c *Range_Context)

	// ExitReservedFieldNames is called when exiting the reservedFieldNames production.
	ExitReservedFieldNames(c *ReservedFieldNamesContext)

	// ExitTopLevelDef is called when exiting the topLevelDef production.
	ExitTopLevelDef(c *TopLevelDefContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumElement is called when exiting the enumElement production.
	ExitEnumElement(c *EnumElementContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitEnumValueOptions is called when exiting the enumValueOptions production.
	ExitEnumValueOptions(c *EnumValueOptionsContext)

	// ExitEnumValueOption is called when exiting the enumValueOption production.
	ExitEnumValueOption(c *EnumValueOptionContext)

	// ExitExtendDef is called when exiting the extendDef production.
	ExitExtendDef(c *ExtendDefContext)

	// ExitExtendBody is called when exiting the extendBody production.
	ExitExtendBody(c *ExtendBodyContext)

	// ExitExtendElement is called when exiting the extendElement production.
	ExitExtendElement(c *ExtendElementContext)

	// ExitMessageDef is called when exiting the messageDef production.
	ExitMessageDef(c *MessageDefContext)

	// ExitMessageBody is called when exiting the messageBody production.
	ExitMessageBody(c *MessageBodyContext)

	// ExitMessageElement is called when exiting the messageElement production.
	ExitMessageElement(c *MessageElementContext)

	// ExitServiceDef is called when exiting the serviceDef production.
	ExitServiceDef(c *ServiceDefContext)

	// ExitServiceElement is called when exiting the serviceElement production.
	ExitServiceElement(c *ServiceElementContext)

	// ExitRpc is called when exiting the rpc production.
	ExitRpc(c *RpcContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitBlockLit is called when exiting the blockLit production.
	ExitBlockLit(c *BlockLitContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitFullIdent is called when exiting the fullIdent production.
	ExitFullIdent(c *FullIdentContext)

	// ExitMessageName is called when exiting the messageName production.
	ExitMessageName(c *MessageNameContext)

	// ExitExtendName is called when exiting the extendName production.
	ExitExtendName(c *ExtendNameContext)

	// ExitEnumName is called when exiting the enumName production.
	ExitEnumName(c *EnumNameContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitOneofName is called when exiting the oneofName production.
	ExitOneofName(c *OneofNameContext)

	// ExitMapName is called when exiting the mapName production.
	ExitMapName(c *MapNameContext)

	// ExitServiceName is called when exiting the serviceName production.
	ExitServiceName(c *ServiceNameContext)

	// ExitRpcName is called when exiting the rpcName production.
	ExitRpcName(c *RpcNameContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitEnumType is called when exiting the enumType production.
	ExitEnumType(c *EnumTypeContext)

	// ExitIntLit is called when exiting the intLit production.
	ExitIntLit(c *IntLitContext)

	// ExitStrLit is called when exiting the strLit production.
	ExitStrLit(c *StrLitContext)

	// ExitBoolLit is called when exiting the boolLit production.
	ExitBoolLit(c *BoolLitContext)

	// ExitFloatLit is called when exiting the floatLit production.
	ExitFloatLit(c *FloatLitContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)
}
