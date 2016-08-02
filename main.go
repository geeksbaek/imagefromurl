package imagesaver

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Save 함수는 지정된 경로로 해당 URL의 이미지를 다운로드 합니다.
func Save(path, URL string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

// Base64 함수는 해당 URL의 이미지를 Base64로 인코딩합니다.
func Base64(URL string) (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(body), nil
}
