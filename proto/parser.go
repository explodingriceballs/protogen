package proto

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"os"
	"path"
)

type ElementType int

const (
	UnknownElementType ElementType = iota
	PackageElementType
	EnumElementType
	MessageElementType
	FieldElementType
	ServiceElementType
	RPCElementType
)

type Element interface {
	Name() string
	GetElementType() ElementType
}

type Parser struct {
	files              []string
	processedFiles     []string
	includeDirectories []string
	ctx                *Context

	messageRegistry map[string]*Message
	enumRegistry    map[string]*Enum
}

func (p *Parser) Parse() error {
	for _, targetFile := range p.files {
		if err := p.parseFile(targetFile, false); err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) parseFile(fileName string, isImported bool) error {
	// Skip over protobuf include files
	// TODO: figure out a better way
	//if strings.HasPrefix(fileName, "google/protobuf") {
	//	return nil
	//}

	// Locate the file
	file, err := p.findFile(fileName)
	if err != nil {
		return err
	}

	// Check if the file was already processed
	for _, processedFile := range p.processedFiles {
		if processedFile == file {
			return nil
		}
	}

	log.Info().Str("file", file).Msg("processing")
	input, err := os.Open(file)
	if err != nil {
		return err
	}

	protoFile, err := protoparser.Parse(input,
		protoparser.WithFilename(fileName),
		protoparser.WithPermissive(true))
	if err != nil {
		return err
	}

	// Version Check
	if protoFile.Syntax == nil {
		return errors.New("unknown or invalid protobuf version")
	}
	if protoFile.Syntax.ProtobufVersion != "proto3" && protoFile.Syntax.ProtobufVersion != "proto2" {
		return errors.New("unknown or invalid protobuf version: " + protoFile.Syntax.ProtobufVersion)
	}

	imports := p.getImports(protoFile)

	//Manage imports
	for _, importedFile := range imports {
		if importedFile[0] == '"' {
			importedFile = importedFile[1 : len(importedFile)-1]
		}
		if err := p.parseFile(importedFile, true); err != nil {
			return err
		}
	}

	packageName := p.getPackage(protoFile)
	p.ctx.AcceptProtoFile(packageName, protoFile)

	// Add the file to processed files
	p.processedFiles = append(p.processedFiles, file)

	return nil
}

func (p *Parser) findFile(file string) (string, error) {
	if _, err := os.Stat(file); err == nil {
		return file, nil
	}
	for _, incDir := range p.includeDirectories {
		foundFile := path.Join(incDir, file)
		if _, err := os.Stat(foundFile); err == nil {
			return foundFile, nil
		}
	}
	return "", errors.New("unable to locate " + file)
}

func (p *Parser) GetContext() *Context {
	return p.ctx
}

func (p *Parser) getImports(file *parser.Proto) []string {
	var imports []string
	for _, body := range file.ProtoBody {
		if importStatement, ok := body.(*parser.Import); ok {
			imports = append(imports, importStatement.Location)
		}
	}
	return imports
}

func (p *Parser) getPackage(file *parser.Proto) string {
	for _, body := range file.ProtoBody {
		if packageStatement, ok := body.(*parser.Package); ok {
			return packageStatement.Name
		}
	}
	return ""
}

func NewParser(files []string, includeDirectors []string) *Parser {
	return &Parser{
		files:              files,
		includeDirectories: includeDirectors,
		ctx:                NewContext(),
	}
}
