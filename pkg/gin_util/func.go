package gin_util

import (
	"github.com/spf13/viper"
)

type runMode = string

const (
	modeDev   runMode = "dev"
	modeTest  runMode = "test"
	modeDebug runMode = "debug"
	modeProd  runMode = "prod"
)

func IsDev() bool {
	return getRunMode() == modeDev
}

func getRunMode() string {
	return viper.GetString("runmode")
}
func IsTest() bool {
	return getRunMode() == modeTest
}
func IsDebug() bool {
	return getRunMode() == modeDebug
}
func IsProd() bool {
	return getRunMode() == modeProd
}
