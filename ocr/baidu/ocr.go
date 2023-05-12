package baidu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BdOcr struct {
	apiKey      string
	apiSecret   string
	accessToken string
	h           http.Client
}

func NewBdOcr(apiKey, apiSecret string) *BdOcr {
	return &BdOcr{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		h: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (m *BdOcr) SetConnectionTimeoutInMillis(timeOut time.Duration) {
	m.h.Timeout = timeOut
}

func (m *BdOcr) InitAccessToken() error {
	var token, err = m.GetAccessToken()
	if err != nil {
		return err
	}
	m.accessToken = token
	return err
}

func (m *BdOcr) ImgOcr(req *GeneralBasicReq) (res GeneralBasicResult, err error) {
	reqValues := url.Values{}
	if req.Image != "" {
		reqValues.Add("image", req.Image)
	}
	if req.Url != "" {
		reqValues.Add("url", req.Url)
	}
	if req.PdfFile != "" {
		reqValues.Add("pdf_file", req.PdfFile)
	}
	if req.PdfFileNum != "" {
		reqValues.Add("pdf_file_num", req.PdfFileNum)
	}
	if req.LanguageType != "" {
		reqValues.Add("language_type", string(req.LanguageType))
	}
	if req.DetectLanguage {
		reqValues.Add("detect_language", "true")
	}
	if req.DetectDirection {
		reqValues.Add("detect_direction", "true")
	}
	if req.Paragraph {
		reqValues.Add("paragraph", "true")
	}
	if req.Probability {
		reqValues.Add("probability", "true")
	}
	reqStr := reqValues.Encode()
	if m.accessToken == "" {
		err = fmt.Errorf("accessToken is empty")
		return
	}
	apiReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?access_token=%s", GeneralBasicUrl, m.accessToken), strings.NewReader(reqStr))
	if err != nil {
		return
	}
	apiReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := m.h.Do(apiReq)
	if err != nil {
		return
	}
	defer func() {
		respErr := resp.Body.Close()
		if respErr != nil {
			return
		}
	}()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("statusCode != 200, response is %v", string(respBytes))
		return
	}
	err = json.Unmarshal(respBytes, &res)
	return
}

func (m *BdOcr) GetAccessToken() (token string, err error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(AccessTokenUrl, m.apiKey, m.apiSecret), nil)
	if err != nil {
		return
	}
	resp, err := m.h.Do(req)
	if err != nil {
		return
	}
	defer func() {
		respErr := resp.Body.Close()
		if respErr != nil {
			return
		}
	}()
	var res AccessTokenResult
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("statusCode != 200, response is %v", string(respBytes))
		return
	}
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return
	}
	token = res.AccessToken
	return
}
