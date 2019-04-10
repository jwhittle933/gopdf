package gopdf

import "io"

//Reader Interface
type Reader interface {
	ReadAt(buf []byte, pos int64) (n int, err error)
	Read(buf []byte) (n int, err error)
	Slice(n int) []byte
	Seek(off int, whence int) (ret int64, err error)
	ReadByte() (c byte, err error)
	UnreadByte() error
	Size() int64
}

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

// PdfReader struct
type PdfReader struct {
	File      string
	rdr       Reader
	StartXref int
	Xref      map[int]int
	Trailer   map[string][]byte
}
