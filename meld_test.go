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

type subSubConfig struct {
	Str    string
	StrArr []string
}

type subConfig struct {
	SubSubCfg subSubConfig
}

type sampleConfig struct {
	SubCfg subConfig
	StrArr []string
	Dura   time.Duration
	Str    string
}

func makeSampleConfig() sampleConfig {
	return sampleConfig{
		SubCfg: subConfig{
			SubSubCfg: subSubConfig{
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
	testCfg := sampleConfig{}
	meldStructs(makeSampleConfig(), &testCfg)

	if !reflect.DeepEqual(testCfg, makeSampleConfig()) {
		t.Fatalf("%+v doesn't equal %+v", testCfg, makeSampleConfig())
	}
}

func TestMeldTwoStructs(t *testing.T) {
	testCfg := makeSampleConfig()
	overCfg := sampleConfig{
		SubCfg: subConfig{
			SubSubCfg: subSubConfig{
				Str:    "second",
				StrArr: []string{"three", "two"},
			},
		},
		StrArr: []string{"3"},
		Dura:   2,
	}
	meldStructs(overCfg, &testCfg)

	targetCfg := sampleConfig{
		SubCfg: subConfig{
			SubSubCfg: subSubConfig{
				Str:    "second",
				StrArr: []string{"three", "two"},
			},
		},
		StrArr: []string{"3"},
		Dura:   2,
		Str:    "somestr",
	}

	if !reflect.DeepEqual(testCfg, targetCfg) {
		t.Fatalf("%+v doesn't equal %+v", testCfg, makeSampleConfig())
	}
}
