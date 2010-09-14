package fig

import "io"
import "io/ioutil"
import "os"

type FileSystem interface {
	OpenReader(path string) (io.ReadCloser, os.Error)
	OpenWriter(path string) (io.WriteCloser, os.Error)
}

func ReadFile(fs FileSystem, path string) ([]byte, os.Error) {
	r, err := fs.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func WriteFile(fs FileSystem, path string, data []byte) os.Error {
	w, err := fs.OpenWriter(path)
	if err != nil {
		return err
	}
	defer w.Close()
	n, err := w.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		panic("unable to write entire file")
	}
	return nil
}

