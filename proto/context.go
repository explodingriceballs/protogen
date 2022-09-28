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
		c.typeDictionary.startScope(c.defaultPackage)
		c.typeDictionary.declareTypes(proto.ProtoBody)
		proto.Accept(c.defaultPackage)
		c.typeDictionary.endScope(c.defaultPackage)
		return
	}

	// Loop thru existing packages
	for _, pkg := range c.packages {
		if pkg.name == pkgName {
			c.typeDictionary.startScope(pkg)
			proto.Accept(pkg)
			c.typeDictionary.endScope(pkg)
			return
		}
	}

	// New a new package
	newPkg := NewPackage(pkgName, c.typeDictionary)
	c.typeDictionary.startScope(newPkg)
	c.typeDictionary.declareTypes(proto.ProtoBody)
	proto.Accept(newPkg)
	c.typeDictionary.endScope(newPkg)
	c.packages = append(c.packages, newPkg)
}

func (c *Context) Packages() []*Package {
	return c.packages
}

func (c *Context) Package(name string) *Package {
	for _, pkg := range c.packages {
		if pkg.name == name {
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
