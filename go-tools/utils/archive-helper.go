package utils

// import (
// 	"archive/zip"
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// // UnzipRezipAndRunOperation starts by deleting the destination directory if it exists,
// // then it goes ahead an unzips the contents into the destination directory
// // once that is done it runs the operation func on the destination folder
// // lastly it rezips the folder back to compress.zip
// func UnzipRezipAndRunOperation(src, dest string, operation func() error) error {
// 	var err error
// 	if FolderExists(dest) {
// 		err = os.RemoveAll(dest)

// 		if err != nil {
// 			WriteError(fmt.Sprintf("failed to delete the destination directory \"%s\": %s", dest, err))
// 		}
// 	}

// 	err = Unzip(src, dest)
// 	if err != nil {
// 		WriteError(fmt.Sprintf("failed to unzip \"%s\": %s", src, err))
// 	}

// 	err = operation()
// 	if err != nil {
// 		WriteError(fmt.Sprintf("failed to run the operation on the unzipped source \"%s\": %s", src, err))
// 	}

// 	err = Rezip(dest, "compress.zip")
// 	if err != nil {
// 		WriteError(fmt.Sprintf("failed to rezip content for source \"%s\": %s", src, err))
// 	}

// 	err = os.RemoveAll(dest)
// 	if err != nil {
// 		WriteError(fmt.Sprintf("failed to cleanup the destination directory \"%s\": %s", dest, err))
// 	}

// 	return nil
// }

// // Unzip is based on https://stackoverflow.com/a/24792688
// func Unzip(src, dest string) error {
// 	r, err := zip.OpenReader(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		if err := r.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	os.MkdirAll(dest, 0755)

// 	// Closure to address file descriptors issue with all the deferred .Close() methods
// 	extractAndWriteFile := func(f *zip.File) error {
// 		rc, err := f.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer func() {
// 			if err := rc.Close(); err != nil {
// 				panic(err)
// 			}
// 		}()

// 		path := filepath.Join(dest, f.Name)

// 		// Check for ZipSlip (Directory traversal)
// 		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
// 			return fmt.Errorf("illegal file path: %s", path)
// 		}

// 		if f.FileInfo().IsDir() {
// 			os.MkdirAll(path, f.Mode())
// 		} else {
// 			os.MkdirAll(filepath.Dir(path), f.Mode())
// 			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
// 			if err != nil {
// 				return err
// 			}
// 			defer func() {
// 				if err := f.Close(); err != nil {
// 					panic(err)
// 				}
// 			}()

// 			_, err = io.Copy(f, rc)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	}

// 	for _, f := range r.File {
// 		err := extractAndWriteFile(f)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // Rezip is based on https://stackoverflow.com/a/63233911
// func Rezip(src, dest string) error {
// 	file, err := os.Create(dest)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	w := zip.NewWriter(file)
// 	defer w.Close()

// 	walker := func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		// skip empty directories
// 		if info.IsDir() {
// 			return nil
// 		}

// 		file, err := os.Open(path)
// 		if err != nil {
// 			return err
// 		}
// 		defer file.Close()

// 		// Ensure that `path` is not absolute; it should not start with "/".
// 		// This snippet happens to work because I don't use
// 		// absolute paths, but ensure your real-world code
// 		// transforms path into a zip-root relative path.
// 		f, err := w.Create(path)
// 		if err != nil {
// 			return err
// 		}

// 		_, err = io.Copy(f, file)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// 	err = filepath.Walk(src, walker)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
