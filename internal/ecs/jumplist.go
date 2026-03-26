package ecs

import (
	"path/filepath"
	"strings"

	"github.com/f0x4n6/flog/internal/types"
)

func MapJumpList(s, src string) (log *Log, err error) {
	m, err := types.NewMap(s)

	if err != nil {
		return
	}

	exe := m.GetString("LocalPath", "Path")
	arg := m.GetString("Arguments")

	log = NewLog(s, src, &Base{
		Timestamp: m.GetTime("LastModified", "TargetAccessed"),
		Message:   strings.TrimSpace(exe + " " + arg),
		Tags:      "JumpList",
		Labels:    make(map[string]interface{}, 1),
	})

	if strings.Contains(s, "DestListVersion") {
		log.Labels["Destination"] = "automatic"
	} else {
		log.Labels["Destination"] = "custom"
	}

	log.Host = &Host{
		Hostname: m.GetString("Hostname", "MachineID"),
		MAC:      m.GetString("MacAddress", "MachineMACAddress"),
	}

	var args []string

	if len(arg) > 0 {
		args = strings.Split(arg, " ")
	}

	log.Process = &Process{
		EntityID:         m.GetString("AppId"),
		Name:             filepath.Base(exe),
		Title:            m.GetString("AppIdDescription"),
		Executable:       exe,
		Args:             args,
		ArgsCount:        int64(len(args)),
		CommandLine:      log.Message,
		WorkingDirectory: m.GetString("WorkingDirectory"),
	}

	return
}
