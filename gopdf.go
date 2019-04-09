package gopdf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
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

// OpenFileHTTP func accepts *multipart.FileHeader and exposes multipart.File
func OpenFileHTTP(f *multipart.FileHeader) (*File, error) {
	file, err := f.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	newFile := &File{
		File: file,
		Size: f.Size,
		data: data,
	}

	return newFile, nil
}

// OpenLocal func
func OpenLocal(filePath string) (*File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		file.Close()
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}

	newFile := &File{
		File: file,
		Size: fi.Size(),
	}

	return newFile, nil

}

/*
 * Example:
 * f, err := gopdf.OpenFileHTTP(http.MultipartForm.File[fileName])
 * err := f.verifyPDF()
 * if err != nil { panic(err) }
 */
func (f *File) verifyPDF() error {
	buf := make([]byte, 10)
	f.File.ReadAt(buf, 0)
	if !bytes.HasPrefix(buf, []byte("%PDF-1.")) || buf[7] < '0' || buf[8] != '\r' && buf[8] != '\n' {
		return fmt.Errorf("This is not a PDF file: invalid header")
	}

	if !bytes.HasSuffix(buf, []byte("%%EOF")) {
		return fmt.Errorf("This is not a PDF file: missing EOF")
	}
	return nil
}
