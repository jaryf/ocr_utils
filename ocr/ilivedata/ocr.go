package ilivedata

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	endpointURL  = "https://isafe.ilivedata.com/api/v1/image/check"
	endpointHost = "isafe.ilivedata.com"
	endpointPath = "/api/v1/image/check"
)

type Ocr struct {
	ProjectID string
	SecretKey string
}

func NewOcr(projectID string, secretKey string) *Ocr {
	return &Ocr{
		ProjectID: projectID,
		SecretKey: secretKey,
	}
}

func (m *Ocr) ImgCheck(image string, imageType int, userID string) (res OcrRes, err error) {
	var now = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	var parameters = map[string]interface{}{
		"type":   imageType,
		"image":  image,
		"userId": userID,
	}
	var queryBody []byte
	queryBody, err = json.Marshal(parameters)
	if err != nil {
		return
	}
	var preparedString = []string{
		"POST",
		endpointHost,
		endpointPath,
		m.sha256AndHexEncode(string(queryBody)),
		"X-AppId:" + m.ProjectID,
		"X-TimeStamp:" + now,
	}
	var stringToSign = strings.Join(preparedString, "\n")
	var signature = m.signAndBase64Encode(stringToSign, m.SecretKey)
	return m.request(string(queryBody), signature, now)
}

func (m *Ocr) signAndBase64Encode(data string, secrectKey string) string {
	var mac = hmac.New(sha256.New, []byte(secrectKey))
	mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (m *Ocr) sha256AndHexEncode(data string) string {
	var sha256Hash = sha256.New()
	sha256Hash.Write([]byte(data))
	return hex.EncodeToString(sha256Hash.Sum(nil))
}

func (m *Ocr) request(body string, signature string, timeStamp string) (res OcrRes, err error) {
	var httpClient = http.Client{}
	apiReq, _ := http.NewRequest("POST", endpointURL, strings.NewReader(body))
	apiReq.Header.Set("X-AppId", m.ProjectID)
	apiReq.Header.Set("X-TimeStamp", timeStamp)
	apiReq.Header.Set("Authorization", signature)
	apiReq.Header.Set("Content-Type", "application/json")
	apiReq.Header.Set("User-Agent", "Golang_HTTP_Client/1.0")

	response, err := httpClient.Do(apiReq)

	if err != nil {
		return
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &res)
	if err != nil {
		return
	}
	return
}
