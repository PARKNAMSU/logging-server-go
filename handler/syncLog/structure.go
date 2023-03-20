package synclog

type SyncLogData struct {
	Alloc       int    `json:"alloc"`
	TotalAlloc  int    `json:"totalAlloc"`
	Sys         int    `json:"sys"`
	NumGC       int    `json:"numGC"`
	RunningTime int    `json:"runningTime"`
	Branch      string `json:"branch"`
}
