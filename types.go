package gopdf

import "io"

// PDF struct for internal type reference.
type PDF struct {
	// Scalar value types
	Integer    int64
	Boolean    bool
	RealNumber float64
	Name       string // i.e. '/text'
	String     string // in the file as either '(...characters...)' or '<...hexadecimal character codes...>'
	// Container (Object) types
	Dictionary struct{}            // in the file as '<<...other objects...>>. These are always in pairs, a Name Obj followed by any other object type.
	Array      []struct{}          // a list of un-delimited objects separated by white space only where necessary
	Stream     map[struct{}]string // This is the most complex type. It's actually a Dictionary Obj mated with a string a bytes
}

// File struct for data and receivers
type File struct {
	File     io.ReaderAt
	Size     int64
	data     []byte
	Content  string
	Font     string
	FontSize float64
}

// Value struct
type Value struct {
	data interface{}
}
