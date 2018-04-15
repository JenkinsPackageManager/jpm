package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// PostFile sends a given file to some URL via HTTP POST
func PostFile(url, filePath string) error {
	data := &bytes.Buffer{}
	bodyWritter := multipart.NewWriter(data)

	fileWritter, err := bodyWritter.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(fileWritter, file)
	if err != nil {
		return err
	}

	bodyWritter.Close()

	resp, err := http.Post(url, bodyWritter.FormDataContentType(), data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

// FetchFile requests a file and saves its contents to a given file path
func FetchFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0644)
}
