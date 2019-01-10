package models

// MinDocRest ...
type MinDocRest struct {
	Token    string `json:"token"`
	Folder   string `json:"folder"`
	Title    string `json:"title"`
	Identify string `json:"identify"`
	TextMD   string `json:"textmd"`
	TextHTML string `json:"texthtml"`
}

type HtmlResult struct {
	output string
}

// NewMinDocRest ...
func NewMinDocRest() *MinDocRest {
	return &MinDocRest{}
}

// NewHtmlResult ...
func NewHtmlResult() *HtmlResult {
	return &HtmlResult{}
}

func (p *HtmlResult) Write(b []byte) (n int, err error) {
	p.output += string(b)
	return len(b), nil
}
