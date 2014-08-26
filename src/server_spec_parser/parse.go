package server_spec_parser

import (
    "io/ioutil"
    "encoding/json"
)

type ServerSpec struct {
    Containers []Container `json:"containers"`
}

type Container struct {
    Name string `json:"name"`
    Image string `json:"image"`
    BuildRequired bool `json:"build_required"`
    Ports []Port `json:"ports"`
    Links []string `json:"links"`
    MappedVolumes []MappedVolume `json:"mapped_volumes"`
    EnvironmentVariables []EnvironmentVariable `json:"environment"`
    Config []json.RawMessage `json:"config"`
}

func (c *Container) GetConfigFormat(config json.RawMessage) string {
    var objmap map[string]*json.RawMessage
    json.Unmarshal(config, &objmap)

    var str string
    json.Unmarshal(*objmap["format"], &str)

    return str;
}

type Port struct {
    Guest int `json:"guest"`
    Expose int `json:"expose"`
}

type MappedVolume struct {
    Guest string `json:"guest"`
    Host string `json:"host"`
}

type EnvironmentVariable struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

func (q *ServerSpec) Parse(file string) error {
    J, err := ioutil.ReadFile(file)
    if err != nil { panic(err) }

    var data = &q
    return json.Unmarshal(J, data)
}