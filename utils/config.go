package utils

import (
	"flag"

	"github.com/stathat/jconfig"
)

var Conf *jconfig.Config
var configFile = flag.String("conf", "../../../test.conf", "Pass in a config file")

func init() {
	flag.Parse()
	Conf = jconfig.LoadConfig(*configFile)
}
