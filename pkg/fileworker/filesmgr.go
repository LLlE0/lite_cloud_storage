package fileworker

import (
	"os"

	"github.com/spf13/viper"
)

//here will be the main logic of working with files, loooot of stuff to do

func CreateDir(s string) {
	os.MkdirAll(viper.GetString("store")+"/"+s, os.FileMode(0777))

}
