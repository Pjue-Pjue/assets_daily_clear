package config

import (
	"github.com/go-yaml/yaml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type Configuration struct {
	DbDirect   string `yaml:"db_direct"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
}

var configuration *Configuration

func init() {
	binDir, err := ExecutableDir()
	if err != nil {
		panic(err)
	}
	ROOT := path.Dir(binDir) + "/"
	//log.Println(ROOT)

	filename := "./conf.yaml"
	configFile := filepath.Join(ROOT, filename)
	if _, err := os.Stat(configFile); err != nil {
		configFile = filepath.Join(ROOT, "..", "exe", filename)
		//log.Println(configFile)
	}

	// Load配置文件
	loadConfiguration(configFile)
}

// 获得可执行程序所在目录
func ExecutableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(pathAbs), nil
}

func loadConfiguration(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Fatalf("error: %v", err)
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logrus.Fatalf("error: %v", err)
	}
	configuration = &config
	logrus.Debugf("config load succeessfully:%v", configuration)
}

func GetConfiguration() *Configuration {
	return configuration
}
