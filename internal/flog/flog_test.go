package flog

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/f0x4n6/futils/pkg/hash"
	"github.com/f0x4n6/futils/pkg/sys"
	"github.com/f0x4n6/futils/pkg/zip"
)

func TestMain(m *testing.M) {
	sys.Progress = nil

	_ = os.Setenv("EZTOOLS", "../../bin")
	os.Exit(m.Run())
}

func TestLogEvent(t *testing.T) {
	cases := []struct {
		name, data, file string
	}{
		{
			name: "Test log event",
			data: filepath.Join("..", "testdata", "system.zip"),
			file: "System.evtx",
		},
	}

	for _, tt := range cases {
		tmp, _ := os.MkdirTemp(os.TempDir(), "log")

		err := zip.Unzip(tt.data, tmp)

		if err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			l, err := LogEvent(filepath.Join(tmp, tt.file), tmp, true)

			if err != nil {
				t.Fatal(err)
			}

			if len(l) != 1 {
				t.Fatal("file count differs")
			}

			b, err := os.ReadFile(l[0])

			if err != nil {
				t.Fatal(err)
			}

			if !json.Valid(b) {
				t.Fatal("invalid json")
			}
		})
	}
}

func TestLogJumpList(t *testing.T) {
	cases := []struct {
		name, data, file string
	}{
		{
			name: "Test log automatic jumplist",
			data: filepath.Join("..", "testdata", "user.zip"),
			file: "0.automaticDestinations-ms",
		},
		{
			name: "Test log custom jumplist",
			data: filepath.Join("..", "testdata", "user.zip"),
			file: "0.customDestinations-ms",
		},
	}

	for _, tt := range cases {
		tmp, _ := os.MkdirTemp(os.TempDir(), "log")

		err := zip.Unzip(tt.data, tmp)

		if err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			l, err := LogJumpList(filepath.Join(tmp, tt.file), tmp, true)

			if err != nil {
				t.Fatal(err)
			}

			if len(l) != 2 {
				t.Fatal("file count differs")
			}

			b, err := os.ReadFile(l[0])

			if err != nil {
				t.Fatal(err)
			}

			if !json.Valid(b) {
				t.Fatal("invalid json")
			}
		})
	}
}

func TestLogShellbag(t *testing.T) {
	cases := []struct {
		name, data, file string
	}{
		{
			name: "Test log usrclass.dat",
			data: filepath.Join("..", "testdata", "user.zip"),
			file: "UsrClass.dat",
		},
	}

	for _, tt := range cases {
		tmp, _ := os.MkdirTemp(os.TempDir(), "log")

		err := zip.Unzip(tt.data, tmp)

		if err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			l, err := LogShellBag(filepath.Join(tmp, tt.file), tmp, true)

			if err != nil {
				t.Fatal(err)
			}

			if len(l) == 0 {
				t.Fatal("file count zero")
			}

			b, err := os.ReadFile(l[0])

			if err != nil {
				t.Fatal(err)
			}

			if !json.Valid(b) {
				t.Fatal("invalid json")
			}
		})
	}
}

func TestLogHistory(t *testing.T) {
	cases := []struct {
		name, data, file string
	}{
		{
			name: "Test log history",
			data: filepath.Join("..", "testdata", "user.zip"),
			file: "History",
		},
		{
			name: "Test log places.sqlite",
			data: filepath.Join("..", "testdata", "user.zip"),
			file: "places.sqlite",
		},
	}

	for _, tt := range cases {
		tmp, _ := os.MkdirTemp(os.TempDir(), "log")

		err := zip.Unzip(tt.data, tmp)

		if err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			l, err := LogHistory(filepath.Join(tmp, tt.file), tmp, true)

			if err != nil {
				t.Fatal(err)
			}

			if len(l) == 0 {
				t.Fatal("file count zero")
			}

			b, err := os.ReadFile(l[0])

			if err != nil {
				t.Fatal(err)
			}

			if !json.Valid(b) {
				t.Fatal("invalid json")
			}
		})
	}
}

func TestStripHash(t *testing.T) {
	cases := []struct {
		name, file, hash string
	}{
		{
			name: "Test StripHash",
			file: "test",
			hash: "68ac906495480a3404beee4874ed853a037a7a8f",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f := StripHash([]string{
				tt.hash + hash.Sep + tt.file,
			})

			if len(f) != 1 {
				t.Fatal("file count wrong")
			}

			if f[0] != tt.file {
				t.Fatal("hash not stripped")
			}
		})
	}
}
