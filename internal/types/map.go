package types

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/f0x4n6/futils/pkg/sys"
)

const (
	KeySep = "/"
	ValSep = "\n"
)

var (
	bom = []byte{0xEF, 0xBB, 0xBF}
)

type Map struct {
	c map[string]string
	o object
}

type object any

func NewMap(s string) (m *Map, err error) {
	m = &Map{
		c: make(map[string]string, 100),
	}

	b := bytes.TrimPrefix([]byte(s), bom)

	return m, json.Unmarshal(b, &m.o)
}

func (m *Map) GetString(keys ...string) (value string) {
	for _, key := range keys {
		if value, ok := m.c[key]; ok {
			return value
		}

		value = rget(m.o, strings.Split(key, KeySep))

		if len(value) > 0 {
			m.c[key] = value
			return
		}
	}

	return
}

func (m *Map) GetInt64(keys ...string) (value int64) {
	value, err := strconv.ParseInt(m.GetString(keys...), 10, 64)

	if err != nil {
		return -1 // default
	}

	return
}

func (m *Map) GetTime(keys ...string) (value time.Time) {
	const layout = "2006-01-02 15:04:05.9999999"

	value, err := time.Parse(layout, m.GetString(keys...))

	if err != nil {
		return time.UnixMicro(0) // default
	}

	return
}

func rget(o object, k []string) (s string) {
	switch v := o.(type) {
	case map[string]any:
		o, ok := v[k[0]]

		if !ok {
			return
		}

		if len(k) > 1 {
			return rget(o, k[1:])
		}

		s, ok = o.(string)

		if !ok {
			sys.Error("map: key too short")
			return
		}

		return

	case []any:
		var l []string

		for _, o = range v {
			l = append(l, rget(o, k))
		}

		return strings.Join(l, ValSep)

	case string: // abort for cut short keys
		return v

	case nil:
		return
	}

	sys.Error("map: type not supported")
	return
}
