package shell

import (
	"fmt"

	"github.com/malba124/godan/help"
	"github.com/malba124/godan/api"
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
	fmt.Printf("\nGodan shell help page\n")
}

func handle(cmd string) {

	if cmd == "help" || cmd == "?" || cmd == "HELP" {
		shellHelp()
	}
}

func Shell(key string, color *bool) error {

	help.Banner()

	fmt.Printf("\nGodan Shell, <help> or <?> for help\n")

	userName, err := getNameOfUser(key)
	var env string
	if err != nil {
	    env = "godan->"
	} else {
	    env = fmt.Sprintf("%s@godan->", userName)
	}

	//var currKey string = key
	for {
		cmd := cmd(env)
		handle(cmd)
	}		
}