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
    MapVolumes []MapVolume `json:"mapped_volumes"`
}

type Port struct {
    Guest int `json:"guest"`
    Expose int `json:"expose"`
}

type MapVolume struct {
    Guest string `json:"guest"`
    Host string `json:"host"`
}

func (q *ServerSpec) Parse(file string) error {
    J, err := ioutil.ReadFile(file)
    if err != nil { panic(err) }

    var data = &q
    return json.Unmarshal(J, data)
}