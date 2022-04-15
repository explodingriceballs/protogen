package main

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/explodingriceballs/protogen/parser"
	"log"
	"os"
	"path"
)

type Parser struct {
	*parser.BaseIonizerListener
	IncludeDirectories []string
	parserState        *ParserState
	file               string
	Imports            []*Import
	ParseErrors        []error
}

func (p *Parser) GetPackage() *Package {
	return p.parserState.pkg
}

func (p *Parser) Parse() error {
	// Locate the file
	file, err := p.findFile()
	if err != nil {
		return err
	}

	input, err := antlr.NewFileStream(file)
	if err != nil {
		return err
	}

	lexer := parser.NewIonizerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	antlrParser := parser.NewIonizerParser(stream)
	antlrParser.RemoveErrorListeners()
	//antlrParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlrParser.BuildParseTrees = true
	syntaxDef := antlrParser.Proto()
	antlr.ParseTreeWalkerDefault.Walk(p, syntaxDef)

	if len(p.ParseErrors) > 0 {
		for _, err := range p.ParseErrors {
			log.Println(err)
		}
	}

	return nil
}

// EnterImportStatement processes import statements
func (p *Parser) EnterImportStatement(ctx *parser.ImportStatementContext) {
	var (
		importFile     = ctx.StrLit().GetText()
		isWeakImport   = false
		isPublicImport = false
	)

	if ctx.WEAK() != nil {
		isWeakImport = true
	}
	if ctx.PUBLIC() != nil {
		isPublicImport = true
	}

	p.Imports = append(p.Imports, NewImport(importFile[1:len(importFile)-1], isWeakImport, isPublicImport))
}

func (p *Parser) EnterPackageStatement(ctx *parser.PackageStatementContext) {
	packageName := ctx.FullIdent().GetText()
	p.parserState.StartPackage(NewPackage(packageName))
}

// EnterServiceDef is called when the parser enters a service definition
func (p *Parser) EnterServiceDef(ctx *parser.ServiceDefContext) {
	// Create a service
	serviceName := ctx.ServiceName().GetText()
	service := NewService(serviceName)

	// Update state
	p.err(p.parserState.StartService(service))
}

// ExitServiceDef finalizes a service definition
func (p *Parser) ExitServiceDef(*parser.ServiceDefContext) {
	// Marshal services
	p.err(p.parserState.FinalizeService())
}

// EnterRpc is called when the parser enters an RPC block
func (p *Parser) EnterRpc(ctx *parser.RpcContext) {
	rpcName := ctx.RpcName().GetText()
	requestMessage := ctx.MessageType(0).GetText()
	responseMessage := ctx.MessageType(1).GetText()

	// Create RPC
	rpc := NewRPC(rpcName, requestMessage, responseMessage)

	streamNodes := ctx.AllSTREAM()
	if len(streamNodes) == 2 {
		rpc.IsRequestStreaming = true
		rpc.IsResponseStreaming = true
	} else if len(streamNodes) == 1 {
		foundReturnNode := false
		// Loop through the children of the RPC
		for _, node := range ctx.GetChildren() {
			// Mark whether we passed the return node
			if node == ctx.RETURNS() {
				foundReturnNode = true
			}

			// If we hit the streaming node, set the streaming flag
			if node == streamNodes[0] {
				if foundReturnNode {
					rpc.IsResponseStreaming = true
				} else {
					rpc.IsRequestStreaming = true
				}
			}
		}
	}

	// Start the RPC block
	p.err(p.parserState.StartRPC(rpc))
}

func (p *Parser) ExitRpc(*parser.RpcContext) {
	p.err(p.parserState.FinalizeRPC())
}

// EnterMessageDef is called when the parser enters a message definition
func (p *Parser) EnterMessageDef(ctx *parser.MessageDefContext) {
	messageName := ctx.MessageName().GetText()
	message := NewMessage(messageName)

	// Update state
	p.err(p.parserState.StartMessage(message))
}

