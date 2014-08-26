package config_creators

import (
)

type ConfigSection struct {
	Format string `json:"format"`
	Target string `json:"target"`
	Domain string `json:"domain"`
	Ports []Port `json:"ports"`
}

type Port struct {
    Guest int `json:"guest"`
    Expose int `json:"expose"`
}

func Create(config_json []interface{}) []string {
	var cmds []string

    for _, c := range config_json {
	    var config &ConfigSection
		json.Unmarshal(*c, &config)
	}

	return cmds
}