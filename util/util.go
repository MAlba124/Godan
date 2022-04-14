package util

import (
    "errors"
    "os"
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    Key string
}

const Version string = "BETA"

func LoadConfig(configFile string) (string, error) {

    if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
        return "", errors.New("File does not exist")
    }

    content, err := ioutil.ReadFile(configFile)
    if err != nil {
        return "", errors.New("Failed to open config file")
    }
 
    var config Config
    err = json.Unmarshal(content, &config)
    if err != nil {
        return "", errors.New("Error during Unmarshal()")
    }

    return config.Key, nil
}

func NotAnArgument(arg string, args []string) bool {

    for i := 0; i < len(args); i++ {
        if(arg == args[i]) {
            return false
        }
    }

    return true
}
