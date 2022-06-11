# protoc-gen-go-json
Simple protoc plugin that fixes json marshaling issues with proto messages
by just adding `jsonpb.Marshal` and `jsonpb.Unmarshal` from [jsonpb](https://pkg.go.dev/github.com/golang/protobuf/jsonpb)

Example:
```go
// Code generated by protoc-gen-go-json. DO NOT EDIT.
package example

import (
	"bytes"
	"io"
	"github.com/golang/protobuf/jsonpb"
)

func (p *Hello) MarshalJSON() ([]byte, error) {
	var bytes bytes.Buffer
	err := (&jsonpb.Marshaler{}).Marshal(io.Writer(&bytes), p)
	return bytes.Bytes(), err
}

func (p *Hello) UnmarshalJSON(data []byte) error {
	return jsonpb.Unmarshal(bytes.NewBuffer(data), p)
}
```
So `json.Marshal` and `json.Unmarshal` uses protojson variant.
Since `jsonpb` supports all of the features of proto, so this plugin should too.


### How to install
```bash
go install github.com/elb3k/protoc-gen-go-json/cmd/protoc-gen-go-json
```


## Usage
Plugin is supported by `protoc` and `buf`


## Protoc usage

After the plugin is installed
```bash
protoc  --go-json_out=. --go-json_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative  ./example/example.proto
```

## Buf usage
Plugin can be added as any other buf plugin:
```yaml
plugins:
  # Other plugins (Go plugin is required)
  - name: protoc-gen-go-json
    path: protoc-gen-go-json
    out: example
    opt:
      - lang=go
      - paths=source_relative
```


Enjoy!