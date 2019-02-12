package set5

import (
	"testing"
)

func TestTestNormalComm(t *testing.T) {
	normalRecv := newB()
	communicate(normalRecv)
}

func TestMITM(t *testing.T) {
	mal := newM()
	communicate(mal)
}
