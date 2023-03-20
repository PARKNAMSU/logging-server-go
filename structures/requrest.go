package structures

type ProfFileData struct {
	MaxIdx int `json:"maxIdx"`
	NowIdx int `json:"nowIdx"`
	ProfType string `json:"profType"`
	FileByte []byte `json:"fileByte"`
	Uuid string `json:"uuid"`
}