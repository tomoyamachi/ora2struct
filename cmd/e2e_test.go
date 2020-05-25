package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tomoyamachi/dbscheme2struct/pkg/generator"
)

func TestE2E(t *testing.T) {
	tests := []struct {
		origin string
		golden string
	}{
		{origin: "./testdata/ddl/single.ddl", golden: "./testdata/golden/single.go.golden"},
	}

	for _, tt := range tests {

		nodes, err := parseFile(tt.origin, false)
		if err != nil {
			t.Errorf("unsupported error: %s", err)
			continue
		}

		var buf bytes.Buffer
		fmt.Println(len(nodes))
		if err = generator.Output(&buf, nodes, "models", ""); err != nil {
			t.Errorf("unsupported error: %s", err)
			continue
		}

		want, err := ioutil.ReadFile(tt.golden)
		if err != nil {
			t.Errorf("unsupported error: %s", err)
			continue
		}

		fmt.Println(buf)
		if buf.String() != string(want) {
			t.Errorf("expected %s, but %s", string(want), buf.String())
		}

	}
}
