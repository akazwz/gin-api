package request

type SetNotify struct {
	NotifyGap   int `json:"notify_gap"`
	NotifyCount int `json:"notify_count"`
}
