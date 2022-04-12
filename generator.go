package main

import (
	"github.com/explodingriceballs/protogen/vm"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type Generator struct {
	sourceDirectory string
	outDirectory    string
	languages       string
	packages        []*Package
}

func (g *Generator) Run() error {
	files, err := g.ListSourceFiles()
	if err != nil {
		return err
	}

	// Loop through all imported files
	for _, file := range files {
		if err := g.processFile(file); err != nil {
			return err
		}
	}

	if err := g.generateFiles(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateFiles() error {
	// Create a new LUA runtime
	tmplDir := g.findTemplateDirectory(g.languages)
	filePath := path.Join(tmplDir, "index.tsx")
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

	engine := NewTemplateEngine(tmplDir)
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

	//packagePath := pkg.Name
	//
	////Create the package output directory
	//if filepath.Base(g.outDirectory) == pkg.Name {
	//	packagePath = g.outDirectory
	//} else {
	//	packagePath = path.Join(g.outDirectory, pkg.Name)
	//}

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
	// Create a parser
	parser := NewParser(file)
	parser.IncludeDirectories = viper.GetStringSlice("includeDirectories")

	// Parse protobuf file
	if err := parser.Parse(); err != nil {
		return err
	}

	// Loop over imports
	for _, importedFile := range parser.Imports {
		if err := g.processFile(importedFile.File); err != nil {
			return err
		}
	}

	// Add the package
	if parser.GetPackage() != nil {
		g.packages = append(g.packages, parser.GetPackage())
	}
	return nil
}

func (g *Generator) findTemplateDirectory(language string) string {
	for _, templateDir := range viper.GetStringSlice("templateDirectories") {
		if _, err := os.Stat(path.Join(templateDir, language, "index.tsx")); err == nil {
			return path.Join(templateDir, language)
		}
	}
	return ""
}

func CreateGenerator(sourceDirectory string, outDirectory string, languages string) *Generator {
	return &Generator{
		sourceDirectory: sourceDirectory,
		outDirectory:    outDirectory,
		languages:       languages,
		packages:        []*Package{},
	}
}
