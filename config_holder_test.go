package configurator

import (
	"reflect"
	"testing"
	"time"
)

type chSampleConfig struct {
	Str    string
	StrArr []string
	Dura   time.Duration
}

type chSampleArr []chSampleConfig

func TestDeserializationAndMerging(t *testing.T) {
	cfg := NewConfig(chSampleConfig{})
	cfg.LoadFile("main", "test_data/ch_sample_config.yml")
	cfg.LoadFile("specific", "test_data/ch_sample_config2.yml")
	x := cfg.Merge([]string{"main", "specific"}).(*chSampleConfig)

	expected := &chSampleConfig{Str: "first_string", StrArr: []string{"two", "three"}, Dura: 2}
	if !reflect.DeepEqual(expected, x) {
		t.Fatalf("%+v != %+v", x, expected)
	}
}
