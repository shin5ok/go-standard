package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	ulid "github.com/oklog/ulid/v2"
)

func New(name string) params {
	p := params{}
	p.Name = name
	return p
}

func (p params) RunCore(ctx context.Context) io.Reader {
	p.Id = ulid.Make().String()
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Errorf("Json Encode failure: %w", err)
	}
	return bytes.NewReader(jsonData)
}
