package configurator

import (
	// "fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

func init() {
	log.SetFlags(0)
}

type SubSubConfig struct {
	Str    string
	StrArr []string
}

type SubConfig struct {
	SubSubCfg SubSubConfig
}

type SampleConfig struct {
	SubCfg SubConfig
	StrArr []string
	Dura   time.Duration
	Str    string
}

func sampleConfig() SampleConfig {
	return SampleConfig{
		SubCfg: SubConfig{
			SubSubCfg: SubSubConfig{
				Str:    "first",
				StrArr: []string{"one", "two"},
			},
		},
		StrArr: []string{"1", "2"},
		Dura:   1,
		Str:    "somestr",
	}
}

func TestMeldToEmptyStruct(t *testing.T) {
	testCfg := SampleConfig{}
	meldStructs(sampleConfig(), &testCfg)

	if !reflect.DeepEqual(testCfg, sampleConfig()) {
		t.Fatalf("%+v doesn't equal %+v", testCfg, sampleConfig())
	}
}

func TestMeldTwoStructs(t *testing.T) {
	testCfg := sampleConfig()
	overCfg := SampleConfig{
		SubCfg: SubConfig{
			SubSubCfg: SubSubConfig{
				Str:    "second",
				StrArr: []string{"three", "two"},
			},
		},
		StrArr: []string{"3"},
		Dura:   2,
	}
	meldStructs(overCfg, &testCfg)

	targetCfg := SampleConfig{
		SubCfg: SubConfig{
			SubSubCfg: SubSubConfig{
				Str:    "second",
				StrArr: []string{"three", "two"},
			},
		},
		StrArr: []string{"3"},
		Dura:   2,
		Str:    "somestr",
	}

	if !reflect.DeepEqual(testCfg, targetCfg) {
		t.Fatalf("%+v doesn't equal %+v", testCfg, sampleConfig())
	}
}
