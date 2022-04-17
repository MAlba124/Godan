package shell

import (
	"fmt"
	"strings"
	"os"

	"github.com/malba124/godan/help"
	"github.com/malba124/godan/printer"
	"github.com/malba124/godan/api"
	"github.com/malba124/godan/color"
)

func getNameOfUser(key string) (string, error) {

	statusCode, profile, err := api.Profile(key, 10)

    if err != nil {
        if statusCode == 404 {
            return "", api.Err404
        } else if statusCode == 401 {
            return "", api.Err401
        } else if statusCode == 429 {
            return "", api.Err429
        }
    } else {
    	if len(profile.DisplayName) > 0 {
    		if strings.Contains(profile.DisplayName, "@") {
    		cleanedName := strings.Split(profile.DisplayName, "@")
    		return cleanedName[0], nil
    		} 
    		return profile.DisplayName, nil	
    	}
    }
    return "", nil
}

func cmd(env string) string {

	var ret string
	fmt.Printf("%s ", env)
	fmt.Scanln(&ret)

	return ret
}

func shellHelp() {
	fmt.Printf("Godan shell help page\n")
}

func handle(cmd string, colored *bool) {

	switch cmd {
	case "help", "?", "HELP":
		shellHelp()
	case "exit":
		os.Exit(0)
	default:
		printer.Error(fmt.Sprintf("Command %s not found", cmd), colored)
	}
}

func Shell(key string, colored *bool) error {

	help.Banner()

	fmt.Printf("\nGodan Shell, <help> or <?> for help\n")

	userName, err := getNameOfUser(key)
	var env string
	if err != nil {
	    env = "godan->"
	} else {
	    env = fmt.Sprintf("%s%s@godan%s->", color.CPurple, userName, color.CReset)
	}

	//var currKey string = key
	for {
		cmd := cmd(env)
		handle(cmd, colored)
	}		
}