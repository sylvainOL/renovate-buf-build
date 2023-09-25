package main

import (
	"fmt"
	"time"

	prompb "buf.build/gen/go/prometheus/prometheus/protocolbuffers/go"
)

func main() {
	timeSeries := []*prompb.TimeSeries{} timeSeries.append(
	timeSeries, &prompb.TimeSeries{
			Labels: []*prompb.Label{
					{
						Name:  "__name__",
						Value: "my_ts",
					},
				}, Samples: []*prompb.Sample{
					{
						Value:     float64(164),
						Timestamp:
						time.Now().UnixMilli(),
					},
				},
			}
	)
}

