package tools

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("EZTOOLS", "../../bin")
	os.Exit(m.Run())
}

func TestPath(t *testing.T) {
	cases := []struct {
		name, tool string
	}{
		{
			name: "Test for EvtxECmd",
			tool: "EvtxECmd.dll",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Path(tt.tool)

			if err != nil {
				t.Fatal(err)
			}

			if len(p) == 0 {
				t.Fatal(tt.tool + " not found")
			}

			if _, err := os.Stat(p); os.IsNotExist(err) {
				t.Fatal(tt.tool + " not found")
			}
		})
	}
}
