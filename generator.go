package main

import (
	"github.com/explodingriceballs/protogen/proto"
	"github.com/explodingriceballs/protogen/vm"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type Generator struct {
	sourceDirectory string
	outDirectory    string
	generator       string
	packages        []*proto.Package
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

	// Loop through all imported files
	//for _, file := range files {
	//	if err := g.processFile(file); err != nil {
	//		return err
	//	}
	//}

	if err := g.generateFiles(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateFiles() error {
	// Create a new LUA runtime
	filePath := path.Join(g.generator, "index.tsx")
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	runtime := vm.NewRuntime(string(fileContents), filePath)

	// Execute once to get hooks
	if err := runtime.Execute(); err != nil {
		return err
	}

	// Get the default exported class
	generatorScript, err := runtime.GetDefaultExport()
	if err != nil {
		return err
	}

	engine := NewTemplateEngine(g.generator, generatorScript)
	engineObj, err := engine.Create(runtime)
	if err != nil {
		return err
	}

	// Create a new class
	if err := generatorScript.Instantiate(engineObj); err != nil {
		return err
	}

	// Loop through all parsed packages
	for _, pkg := range g.packages {
		if err := generatorScript.InvokeVoid("processPackage", pkg); err != nil {
			return err
		}
	}

	var outputFiles []File
	if err := generatorScript.Invoke("getFiles", &outputFiles); err != nil {
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

func (g *Generator) processFile(file string) error {
	// Create a proto
	//proto := NewParser(file)
	//proto.IncludeDirectories = g.inclDirs
	//
	//// Parse protobuf file
	//if err := proto.Parse(); err != nil {
	//	return err
	//}
	//
	//// Loop over imports
	//for _, importedFile := range proto.Imports {
	//	if err := g.processFile(importedFile.File); err != nil {
	//		return err
	//	}
	//}

	// Add the package
	//if proto.GetPackage() != nil {
	//	g.packages = append(g.packages, proto.GetPackage())
	//}
	return nil
}

func (g *Generator) SetInclDirs(dirs []string) {
	g.inclDirs = dirs
}

func CreateGenerator(sourceDirectory string, outDirectory string, generator string) *Generator {
	return &Generator{
		sourceDirectory: sourceDirectory,
		outDirectory:    outDirectory,
		generator:       generator,
		packages:        []*proto.Package{},
	}
}
