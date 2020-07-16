package random

import (
	"fmt"
	"log"
	"testing"

	"github.com/DazWilkin/protobuf-extras/example"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestRandom(t *testing.T) {
	files, err := protodesc.NewFiles(example.Descriptor)
	if err != nil {
		t.Error(err)
	}

	d, err := files.FindDescriptorByName(protoreflect.FullName(fmt.Sprintf("%s.%s", *example.Package, *example.Input)))
	if err != nil {
		log.Fatal(err)
	}

	// Assert it
	md, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		t.Error("Unable to assert to a MessageDescriptor")
	}

	m := Message(md)
	fd := md.Fields().ByName(protoreflect.Name("name"))
	v := m.Get(fd)
	got := v.String()
	want := "X"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
