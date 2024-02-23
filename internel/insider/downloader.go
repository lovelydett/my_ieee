// Package insider implements the downloader from different sources.

package insider

import (
	"filepath"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Downloader struct{}

func (d *Downloader) Process(jsonMap map[string]interface{}) {
	url := jsonMap["url"]

	// make HTTP request to get the page content
	res, err := http.Get(url.(string))
	if err != nil {
		panic(err)
	}

}

func saveFile(res *http.Response, format string) string {
	// Use the webpage title as the file title
	title := "test_file"
	nsec := time.Now().UnixNano()
	fileId := title + strconv.FormatInt(nsec, 10) + "." + format

	tmpPath := "/tmp"
	filePath := filepath.Join(tmpPath, fileId)
	dst, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, res.Body); err != nil {
		panic(err)
	}

}

func uploadToServer(fileId string) {

}
