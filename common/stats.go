package common

import (
	"fmt"
	"github.com/gonum/stat"
	"time"
)

func StandardDeviation(ts []time.Duration) {
	datalen := len(ts)
	series := []float64{}
	for i := 0; i < datalen; i++ {
		series = append(series, ts[i].Seconds())
	}
	sd := stat.StdDev(series, nil)
	fmt.Printf("sd: %.8f\n", sd)
}
