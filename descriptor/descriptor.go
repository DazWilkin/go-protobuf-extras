package descriptor

import (
	"bytes"
	"fmt"

	"google.golang.org/protobuf/types/descriptorpb"
)

// ToString is a function that converts a FileDescriptorSet to a protobuf
func ToString(fds *descriptorpb.FileDescriptorSet) string {
	b := &bytes.Buffer{}
	for _, f := range fds.File {
		fmt.Fprintf(b, "%s\n", *f.Name)
		fmt.Fprintf(b, "package %s\n", *f.Package)
		for _, s := range f.Service {
			fmt.Fprintf(b, "service %s {\n", *s.Name)
			for _, m := range s.Method {
				fmt.Fprintf(b, "\trpc %s(%s) returns (%s)\n", *m.Name, *m.InputType, *m.OutputType)
			}
			fmt.Fprint(b, "}\n")
		}
		for _, m := range f.MessageType {
			fmt.Fprintf(b, "message %s {\n", *m.Name)
			for _, f := range m.Field {
				fmt.Fprintf(b, "\t%s %s = %d\n", *f.Type, *f.Name, *f.Number)
			}
			fmt.Fprint(b, "}\n")
		}
	}
	return b.String()
}
