// Code generated by protoc-gen-go-json. DO NOT EDIT.
package example

import "google.golang.org/protobuf/encoding/protojson"

func (p *Hello) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(p)
}

func (p *Hello) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, p)
}