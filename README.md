# Go SDK for some Protobuf APIv2 "Extras"

### `ToString(*descriptorpb.FileDescriptorSet) string`

Given a `descriptorpb.FileDescriptorSet` generates something resembling the original protobuf (source) files.


### `Message(protoreflect.MessageDescriptor) *dynamicpb.Message`

Given a `protoreflect.MessageDescriptor` generates a randomly constructed message

Currently supports:

|Type|
|----|
|protoreflect.Int32Kind|
|protoreflect.FloatKind|
|protoreflect.MessageKind|
|protoreflect.StringKind|

Laziness means I've not yet implemented the others ;-)
