package artifacts

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/f0x4n6/futils/pkg/sys"
)

func BaseFile(name string) string {
	b := filepath.Base(name)

	return strings.TrimSuffix(b, filepath.Ext(b))
}

func Copy(dir, src string) (err error) {
	dst := filepath.Join(dir, filepath.Base(src))

	b, err := os.ReadFile(src)

	if err != nil {
		return
	}

	err = os.WriteFile(dst, b, sys.MODE_FILE)

	if os.IsNotExist(err) {
		err = nil
	}

	return
}

func ConsumeJson(name string) (lines []string, err error) {
	f, err := os.Open(name)

	if err != nil {
		return
	}

	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		lines = append(lines, fs.Text())
	}

	_ = f.Close()

	err = os.Remove(name)

	return
}

func ConsumeCsv(name string) (lines []string, err error) {
	f, err := os.Open(name)

	if err != nil {
		return
	}

	rr, err := csv.NewReader(f).ReadAll()

	if len(rr) <= 1 {
		_ = f.Close()
		return
	}

	m := map[string]string{}

	for _, r := range rr[1:] {
		for i, c := range r {
			m[rr[0][i]] = c
		}

		b, err := json.Marshal(m)

		if err != nil {
			sys.Error(err)
			continue
		}

		lines = append(lines, string(b))
	}

	_ = f.Close()

	err = os.Remove(name)

	return
}
