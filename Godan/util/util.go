package util

import (
    "fmt"
    "errors"
    "os"
    "encoding/json"
    "io/ioutil"

    "github.com/malba124/godan/printer"
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

func SaveKey(key string, color *bool) {

    var answ string
    printer.Print("\nDo you want to save provided API key to config file? [y/n]> ")
    fmt.Scanln(&answ)

    if answ != "y" {
        return
    }

    file, err := os.Create("config.json")    
    if err != nil {
        printer.Error("Failed to create config file", color)
    }

    defer file.Close()

    _, err = file.WriteString(fmt.Sprintf("{\n\t\"Key\": \"%s\"\n}", key))
    if err != nil {
        printer.Error("Failed to write to config file", color)
    }
}