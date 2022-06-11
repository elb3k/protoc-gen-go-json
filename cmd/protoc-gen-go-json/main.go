package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	code = `
func (p *%[1]s) MarshalJSON() ([]byte, error) {
	var bytes bytes.Buffer
	err := (&jsonpb.Marshaler{}).Marshal(io.Writer(&bytes), p)
	return bytes.Bytes(), err
}

func (p *%[1]s) UnmarshalJSON(data []byte) error {
	return jsonpb.Unmarshal(bytes.NewBuffer(data), p)
}
`
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_protojson.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-json. DO NOT EDIT.")
	g.P("package ", file.GoPackageName)
	g.P(`import (`)
	g.P(`"bytes"`)
	g.P(`"io"`)
	g.P(`"github.com/golang/protobuf/jsonpb"`)
	g.P(")")
	for _, msg := range file.Messages {
		g.P(fmt.Sprintf(code, msg.GoIdent.GoName))
	}
	return g
}
