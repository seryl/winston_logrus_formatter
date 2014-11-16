package winston_logrus_formatter

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"time"
)

type WinstonJSONFormatter struct{}

func prefixFieldClashes(entry *log.Entry) {
	_, ok := entry.Data["t"]
	if ok {
		entry.Data["fields.t"] = entry.Data["t"]
	}

	_, ok = entry.Data["m"]
	if ok {
		entry.Data["fields.m"] = entry.Data["m"]
	}

	_, ok = entry.Data["s"]
	if ok {
		entry.Data["fields.s"] = entry.Data["s"]
	}
}

func (f *WinstonJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	prefixFieldClashes(entry)
	entry.Data["s"] = entry.Level.String()
	entry.Data["m"] = entry.Message
	entry.Data["t"] = entry.Time.Format(time.RFC3339)

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
