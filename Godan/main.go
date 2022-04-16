package main

import (
    "flag"
    "os"
    "errors"

    "github.com/malba124/godan/util"
    "github.com/malba124/godan/printer"
    "github.com/malba124/godan/shodan"
    "github.com/malba124/godan/help"
    "github.com/malba124/godan/api"
    "github.com/malba124/godan/shell"
)

type Host struct {
    ip *string
    color *bool
    verbose *bool
}

type ApiInfo struct {
    verbose *bool
    color *bool
}

func main() {

    if len(os.Args) < 2 {
        help.MainHelp()
        os.Exit(1)
    }

    if util.NotAnArgument("-h", os.Args) && util.NotAnArgument("--help", os.Args) {
        key, err := util.LoadConfig("config.json")

        var client shodan.Client

        hostCmd := flag.NewFlagSet("host", flag.ExitOnError)
        hIp := hostCmd.String("ip", "", "Host ip")
        hVerbose := hostCmd.Bool("v", false, "Verbose mode")
        hColor := hostCmd.Bool("c", true, "Disable colored output")
        hTimeout := hostCmd.Int("t", 10, "Specify request timeout")
        hKey := hostCmd.String("x", "", "Use specific api key")

        apiInfoCmd := flag.NewFlagSet("api-info", flag.ExitOnError)
        aiVerbose := apiInfoCmd.Bool("v", false, "Verbose mode")
        aiColor := apiInfoCmd.Bool("c", true, "Disable colored output")
        aiTimeout := apiInfoCmd.Int("t", 10, "Specify request timeout")
        aiKey := apiInfoCmd.String("x", "", "Use specific api key")

        shellCmd := flag.NewFlagSet("shell", flag.ExitOnError)    
        shellColor := shellCmd.Bool("c", true, "Disable colored output")
        shellKey := shellCmd.String("x", "", "Use specific api key")

        switch os.Args[1] {
        case "host":
            hostCmd.Parse(os.Args[2:])
            client.Timeout = *hTimeout
            var h Host = Host{
                    ip: hIp,
                    color: hColor,
                    verbose: hVerbose,
                }
            if len(*hKey) > 0 {
                client.Key = *hKey
            } else if err == nil && len(key) > 0 {
                client.Key = key 
            } else {
                printer.Error("Config file is bad or missing api key!", hColor)
                os.Exit(1)
            }
            printer.Verbose("Executing Host()", h.verbose)
            err := client.Host(*h.ip, h.color)                    
            if (errors.Is(err, api.Err401) || err == nil) && len(*hKey) > 0 {
                util.SaveKey(*aiKey, h.color)
            } 
        case "api-info":
            apiInfoCmd.Parse(os.Args[2:])
            client.Timeout = *aiTimeout
            var ai ApiInfo = ApiInfo {
                verbose: aiVerbose,
                color: aiColor,
            }
            if len(*aiKey) > 0 {
                client.Key = *aiKey
            } else if err == nil && len(key) > 0 {
                client.Key = key 
            } else {
                printer.Error("Config file is bad or missing api key!", aiColor)
                os.Exit(1)
            }
            printer.Verbose("Executing ApiInfo()", ai.verbose)
            err := client.ApiInfo(ai.color)
            if (!errors.Is(err, api.Err401) || err == nil) && len(*aiKey) > 0 {
                util.SaveKey(*aiKey, ai.color)
            } 
        case "shell":
            shellCmd.Parse(os.Args[2:])
            var sKey string
            if len(*shellKey) > 0 {
                sKey = *shellKey
            } else if err == nil && len(key) > 0 {
                sKey = key 
            } else {
                printer.Error("Config file is bad or missing api key!", hColor)
                os.Exit(1)
            }
            shell.Shell(sKey, shellColor)
        default:
            help.MainHelp()
            printer.Error("Missing arguments", nil)
            os.Exit(1)
        }
    } else {
        help.MainHelp()
    }
}
