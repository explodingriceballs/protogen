package main

import (
	"github.com/explodingriceballs/protogen/js"
	"github.com/explodingriceballs/protogen/proto"
	"github.com/explodingriceballs/protogen/template"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type File struct {
	Name     string
	Contents string
}

type Generator struct {
	sourceDirectory string
	outDirectory    string
	generator       string
	inclDirs        []string
}

func (g *Generator) Run() error {
	files, err := g.ListSourceFiles()
	if err != nil {
		return err
	}

	parser := proto.NewParser(files, g.inclDirs)
	if err := parser.Parse(); err != nil {
		return err
	}

	if err := g.generateFiles(parser.GetContext()); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateFiles(ctx *proto.Context) error {
	// New a new LUA runtime
	filePath := path.Join(g.generator, "index.tsx")
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	runtime := js.NewRuntime(string(fileContents), filePath)
	err = runtime.RegisterModule("protogen/template", &template.Module{TemplateDir: g.generator})
	if err != nil {
		return err
	}

	// Execute once to get hooks
	if err := runtime.Execute(); err != nil {
		return err
	}

	// Get the default exported class
	generatorScript, err := runtime.GetDefaultExport()
	if err != nil {
		return err
	}

	// Create a new class
	if err := generatorScript.Instantiate(); err != nil {
		return err
	}

	// Loop through all parsed packages
	if err := generatorScript.InvokeVoid("Process", ctx); err != nil {
		return err
	}

	var outputFiles []File
	if err := generatorScript.Invoke("GetFiles", &outputFiles); err != nil {
		return err
	}

	for _, outputFile := range outputFiles {
		targetFile := path.Join(g.outDirectory, outputFile.Name)
		targetPath := path.Dir(targetFile)
		if err := os.MkdirAll(targetPath, 0777); err != nil {
			return err
		}
		if err := os.WriteFile(targetFile, []byte(outputFile.Contents), 0777); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) ListSourceFiles() ([]string, error) {
	info, err := os.Stat(g.sourceDirectory)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return []string{g.sourceDirectory}, nil
	}

	var files []string
	err = filepath.Walk(g.sourceDirectory, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".proto" {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (g *Generator) SetInclDirs(dirs []string) {
	g.inclDirs = dirs
}

func CreateGenerator(sourceDirectory string, outDirectory string, generator string) *Generator {
	return &Generator{
		sourceDirectory: sourceDirectory,
		outDirectory:    outDirectory,
		generator:       generator,
	}
}
