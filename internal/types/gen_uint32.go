/*
 This file was autogenerated via
 -------------------------------------------------------------
 gen-builtin --declarable --native --type Uint32 --name uint32
 -------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Uint32("")

// Uint32 represents field of type uint32
type Uint32 string

// Name returns field name
func (i Uint32) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Uint32) TypeName() string {
	return "uint32"
}

// Register registers a field
func (i Uint32) Register(registrator FieldRegistrator) {
	registrator.AddUint32(i.Name())
}

// GoName returns Go's representation of this field's type
func (i Uint32) GoName() string {
	return "uint32"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["uint32"] = func(fieldName string) Field {
		return Uint32(fieldName)
	}
	if natives == nil {
		natives = map[string]struct{}{}
	}
	natives["uint32"] = struct{}{}
	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["uint32"] = struct{}{}

}
