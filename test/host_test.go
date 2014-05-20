package test

import (
	"fmt"
	"testing"

	"github.com/shirou/gopsutil"
)

func TestHostInfo(t *testing.T) {
	v, err := gopsutil.HostInfo()
	if err != nil {
		t.Errorf("error %v", err)
	}
	empty := &gopsutil.HostInfoStat{}
	if v == empty {
		t.Errorf("Could not get hostinfo %v", v)
	}
}

func TestBoot_time(t *testing.T) {
	v, err := gopsutil.BootTime()
	if err != nil {
		t.Errorf("error %v", err)
	}
	if v == 0 {
		t.Errorf("Could not boot time %v", v)
	}
}

func TestUsers(t *testing.T) {
	v, err := gopsutil.Users()
	if err != nil {
		t.Errorf("error %v", err)
	}
	empty := gopsutil.UserStat{}
	for _, u := range v {
		if u == empty {
			t.Errorf("Could not Users %v", v)
		}
	}
}

func TestHostInfoStat_String(t *testing.T) {
	v := gopsutil.HostInfoStat{
		Hostname: "test",
		Uptime:   3000,
		Procs:    100,
		OS:       "linux",
		Platform: "ubuntu",
	}
	e := `{"hostname":"test","uptime":3000,"procs":100,"os":"linux","platform":"ubuntu","platformFamily":"","platformVersion":"","virtualizationSystem":"","virtualizationRole":""}`
	if e != fmt.Sprintf("%v", v) {
		t.Errorf("HostInfoStat string is invalid: %v", v)
	}
}

func TestUserStat_String(t *testing.T) {
	v := gopsutil.UserStat{
		User:     "user",
		Terminal: "term",
		Host:     "host",
		Started:  100,
	}
	e := `{"user":"user","terminal":"term","host":"host","started":100}`
	if e != fmt.Sprintf("%v", v) {
		t.Errorf("UserStat string is invalid: %v", v)
	}
}