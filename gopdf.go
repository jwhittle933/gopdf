package gopdf

import (
	"io/ioutil"
	"mime/multipart"
)

/*
 !! http://www.planetpdf.com/developer/article.asp?ContentID=navigating_the_internal_struct&page=0
 * See source above for authorial credit of some of the following info.
 *
 * Credit is also due to github.com/rsc/pdf as well as a fork from this repo
 * at github.com/ledongthuc/pdf as mental aids and idea completion.
 *
 * The PDF file format is text with some binary data mixed in.
*/

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
	data     []byte
	Content  string
	Font     string
	FontSize float64
}

// ExposeFileHTTP func accepts *multipart.FileHeader and exposes multipart.File
func ExposeFileHTTP(f *multipart.FileHeader) ([]byte, error) {
	file, err := f.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
