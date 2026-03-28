package ecs

import (
	"go.foxforensics.dev/flog/internal/types"
)

func MapEvent(s, src string) (log *Log, err error) {
	m, err := types.NewMap(s)

	if err != nil {
		return
	}

	log = NewLog(s, src, &Base{
		Timestamp: m.GetTime("Event/System/TimeCreated/@SystemTime"),
		Message:   m.GetString("Event/EventData/Data/#text"),
		Tags:      "EventLog",
		Labels: map[string]interface{}{
			"Channel": m.GetString("Event/System/Channel"),
			"Level":   m.GetInt64("Event/System/Level"),
			"Task":    m.GetInt64("Event/System/Task"),
		},
	})

	log.Event.Kind = "event"
	log.Event.Module = "EventLog"
	log.Event.Dataset = "EventLog." + log.Labels["Channel"].(string)
	log.Event.Severity = m.GetInt64("Event/System/Level")
	log.Event.ID = m.GetString("Event/System/EventRecordID")
	log.Event.Code = m.GetString("Event/System/EventID/#text")
	log.Event.Provider = m.GetString("Event/System/Provider/@Name")

	log.Host = &Host{
		Hostname: m.GetString("Event/System/Computer"),
	}

	log.User = &User{
		ID: m.GetString("Event/System/Security/@UserID"),
	}

	log.Process = &Process{
		PID: m.GetInt64("Event/System/Execution/@ProcessID"),
		Thread: &Thread{
			ID: m.GetInt64("Event/System/Execution/@ThreadID"),
		},
	}

	return
}
