package descriptor

import (
	"testing"

	"github.com/DazWilkin/go-protobuf-extras/example"
)

func TestToString(t *testing.T) {
	got := ToString(example.Descriptor)
	want := `name
package package
service service {
	rpc name(input) returns (output)
}
message input {
	TYPE_STRING name = 1
}
message output {
	TYPE_STRING name = 1
}
`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s\n", got, want)
	}
}
