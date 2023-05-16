package baidu

const (
	CHN_ENG LanguageType = "CHN_ENG" // 中英文混合
	ENG     LanguageType = "ENG"     // 英文
	JAP     LanguageType = "JAP"     // 日语
	KOR     LanguageType = "KOR"     // 韩语
	FRE     LanguageType = "FRE"     // 法语
	SPA     LanguageType = "SPA"     // 西班牙语
	POR     LanguageType = "POR"     // 葡萄牙语
	GER     LanguageType = "GER"     // 德语
	ITA     LanguageType = "ITA"     // 意大利语
	RUS     LanguageType = "RUS"     // 俄语
)

type AccessTokenResult struct {
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	SessionSecret    string `json:"session_secret"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type LanguageType string

// GeneralBasicReq 通用文字识别（标准版） https://cloud.baidu.com/doc/OCR/s/zk3h7xz52?_=1683880636284#%E8%AF%B7%E6%B1%82%E8%AF%B4%E6%98%8E
type GeneralBasicReq struct {
	Image           string       `json:"image,omitempty"`            // 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px，支持jpg/jpeg/png/bmp格式 优先级：image > url > pdf_file，当image字段存在时，url、pdf_file字段失效
	Url             string       `json:"url,omitempty"`              // 图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px，支持jpg/jpeg/png/bmp格式 优先级：image > url > pdf_file，当image字段存在时，url、pdf_file字段失效
	PdfFile         string       `json:"pdf_file,omitempty"`         // PDF文件，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px
	PdfFileNum      string       `json:"pdf_file_num,omitempty"`     // 需要识别的PDF文件的对应页码，当 pdf_file 参数有效时，识别传入页码的对应页面内容，若不传入，则默认识别第 1 页
	LanguageType    LanguageType `json:"language_type,omitempty"`    // 识别语言类型，默认为CHN_ENG
	DetectDirection bool         `json:"detect_direction,omitempty"` // 是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。
	DetectLanguage  bool         `json:"detect_language,omitempty"`  // 是否检测语言，默认不检测，即：false。当前支持中文、英语、日语、韩语
	Paragraph       bool         `json:"paragraph,omitempty"`        // 是否输出段落信息
	Probability     bool         `json:"probability,omitempty"`      // 是否返回识别结果中每一行的置信度
}

type FreeImgOcrResult struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
	Data  struct {
		WordsResult []struct {
			Words string `json:"words"`
		} `json:"words_result"`
		WordsResultNum int    `json:"words_result_num"`
		LogId          string `json:"log_id"`
	} `json:"data"`
}

type GeneralBasicResult struct {
	ParagraphsResult []struct {
		WordsResultIdx []int `json:"words_result_idx"`
	} `json:"paragraphs_result"`
	ParagraphsResultNum int `json:"paragraphs_result_num"`
	Direction           int `json:"direction"`
	Language            int `json:"language"`
	WordsResult         []struct {
		Probability struct {
			Average  float64 `json:"average"`
			Min      float64 `json:"min"`
			Variance float64 `json:"variance"`
		} `json:"probability"`
		Words string `json:"words"`
	} `json:"words_result"`
	WordsResultNum int   `json:"words_result_num"`
	LogId          int64 `json:"log_id"`
}
