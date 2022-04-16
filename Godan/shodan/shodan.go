package shodan

import (
    "fmt"
    "errors"

    "github.com/malba124/godan/api"
    "github.com/malba124/godan/printer"
    "github.com/malba124/godan/color"
)

type Client struct {
    Key string
    Timeout int
}

func (c *Client)ApiInfo(colored *bool) error {

    if *colored {
        printer.Print(fmt.Sprintf("%s[ ... ] Sending request%s", color.CGray, color.CReset))
    } else {
        printer.Print("[ ... ] Sending request")
    }

    statusCode, info, err := api.ApiInfo(c.Key, c.Timeout)

    if err != nil {
        if statusCode == 404 {
            printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", 
                printer.WHITESPACE, color.CRed, statusCode, color.CReset))
            printer.Error("Invalid host", colored)
            return errors.New("404")
        } else if statusCode == 401 {
            printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", 
                printer.WHITESPACE, color.CRed, statusCode, color.CReset))
            printer.Error("Please verify that you API key is valid!", colored)
            return errors.New("401")
        } else if statusCode == 429 {
            printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", 
                printer.WHITESPACE, color.CRed, statusCode, color.CReset))
            return errors.New("429")
        }
    } else {
        printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", 
            printer.WHITESPACE, color.CGreen, statusCode, color.CReset))

        printer.Print(fmt.Sprintf("\nAPI Information about Key: %s...\n", c.Key[:5]))
        
        if info.QueryCredits < 1 {
            printer.Print(fmt.Sprintf(" ├─ Query Credits: %s%d%s\n", 
                color.CRed, info.QueryCredits, color.CReset))    
        } else {
            printer.Print(fmt.Sprintf(" ├─ Query Credits: %s%d%s\n", 
                color.CGreen, info.QueryCredits, color.CReset))
        }
        if info.ScanCredits < 1 {
            printer.Print(fmt.Sprintf(" ├─ Scan Credits: %s%d%s\n", 
                color.CRed, info.ScanCredits, color.CReset))    
        } else {
            printer.Print(fmt.Sprintf(" ├─ Scan Credits: %s%d%s\n", 
                color.CGreen, info.ScanCredits, color.CReset))
        }
        if len(info.Plan) < 1 {
            printer.Print(" ├─ Plan: ???\n")    
        } else {
            printer.Print(fmt.Sprintf(" ├─ Plan: %s\n", info.Plan))
        }
        printer.Print(fmt.Sprintf(" └─ Monitored IPs: %d\n", info.MonitoredIps))

        return nil
    }
    return nil
}

func (c *Client)Host(ip string, colored *bool) error {

    if *colored {
        printer.Print(fmt.Sprintf("%s[ ... ] Sending request%s", color.CGray, color.CReset))
    } else {
        printer.Print("[ ... ] Sending request")
    }

    statusCode, host, err := api.Host(ip, c.Key, c.Timeout)
    if err != nil {
        if statusCode == 404 {
            printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", 
                printer.WHITESPACE, color.CRed, statusCode, color.CReset))
            printer.Error("Invalid host", colored)
            return errors.New("404")
        } else if statusCode == 401 {
            printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", 
                printer.WHITESPACE, color.CRed, statusCode, color.CReset))
            printer.Error("Please verify that you API key is valid!", colored)
            return errors.New("401")
        } else if statusCode == 429 {
            printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", 
                printer.WHITESPACE, color.CRed, statusCode, color.CReset))
            return errors.New("429")
        }
    } else {
        printer.Print(fmt.Sprintf("\r%s\r[ %s%d%s ] Request sent\n", printer.WHITESPACE, color.CGreen, statusCode, color.CReset))

        printer.Print(fmt.Sprintf("\nHost information for %s\n", ip))
        printer.Print(fmt.Sprintf("Shodan link: https://www.shodan.io/host/%s\n",
            ip))
        if len(host.CountryName) > 0 {
            printer.Print(fmt.Sprintf(" ├─ Country: %s\n", host.CountryName))
            if len(host.CountryCode) > 0 {
                printer.Print(fmt.Sprintf("   └─ %s\n", host.CountryCode))
            }
        } else if len(host.CountryCode) > 0 {
            printer.Print(fmt.Sprintf(" ├─ Country: %s\n", host.CountryCode))
        }else {
            printer.Print(" ├─ Country: N/A\n")
        }
        if len(host.City) > 0 {
            printer.Print(fmt.Sprintf(" ├─ City: %s\n", host.City))
        } else {
            printer.Print(" ├─ City: N/A\n")
        }
        if len(host.LastUpdate) > 0 {
            printer.Print(fmt.Sprintf(" ├─ Last Update: %s\n", host.LastUpdate))
        }
        if len(host.Isp) > 0 {
            printer.Print(fmt.Sprintf(" ├─ ISP: %s\n", host.Isp))
        }
        if len(host.Org) > 0 {
            printer.Print(fmt.Sprintf(" ├─ Org: %s\n", host.Org))
        } 
        if len(host.Os) > 0 {
            printer.Print(fmt.Sprintf("  ├─ Os: %s\n", host.Os))
        }
        if len(host.Ports) > 0 {
            printer.Print(" ├─ Ports\n")
            for i := 0; i < len(host.Ports); i++ {
                if i >= len(host.Ports)-1 {
                    printer.Print(fmt.Sprintf(" │  └─ %d\n", host.Ports[i]))
                } else {
                    printer.Print(fmt.Sprintf(" │  ├─ %d\n", host.Ports[i]))
                }
            }
        }
        if len(host.Domains) > 0 {
            printer.Print(" └─ Domains\n")
            for i := 0; i < len(host.Domains); i++ {
                if i >= len(host.Domains)-1 {
                    printer.Print(fmt.Sprintf("    └─ %s\n", host.Domains[i]))
                } else {
                    printer.Print(fmt.Sprintf("    ├─ %s\n", host.Domains[i]))
                }
            }
        }
        return nil
    }
    return nil
}
