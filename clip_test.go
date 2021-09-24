package clip

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	value    interface{}
	expected string
}

type test1 struct {
	Foo string `clip:"foo"`
	Bar string `clip:"bar"`
}

type test2 struct {
	Foo []string `clip:"foo"`
}

type test3 struct {
	Foo string            `clip:"foo"`
	Map map[string]string `clip:"map"`
	Bar string            `clip:"bar"`
}

var cases = []testCase{
	{
		value: "abcd",
		expected: `
abcd
`,
	},
	{
		value: []test1{{Foo: "abcd", Bar: "efgh"}},
		expected: `
- foo: abcd
  bar: efgh
`,
	},
	{
		value: test2{Foo: []string{"abcd", "efgh"}},
		expected: `
foo:
  - abcd
  - efgh
`,
	},
	{
		value: test3{Foo: "abcd", Bar: "efgh", Map: nil},
		expected: `
foo: abcd
map:
bar: efgh
`,
	},
	{
		value: true,
		expected: `
true
`,
	},
	{
		value: false,
		expected: `
false
`,
	},
}

func TestCases(t *testing.T) {
	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			buf := &bytes.Buffer{}
			Fprint(buf, tc.value)
			require.Equal(t, tc.expected[1:], buf.String())
		})
	}
}
