package utils

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
)

func GetImgBase64FromUrl(imgUrl string) (imgBase64 string, err error) {
	var (
		resp     *http.Response
		respByte []byte
	)
	resp, err = http.DefaultClient.Get(imgUrl)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("获取图片响应码非200")
		return
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			return
		}
	}()
	respByte, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	imgBase64 = base64.StdEncoding.EncodeToString(respByte)
	return
}

func GetImgBase64FromPath(imgPath string) (imgBase64 string, err error) {
	var (
		f     *os.File
		fByte []byte
	)
	if f, err = os.Open(imgPath); err != nil {
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			return
		}
	}()
	if fByte, err = io.ReadAll(f); err != nil {
		return
	}
	imgBase64 = base64.StdEncoding.EncodeToString(fByte)
	return
}
