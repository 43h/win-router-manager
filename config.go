package main

import (
	"fmt"
)

type Config struct {
	Local  Connection `yaml:"local"`
	Remote Connection `yaml:"remote"`
	Listen Connection `yaml:"listen"`
}

type Connection struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

var config Config
var cfgFile = "./conf.yaml"

func loadConfig() error {
	//if configFile != "" {
	//	cfgFile = configFile
	//}
	//data, err := ioutil.ReadFile(cfgFile)
	//if err != nil {
	//	return fmt.Errorf("error reading config file: %v", err)
	//}
	//
	//err = yaml.Unmarshal(data, &config)
	//if err != nil {
	//	return fmt.Errorf("error unmarshalling config file: %v", err)
	//}

	return nil
}

func dumpConfig() {
	fmt.Printf("Local IP: %s\n", config.Local.IP)
	fmt.Printf("Local Port: %d\n", config.Local.Port)
	fmt.Printf("Remote IP: %s\n", config.Remote.IP)
	fmt.Printf("Remote Port: %d\n", config.Remote.Port)
	fmt.Printf("Listen IP: %s\n", config.Listen.IP)
	fmt.Printf("Lisen Port: %d\n", config.Listen.Port)
}
