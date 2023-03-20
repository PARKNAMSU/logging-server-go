package structure

type ClientData struct {
	Protocol string                 `json:"protocol"`
	Data     map[string]interface{} `json:"data"`
}
