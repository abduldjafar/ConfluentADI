package opencage

import (
	"encoding/json"
	"fmt"
)

type Opencage struct {
	Documentation string `json:"documentation"`
	Licenses      []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"licenses"`
	Rate struct {
		Limit     int `json:"limit"`
		Remaining int `json:"remaining"`
		Reset     int `json:"reset"`
	} `json:"rate"`
	Request struct {
		Abbrv         int    `json:"abbrv"`
		AddRequest    int    `json:"add_request"`
		Autocomplete  int    `json:"autocomplete"`
		Countrycode   string `json:"countrycode"`
		Format        string `json:"format"`
		Key           string `json:"key"`
		Language      string `json:"language"`
		Limit         int    `json:"limit"`
		MinConfidence int    `json:"min_confidence"`
		NoAnnotations int    `json:"no_annotations"`
		NoDedupe      int    `json:"no_dedupe"`
		NoRecord      int    `json:"no_record"`
		OnlyNominatim int    `json:"only_nominatim"`
		Pretty        int    `json:"pretty"`
		Proximity     string `json:"proximity"`
		Query         string `json:"query"`
		Roadinfo      int    `json:"roadinfo"`
		Version       string `json:"version"`
	} `json:"request"`
	Results []struct {
		Annotations struct {
			DMS struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"DMS"`
			MGRS       string `json:"MGRS"`
			Maidenhead string `json:"Maidenhead"`
			Mercator   struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"Mercator"`
			OSM struct {
				EditURL string `json:"edit_url"`
				URL     string `json:"url"`
			} `json:"OSM"`
			UNM49 struct {
				Regions struct {
					EUROPE         string `json:"EUROPE"`
					GB             string `json:"GB"`
					NORTHERNEUROPE string `json:"NORTHERN_EUROPE"`
					WORLD          string `json:"WORLD"`
				} `json:"regions"`
				StatisticalGroupings []string `json:"statistical_groupings"`
			} `json:"UN_M49"`
			Callingcode int `json:"callingcode"`
			Currency    struct {
				AlternateSymbols     []interface{} `json:"alternate_symbols"`
				DecimalMark          string        `json:"decimal_mark"`
				HTMLEntity           string        `json:"html_entity"`
				IsoCode              string        `json:"iso_code"`
				IsoNumeric           string        `json:"iso_numeric"`
				Name                 string        `json:"name"`
				SmallestDenomination int           `json:"smallest_denomination"`
				Subunit              string        `json:"subunit"`
				SubunitToUnit        int           `json:"subunit_to_unit"`
				Symbol               string        `json:"symbol"`
				SymbolFirst          int           `json:"symbol_first"`
				ThousandsSeparator   string        `json:"thousands_separator"`
			} `json:"currency"`
			Flag     string  `json:"flag"`
			Geohash  string  `json:"geohash"`
			Qibla    float64 `json:"qibla"`
			Roadinfo struct {
				DriveOn string `json:"drive_on"`
				SpeedIn string `json:"speed_in"`
			} `json:"roadinfo"`
			Sun struct {
				Rise struct {
					Apparent     int `json:"apparent"`
					Astronomical int `json:"astronomical"`
					Civil        int `json:"civil"`
					Nautical     int `json:"nautical"`
				} `json:"rise"`
				Set struct {
					Apparent     int `json:"apparent"`
					Astronomical int `json:"astronomical"`
					Civil        int `json:"civil"`
					Nautical     int `json:"nautical"`
				} `json:"set"`
			} `json:"sun"`
			Timezone struct {
				Name         string `json:"name"`
				NowInDst     int    `json:"now_in_dst"`
				OffsetSec    int    `json:"offset_sec"`
				OffsetString string `json:"offset_string"`
				ShortName    string `json:"short_name"`
			} `json:"timezone"`
			What3Words struct {
				Words string `json:"words"`
			} `json:"what3words"`
			Wikidata string `json:"wikidata"`
		} `json:"annotations,omitempty"`
		Bounds struct {
			Northeast struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"northeast"`
			Southwest struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"southwest"`
		} `json:"bounds"`
		Components struct {
			ISO31661Alpha2 string `json:"ISO_3166-1_alpha-2"`
			ISO31661Alpha3 string `json:"ISO_3166-1_alpha-3"`
			Type           string `json:"_type"`
			Continent      string `json:"continent"`
			Country        string `json:"country"`
			CountryCode    string `json:"country_code"`
			County         string `json:"county"`
			CountyCode     string `json:"county_code"`
			PoliticalUnion string `json:"political_union"`
			Postcode       string `json:"postcode"`
			State          string `json:"state"`
			StateCode      string `json:"state_code"`
			StateDistrict  string `json:"state_district"`
			Town           string `json:"town"`
		} `json:"components,omitempty"`
		Confidence int    `json:"confidence"`
		Formatted  string `json:"formatted"`
		Geometry   struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	StayInformed struct {
		Blog    string `json:"blog"`
		Twitter string `json:"twitter"`
	} `json:"stay_informed"`
	Thanks    string `json:"thanks"`
	Timestamp struct {
		CreatedHTTP string `json:"created_http"`
		CreatedUnix int    `json:"created_unix"`
	} `json:"timestamp"`
	TotalResults int `json:"total_results"`
}

type ResultOpenCage struct {
	TpsDescription string `json:"TPS_DESCRIPTION"`
	TotalResults   int    `json:"TOTAL_RESULTS"`
	Geohash        string `json:"GEOHASH"`
	OsmUrl         string `json:"OSM_URL"`
	Components     struct {
		ISO31661Alpha2 string `json:"ISO_3166-1_alpha-2"`
		ISO31661Alpha3 string `json:"ISO_3166-1_alpha-3"`
		Type           string `json:"_type"`
		Continent      string `json:"continent"`
		Country        string `json:"country"`
		CountryCode    string `json:"country_code"`
		County         string `json:"county"`
		CountyCode     string `json:"county_code"`
		PoliticalUnion string `json:"political_union"`
		Postcode       string `json:"postcode"`
		State          string `json:"state"`
		StateCode      string `json:"state_code"`
		StateDistrict  string `json:"state_district"`
		Town           string `json:"town"`
	}
	Geolatlan string `json:"GEO_LATLON"`
}

func (r ResultOpenCage) Encode(value interface{}) (data []byte, err error) {
	if _, isUser := value.(*ResultOpenCage); !isUser {
		return nil, fmt.Errorf("Codec requires value *user, got %T", value)
	}
	return json.Marshal(value)
}

func (r ResultOpenCage) Decode(data []byte) (value interface{}, err error) {
	var (
		c ResultOpenCage
		_ error
	)
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling user: %v", err)
	}
	return &c, nil
}
