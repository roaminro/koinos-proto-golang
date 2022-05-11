package canonical

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// Marshal ensures the object is serialized in a canonical way
func Marshal(m proto.Message) ([]byte, error) {
	if containsMap(m.ProtoReflect()) {
		return nil, fmt.Errorf("%w: the given object contains a map", ErrNoCanonical)
	}

	return proto.Marshal(m)
}

func containsMap(pr protoreflect.Message) bool {
	fields := pr.Descriptor().Fields()

	// Iterate over the fields
	for i := 0; i < fields.Len(); i++ {
		fdesc := fields.Get(i)

		// If the field is a map we're done here
		if fdesc.IsMap() {
			return true
		}

		// Recurse into the field if it is a message
		if fdesc.Kind() == protoreflect.MessageKind {
			var message protoreflect.Message

			if !fdesc.IsList() {
				message = pr.Get(fdesc).Message()
			} else if pr.Get(fdesc).List().Len() > 0 {
				message = pr.Get(fdesc).List().Get(0).Message()
			} else {
				continue
			}

			if containsMap(message) {
				return true
			}
		}
	}

	return false
}
