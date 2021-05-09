package config

import (
    "flag"
    "sync"
    "io/ioutil"
    "errors"

    yaml "gopkg.in/yaml.v2"
)

type Config struct {
    API API `yaml:"api"`
    Logs Logs `yaml:"logs"`
    Db     Db  `yaml: "mysql"`
}

type API struct {
    Debug bool `yaml:"debug"`
    Port uint `yaml:"port"`
}

type Logs struct {
    Level string `yaml:"level"`
    LogPath string `yaml:"log_path"`
}

// TODO: improve; support pg etc.
type Db struct {
    Host string `yaml:"host"`
    Port string `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
}


var (
    config = new(Config)
    cLock = new(sync.RWMutex)
)

func Cfg() *Config {
    cLock.RLock()
    defer cLock.RUnlock()
    return config
}

func  ParseConfig(filepath string) error {
    if len(filepath) == 0 {
        return errors.New("缺少配置文件")
    }

    b, err := ioutil.ReadFile(filepath)
    if err != nil {
        return err
    }

    err = yaml.Unmarshal(b, config)
    return err
}

func init() {
    configFile := flag.String("c", "config.yml", "配置文件")
    flag.Parse()
    if err := ParseConfig(*configFile); err != nil {
        panic(err)
    }
}