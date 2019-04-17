package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PingThingsIO/c37wavemq"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: c37wavemq config.yaml\n")
		os.Exit(1)
	}
	conf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("could not read config file: %v\n", err)
		os.Exit(2)
	}
	params := &c37wavemq.ProtocolAdapterConfig{}
	err = yaml.Unmarshal(conf, &params)
	if err != nil {
		fmt.Printf("could not parse config: %v\n", err)
		os.Exit(3)
	}
	c37wavemq.StartProtocolAdapter(params)
	for {
		fmt.Printf("start ended\n")
	}
}
