package atfield

import (
	"os"
	"path/filepath"
)

func absolutePath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = filepath.Join(wd, path)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errInvalidPath
	}
	return path, nil
}

func (e *engine) SetInfileDir(dir string) {
	e.errorIntercept(func() {
		infileDir, err := absolutePath(dir)
		if err != nil {
			e.err = err
		} else {
			e.infileDir = infileDir
		}
	})
}

func (e *engine) SetoutfileDir(dir string) {
	e.errorIntercept(func() {
		outfileDir, err := absolutePath(dir)
		if err != nil {
			e.err = err
		} else {
			e.outfileDir = outfileDir
		}
	})
}
