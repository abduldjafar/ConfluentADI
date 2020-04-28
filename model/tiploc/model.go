package tiploc

import (
	"encoding/json"
	"fmt"
)

type RequestTiploc struct {
	TiplocV1 struct {
		TransactionType string      `json:"transaction_type"`
		TiplocCode      string      `json:"tiploc_code"`
		Nalco           string      `json:"nalco"`
		Stanox          string      `json:"stanox"`
		CrsCode         interface{} `json:"crs_code"`
		Description     interface{} `json:"description"`
		TpsDescription  string      `json:"tps_description"`
	} `json:"TiplocV1"`
}

func (r RequestTiploc) Encode(value interface{}) (data []byte, err error) {
	if _, isUser := value.(*RequestTiploc); !isUser {
		return nil, fmt.Errorf("Codec requires value *user, got %T", value)
	}
	return json.Marshal(value)
}

func (r RequestTiploc) Decode(data []byte) (value interface{}, err error) {
	var (
		c RequestTiploc
		_ error
	)
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling user: %v", err)
	}
	return &c, nil
}
