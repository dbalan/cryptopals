// these tests depend on server insecure_cmp running!
package set4

import (
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMakeRequest(t *testing.T) {
	_, ts := makeRequest("foo", startingPoint)
	assert.NotEqual(t, 0, ts)
}

func TestMakeRequestTimes(t *testing.T) {
	var series []time.Duration
	for i := 0; i < 1000000000000; i++ {
		_, ts := makeRequest("foo", startingPoint)
		series = append(series, ts)
	}
	common.StandardDeviation(series)
	//	assert.NotEqual(t, 0, ts)
}

func TestBruteForceSig(t *testing.T) {
	//	bruteforceSig("findMySig")
}
