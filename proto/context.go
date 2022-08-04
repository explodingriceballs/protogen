package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

// Context captures the global parsing state of the files
type Context struct {
	packages       []*Package
	defaultPackage *Package
	typeDictionary *TypeDictionary
}

func (c *Context) AcceptProtoFile(pkgName string, proto *parser.Proto) {
	// Check if it should go into the default package
	if pkgName == "" {
		c.typeDictionary.startScope("")
		proto.Accept(c.defaultPackage)
		c.typeDictionary.endScope("")
		return
	}

	// Loop thru existing packages
	for _, pkg := range c.packages {
		if pkg.Name == pkgName {
			c.typeDictionary.startScope(pkg.Name)
			proto.Accept(pkg)
			c.typeDictionary.endScope(pkg.Name)
			return
		}
	}

	// Create a new package
	newPkg := NewPackage(pkgName, c.typeDictionary)
	c.typeDictionary.startScope(pkgName)
	proto.Accept(newPkg)
	c.typeDictionary.endScope(pkgName)
	c.packages = append(c.packages, newPkg)
}

func (c *Context) Packages() []*Package {
	return c.packages
}

func (c *Context) GetPackage(name string) *Package {
	for _, pkg := range c.packages {
		if pkg.Name == name {
			return pkg
		}
	}
	return nil
}

func (c *Context) GlobalScope() *Package {
	return c.defaultPackage
}

func NewContext() *Context {
	typeDictionary := NewTypeDictionary()
	c := &Context{
		packages:       []*Package{},
		defaultPackage: NewPackage("", typeDictionary),
		typeDictionary: typeDictionary,
	}
	return c
}
