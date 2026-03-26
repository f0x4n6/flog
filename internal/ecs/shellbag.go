package ecs

import (
	"github.com/f0x4n6/flog/internal/types"
)

func MapShellBag(s, src string) (log *Log, err error) {
	m, err := types.NewMap(s)

	if err != nil {
		return
	}

	log = NewLog(s, src, &Base{
		Timestamp: m.GetTime("LastInteracted", "LastWriteTime"),
		Message:   m.GetString("AbsolutePath"),
		Tags:      "ShellBag",
	})

	log.Registry = &Registry{
		Hive: "HKU",
	}

	return
}
