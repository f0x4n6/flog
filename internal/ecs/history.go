package ecs

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"go.foxforensics.dev/flog/internal/artifacts"
)

func MapHistory(fu *artifacts.Url, src string) (log *Log, err error) {
	log = NewLog(fmt.Sprint(fu), src, &Base{
		Timestamp: time.Unix(fu.Time/1000000, 0).UTC(),
		Message:   fu.Url,
		Tags:      "History",
		Labels: map[string]interface{}{
			"Title": fu.Title,
		},
	})

	u, err := url.Parse(fu.Url)

	if err != nil {
		return log, nil
	}

	log.Url = &Url{
		Original: fu.Url,
		Full:     u.String(),
		Scheme:   u.Scheme,
		Domain:   u.Host,
		Path:     u.Path,
		Query:    u.RawQuery,
		Fragment: u.Fragment,
		Username: u.User.Username(),
	}

	log.Url.Password, _ = u.User.Password()

	_, p, err1 := net.SplitHostPort(u.Host)
	port, err2 := strconv.ParseInt(p, 10, 64)

	if err1 == nil || err2 == nil {
		log.Url.Port = port
	}

	return
}
