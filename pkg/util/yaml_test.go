package util

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

type DummyContext struct {
	Name string
}

type DummyContextWithTag struct {
	Name string `json:"newName"`
}

func TestYAMLEncoder(t *testing.T) {
	testSets := map[string]struct {
		values   []interface{}
		expected string
	}{
		"single value": {
			values: []interface{}{
				DummyContext{Name: "clusterA"},
			},
			expected: `Name: clusterA
`,
		},
		"single value with json tag": {
			values: []interface{}{
				DummyContextWithTag{Name: "clusterA"},
			},
			expected: `newName: clusterA
`,
		},
		"multiple values": {
			values: []interface{}{
				DummyContext{Name: "clusterA"},
				DummyContextWithTag{Name: "clusterB"},
			},
			expected: `Name: clusterA
---
newName: clusterB
`,
		},
	}
	for name, testSet := range testSets {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			enc := NewYAMLEncoder(&buf)
			for _, v := range testSet.values {
				assert.NilError(t, enc.Encode(v))
			}
			assert.NilError(t, enc.Close())
			assert.Equal(t, testSet.expected, buf.String())
		})
	}
}
