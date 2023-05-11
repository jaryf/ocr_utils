package xfyun

import (
	"testing"
)

const (
	APPID     = "xxxxxxxxxx"
	APISecret = "xxxxxxxxxx"
	APIKey    = "xxxxxxxxxx"
	xfWebAPI  = "https://api.xf-yun.com/v1/private/sf8e6aca1"
	xfHost    = "api.xf-yun.com"
	imgUrl    = "https://www.xiuren5.vip/uploadfile/202305/11/681223583.jpg"
	imgUrl2   = "https://www.xiuren5.vip/uploadfile/202305/11/F11223155.jpg"
	imgUrl3   = "https://www.xiuren5.vip/uploadfile/202305/11/5C1223262.jpg"
	imgPath   = "961223216.jpg"
)

var xf *XfOcr

func init() {
	xf = NewXfOcr(APPID, APISecret, APIKey, xfWebAPI, xfHost)
}

func TestXfOcr_ImgOcrXfFromUrl(t *testing.T) {
	xfOcr := NewXfOcr(APPID, APISecret, APIKey, xfWebAPI, xfHost)
	wordList, err := xfOcr.ImgOcrXfFromUrl(imgUrl2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(wordList)
}

func TestXfOcr_ImgOcrXfFromPath(t *testing.T) {
	wordList, err := xf.ImgOcrXfFromPath(imgPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(wordList)
}

func BenchmarkXfOcr_ImgOcrXfFromPath(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := xf.ImgOcrXfFromPath(imgPath)
		if err != nil {
			return
		}
	}
}
