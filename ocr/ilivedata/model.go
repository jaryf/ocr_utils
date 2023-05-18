package ilivedata

type OcrRes struct {
	ErrorCode  int `json:"errorCode"`
	Code       int `json:"code"`
	Result     int `json:"result"`
	ImageSpams []struct {
		Code   int `json:"code"`
		Result int `json:"result"`
		Tags   []struct {
			Tag        int    `json:"tag"`
			Level      int    `json:"level"`
			Confidence int    `json:"confidence"`
			TagName    string `json:"tagName"`
			TagNameEn  string `json:"tagNameEn"`
			SubTags    []struct {
				SubTag       int    `json:"subTag"`
				SubTagName   string `json:"subTagName"`
				SubTagNameEn string `json:"subTagNameEn"`
				Level        int    `json:"level"`
				Confidence   int    `json:"confidence"`
			} `json:"subTags"`
		} `json:"tags"`
	} `json:"imageSpams"`
	Ocr       []string      `json:"ocr"`
	Labels    []interface{} `json:"labels"`
	TaskId    string        `json:"taskId"`
	ExtraInfo struct {
		CartoonScore int           `json:"cartoonScore"`
		GenderResult []interface{} `json:"genderResult"`
		NumHuman     int           `json:"numHuman"`
		NumFace      int           `json:"numFace"`
	} `json:"extraInfo"`
	IsOcr int `json:"isOcr"`
}
