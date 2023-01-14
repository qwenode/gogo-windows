package process

import (
	"log"
	"strings"
	"testing"
)

func TestProcesses(t *testing.T) {
	tests := []struct {
		name    string
		want    []Process
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Processes()
			has := false
			for _, process := range got {
				if strings.Contains(process.Executable(), "explorer.exe") {
					has = true
					log.Printf("found:%v,PID:%d,PPID:%d", process.Executable(), process.Pid(), process.PPid())
					break
				}
			}
			if !has {
				t.Errorf("Processes() error = %v, wantErr %v", "not found", "explorer.exe")
				return
			}

		})
	}
}
