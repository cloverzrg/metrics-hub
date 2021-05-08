package config

import "flag"

type cmdParams struct {
	ConsulAddress  string
}

func ParseCmdParams() (params *cmdParams) {
	params = &cmdParams{}
	flag.StringVar(&params.ConsulAddress, "consul", "127.0.0.1:8500", "")
	flag.Parse()
	return params
}
