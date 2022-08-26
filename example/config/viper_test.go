package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func Test1(t *testing.T) {

	//os.Getenv("")

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if nil != err {
		t.Fatal(err)
	}
	env := viper.GetString("env")
	t.Log(env)
	viper.SetConfigName(fmt.Sprintf("config-%s.yaml", env))
	viper.ReadInConfig()
	t.Log(viper.GetString("config"))
	t.Log(viper.GetString("user1.name"))
	t.Log(viper.GetInt64("user1.pwd"))
	t.Log(viper.GetString("user2.name"))
	t.Log(viper.GetInt64("user2.pwd"))
}
