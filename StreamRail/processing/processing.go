package processing

import (
	"ConfluentADI/model/opencage"
	"ConfluentADI/model/tiploc"
	"encoding/json"
	"fmt"
	"github.com/lovoo/goka"
)

func OpenCage(ctx goka.Context, msg interface{}) {
	datas := msg.(string)
	_ = datas

	var data opencage.Opencage
	var results opencage.ResultOpenCage

	err := json.Unmarshal([]byte(datas), &data)
	if err == nil {
		if len(data.Results) > 0 {
			results.TpsDescription = data.Request.Query
			results.TotalResults = data.TotalResults

			results.Geohash = data.Results[0].Annotations.Geohash
			results.OsmUrl = data.Results[0].Annotations.OSM.URL
			results.Geolatlan = fmt.Sprintf("%f", data.Results[0].Geometry.Lat) +
				"," + fmt.Sprintf("%f", data.Results[0].Geometry.Lng)
			results.Components = data.Results[0].Components
		}
	}
	ctx.SetValue(&results)
}

func TiplocV1(ctx goka.Context, msg interface{}) {
	datas := msg.(string)
	_ = datas

	var data tiploc.RequestTiploc

	err := json.Unmarshal([]byte(datas), &data)
	if err == nil && data.TiplocV1.TransactionType != "" {
		ctx.SetValue(&data)
	}

}
