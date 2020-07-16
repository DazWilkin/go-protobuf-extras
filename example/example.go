package example

import "google.golang.org/protobuf/types/descriptorpb"

func intptr(i int32) *int32   { return &i }
func strptr(s string) *string { return &s }
func typptr(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}

var (
	Package    = strptr("package")
	Service    = strptr("service")
	Input      = strptr("input")
	Output     = strptr("output")
	Descriptor = &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{
			{
				Name:    strptr("name"),
				Package: Package,
				Syntax:  strptr("proto3"),
				Service: []*descriptorpb.ServiceDescriptorProto{
					{
						Name: Service,
						Method: []*descriptorpb.MethodDescriptorProto{
							{
								Name:       strptr("name"),
								InputType:  strptr("input"),
								OutputType: strptr("output"),
							},
						},
					},
				},
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: Input,
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Type:   typptr(descriptorpb.FieldDescriptorProto_TYPE_STRING),
								Name:   strptr("name"),
								Number: intptr(1),
							},
						},
					},
					{
						Name: Output,
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Type:   typptr(descriptorpb.FieldDescriptorProto_TYPE_STRING),
								Name:   strptr("name"),
								Number: intptr(1),
							},
						},
					},
				},
			},
		},
	}
)
