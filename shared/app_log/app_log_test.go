package app_log_test

import (
	"bytes"
	. "github.com/marugoshi/gobm/shared/app_log"
	"regexp"
	"testing"
)

func TestDebugf(t *testing.T) {
	buf := buf()

	Debugf("this is debug message: %s.", "Debugf")
	if m, err := regexp.Match(`\[DEBUG\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[DEBUG] not found.")
	}
	if m, err := regexp.Match(`Debugf`, buf.Bytes()); err != nil || !m {
		t.Errorf("Debugf not found.")
	}
}

func TestDebug(t *testing.T) {
	buf := buf()

	Debugf("this is debug message.")
	if m, err := regexp.Match(`\[DEBUG\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[DEBUG] not found.")
	}
}

func TestInfof(t *testing.T) {
	buf := buf()
	Infof("this is info message: %s.", "Infof")
	if m, err := regexp.Match(`\[INFO\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[INFO] not found.")
	}
	if m, err := regexp.Match(`Infof`, buf.Bytes()); err != nil || !m {
		t.Errorf("Infof not found.")
	}
}

func TestInfo(t *testing.T) {
	buf := buf()
	Info("this is info message.")
	if m, err := regexp.Match(`\[INFO\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[INFO] not found.")
	}
}

func TestWarnf(t *testing.T) {
	buf := buf()
	Warnf("this is warn message: %s.", "Warnf")
	if m, err := regexp.Match(`\[WARN\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[WARN] not found.")
	}
	if m, err := regexp.Match(`Warnf`, buf.Bytes()); err != nil || !m {
		t.Errorf("Warnf not found.")
	}
}

func TestWarn(t *testing.T) {
	buf := buf()
	Warn("this is warn message.")
	if m, err := regexp.Match(`\[WARN\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[WARN] not found.")
	}
}

func TestFatalf(t *testing.T) {
	buf := buf()
	Fatalf("this is fatal message: %s.", "Fatalf")
	if m, err := regexp.Match(`\[FATAL\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[FATAL] not found.")
	}
	if m, err := regexp.Match(`Fatalf`, buf.Bytes()); err != nil || !m {
		t.Errorf("Fatalf not found.")
	}
}

func TestFatal(t *testing.T) {
	buf := buf()
	Fatal("this is Fatal message.")
	if m, err := regexp.Match(`\[FATAL\] `, buf.Bytes()); err != nil || !m {
		t.Errorf("[FATAL] not found.")
	}
}

func buf() *bytes.Buffer {
	buf := bytes.NewBuffer(make([]byte, 0))
	SetOutput(buf)
	return buf
}