// ExitMessageDef is called when the parser leaves a message definition
func (p *Parser) ExitMessageDef(*parser.MessageDefContext) {
	// Marshal message
	p.err(p.parserState.FinalizeMessage())
}

// EnterExtendDef is called when production extendDef is entered.
func (p *Parser) EnterExtendDef(ctx *parser.ExtendDefContext) {
	messageName := ctx.ExtendName().GetText()
	message := NewMessage(messageName)

	// Update state
	p.err(p.parserState.StartExtension(message))
}

// ExitExtendDef is called when production extendDef is exited.
func (p *Parser) ExitExtendDef(_ *parser.ExtendDefContext) {
	p.err(p.parserState.FinalizeExtension())
}

// EnterField is called when a field is entered.
func (p *Parser) EnterField(ctx *parser.FieldContext) {
	fieldName := ctx.FieldName().GetText()

	// Create a new field
	field := NewField(fieldName, ctx.Type_().GetText())
	if ctx.FieldNumber() != nil {
		field.FieldNumber = ctx.FieldNumber().GetText()
	}

	if ctx.REPEATED() != nil {
		field.IsRepeated = true
	}

	// Start the field
	p.err(p.parserState.StartField(field))
}

// ExitField is called when a field is exited.
func (p *Parser) ExitField(*parser.FieldContext) {
	// Marshal message
	p.err(p.parserState.FinalizeField())
}

// EnterOptionStatement is called when the parser enters an option statement
func (p *Parser) EnterOptionStatement(ctx *parser.OptionStatementContext) {
	key := ctx.OptionName().GetText()
	value := ctx.Constant().GetText()

	option := NewOption(key, value[1:len(value)-1])
	p.err(p.parserState.AddOption(option))
}

// EnterFieldOption is called when an option on a field is entered (not required to exit)
func (p *Parser) EnterFieldOption(ctx *parser.FieldOptionContext) {
	key := ctx.OptionName().GetText()
	value := ctx.Constant().GetText()

	if key[0] == '"' || key[0] == '(' {
		key = key[1 : len(key)-1]
	}
	if value[0] == '"' {
		value = value[1 : len(value)-1]
	}

	option := NewOption(key, value)
	p.err(p.parserState.AddOption(option))
}

// EnterEnumDef is called when an enumDef is entered.
func (p *Parser) EnterEnumDef(ctx *parser.EnumDefContext) {
	enumName := ctx.EnumName().GetText()
	p.err(p.parserState.StartEnum(NewEnum(enumName)))
}

// ExitEnumDef is called when an enumDef is exited.
func (p *Parser) ExitEnumDef(_ *parser.EnumDefContext) {
	p.err(p.parserState.FinalizeEnum())
}

// EnterEnumField is called when an enumField is entered.
func (p *Parser) EnterEnumField(ctx *parser.EnumFieldContext) {
	identifier := ctx.Ident().GetText()
	value := ctx.IntLit().GetText()

	enumValue := NewEnumValue(identifier, value)
	p.err(p.parserState.StartEnumValue(enumValue))
}

// ExitEnumField is called when an enumField is exited.
func (p *Parser) ExitEnumField(_ *parser.EnumFieldContext) {
	p.err(p.parserState.FinalizeEnumValue())
}

func (p *Parser) err(err error) {
	if err != nil {
		p.ParseErrors = append(p.ParseErrors, err)
	}
}

func (p *Parser) findFile() (string, error) {
	if _, err := os.Stat(p.file); err == nil {
		return p.file, nil
	}
	for _, incDir := range p.IncludeDirectories {
		file := path.Join(incDir, p.file)
		if _, err := os.Stat(file); err == nil {
			return file, nil
		}
	}
	return "", errors.New("unable to locate " + p.file)
}

func NewParser(file string) *Parser {
	return &Parser{
		parserState: &ParserState{},
		file:        file,
	}
}
