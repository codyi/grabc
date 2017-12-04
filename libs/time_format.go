package libs

import (
	"time"
)

func UnixTimeFormat(unix int32, Format string) string {
	tm := time.Unix(int64(unix), 0)
	return tm.Format(Format)
}
