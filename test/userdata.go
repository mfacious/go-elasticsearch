// Code generated by ent-elastic, DO NOT EDIT.

package test

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"time"
)

type UserData struct {
	ID string `json:"id,omitempty"`

	KeywordKkkk string `json:"keyword_kkkk,omitempty"`

	ByteBbbb int8 `json:"byte_bbbb,omitempty"`

	ShortSsss int16 `json:"short_ssss,omitempty"`

	IntIiii int32 `json:"int_iiii,omitempty"`

	LongLlll int64 `json:"long_llll,omitempty"`

	FloatFfff float32 `json:"float_ffff,omitempty"`

	DoubleDdddd float64 `json:"double_ddddd,omitempty"`

	BoolBbbb bool `json:"bool_bbbb,omitempty"`

	DateDddd time.Time `json:"date_dddd,omitempty"`

	StingsSssss []string `json:"stings_sssss,omitempty"`

	IntsIiiiii []int `json:"ints_iiiiii,omitempty"`

	Int64sIiiiii []int64 `json:"int64s_iiiiii,omitempty"`

	FloatsLlllll []float32 `json:"floats_llllll,omitempty"`
}

func UnserizerUserData(result *elastic.SearchResult) ([]*UserData, error) {
	datas := make([]*UserData, 0)
	for _, h := range result.Hits.Hits {
		data := &UserData{}
		err := json.Unmarshal(h.Source, data)
		if err != nil {
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}
