package configurator

import (
	"io/ioutil"
	"reflect"

	"launchpad.net/goyaml"
)

type ConfigHolder struct {
	template reflect.Value
	configs  map[string]reflect.Value

	ConfigFilePath string
}

func NewConfig(cfg interface{}) *ConfigHolder {
	holder := ConfigHolder{configs: make(map[string]reflect.Value)}
	holder.template = reflect.ValueOf(cfg)
	if holder.template.Kind() != reflect.Struct {
		panic("NewConfig requires struct type passed to it.")
	}
	return &holder
}

func (c *ConfigHolder) LoadFile(id string, filePath string) error {
	_, notEmpty := c.configs[id]
	if notEmpty {
		panic("config already declared " + id)
	}
	cfg := reflect.New(c.template.Type())

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = goyaml.Unmarshal([]byte(contents), cfg.Interface())
	if err != nil {
		return err
	}

	c.configs[id] = cfg
	return nil
}

func (c *ConfigHolder) Add(id string, cfg interface{}) {
	_, notEmpty := c.configs[id]
	if notEmpty {
		panic("config already declared " + id)
	}
	c.configs[id] = reflect.ValueOf(cfg)
}

func (c *ConfigHolder) Merge(ids []string) interface{} {
	retCfg := reflect.New(c.template.Type())
	for _, id := range ids {
		cfg, ok := c.configs[id]
		if ok == true {
			meldValueStructs(cfg, retCfg)
		}
	}
	return retCfg.Interface()
}
