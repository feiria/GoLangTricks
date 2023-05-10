package file_download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(filepath string, url string) error {

	// get data from remote
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// create empty file in local
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	// write data of body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)

	return n, nil
}

func (wc *WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %d complete", wc.Total)
}

func DownloadBigFile(filepath string, url string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	err = out.Close()
	if err != nil {
		return err
	}
	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")
	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}
