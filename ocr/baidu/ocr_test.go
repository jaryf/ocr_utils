package baidu

import (
	"log"
	"testing"
)

const (
	APPID     = "xxx"
	APIKey    = "xxx"
	APISecret = "xxx"
	imgUrl    = "https://www.xiuren5.vip/uploadfile/202305/11/681223583.jpg"
	imgUrl2   = "https://www.xiuren5.vip/uploadfile/202305/11/F11223155.jpg"
	imgUrl3   = "https://www.xiuren5.vip/uploadfile/202305/11/5C1223262.jpg"
	imgPath   = "961223216.jpg"
)

var bd *BdOcr

func init() {
	bd = NewBdOcr(APIKey, APISecret)
	err := bd.InitAccessToken()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func TestBdOcr_ImgOcr(t *testing.T) {
	res, err := bd.ImgOcr(&GeneralBasicReq{
		Url:             imgUrl2,
		LanguageType:    CHN_ENG,
		DetectDirection: true,
		DetectLanguage:  true,
		Paragraph:       true,
		Probability:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
