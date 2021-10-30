package response

import "time"

type Notify struct {
	NotifyGap        int       `json:"notify_gap"`
	NotifyCount      int       `json:"notify_count"`
	AllNotifyCount   int       `json:"all_notify_count"`
	LastNotify       time.Time `json:"last_notify"`
	NotifyLimitCount int       `json:"notify_limit_count"`
}
