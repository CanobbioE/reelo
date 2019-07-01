package dto

// UploadInfo represents the information associated w/
// an uploaded ranking file.
type UploadInfo struct {
	Category string `json:category`
	Year     string `json:year`
	IsParis  bool   `json:isParis`
	Format   string `json:format`
	Token    string `json:token`
	Start    string `json:start`
	End      string `json:end`
}
