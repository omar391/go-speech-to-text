package models

type ApiResponse struct {
	IsSuccess bool        `json:"is_success"`
	Msg       string      `json:"msg"`
	Token     string      `json:"token"`
	Data      interface{} `json:"data"`
}

type DataResult struct {
	OriginalAudioFileName string
	OriginalAudioFilePath string
	Text                  string
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name for DataResult
func (DataResult) TableName() string {
	return "stt_data"
}
