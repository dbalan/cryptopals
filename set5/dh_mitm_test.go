package set5

import (
	"testing"
)

func TestTestNormalComm(t *testing.T) {
	normalRecv := newB()
	communicate(normalRecv)
}

func TestKey0MITM(t *testing.T) {
	mal := newK0Mitm()
	communicate(mal)
}
