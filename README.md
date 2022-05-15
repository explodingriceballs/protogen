# Protogen

![go report](https://goreportcard.com/badge/github.com/explodingriceballs/protogen)

A tool for generating arbitrary code based on protobuf files using go templates & typescript.

## How to use

```
Usage:
  protogen [flags]

Flags:
      --generator string   directory containing the generator
  -h, --help               help for protogen
      --out string         output directory
      --source string      source directory / files
      --include string     include directories
```

### Generating code

Invoking the following command: `protogen --generator ./mygenerator --out ./pkg/service --source ./api`
will:

1. Read all `*.proto` files in the `./api` directory and parse them
2. Execute the `index.tsx` generator class in the `./mygenerator` directory
3. Place all files generated in the `./pkg/service` directory

### Writing generators

The `examples` directory contains examples of generators.