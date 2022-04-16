package api

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
    "errors"

    "github.com/malba124/godan/util"
)

func NewHttpClient(timeout int) (*http.Client) {

    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
    }

    return client
}

var Err404 = errors.New("404")
var Err401 = errors.New("401")
var Err429 = errors.New("429")

type Client struct {
    Key string
}

type ProfileS struct {
    Member bool             `json:"member"`
    Credits int             `json:"credits"`
    DisplayName string      `json:"display_name"`
    Created string          `json:"created"`
}

type ApiInfoS struct {
    QueryCredits int        `json:"query_credits"`
    ScanCredits int         `json:"scan_credits"`
    Plan string             `json:"plan"`
    MonitoredIps int        `json:"monitored_ips"`
}

type HostS struct {
    RegionCode string       `json:"region_code"`
    PostalCode string       `json:"postal_code"`
    CountryCode string      `json:"country_code"`
    City string             `json:"city"`
    LastUpdate string       `json:"last_update"`
    Latitude float64        `json:"latitude"`
    Tags []string           `json:"tags"`
    AreaCode string         `json"area_code"`
    CountryName string      `json"country_name"`
    Hostnames []string      `json"hostnames"`
    Org string              `json:"org"`
    Asn string              `json."asn"`
    Isp string              `json:"isp"`
    Longitude float64       `json:"longitude"`
    Domains []string        `json:"domains"`
    Os string               `json:"os"`
    Ports []int             `json:"ports"`
}

func ApiInfo(apiKey string, timeout int) (int, *ApiInfoS, error) {

    client := NewHttpClient(timeout)

    var url string
    url = fmt.Sprintf("https://api.shodan.io/api-info?key=%s", apiKey)

    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, nil, err
    }

    request.Header.Set("User-Agent", fmt.Sprintf("Godan-v%s", util.Version))

    resp, err := client.Do(request)
    if err != nil {
        return 0, nil, err
    }

    defer resp.Body.Close()

    var statusCode int 
    statusCode = resp.StatusCode
    if statusCode == 404 {
        return statusCode, nil, Err401 
    } else if statusCode == 401 {
        return statusCode, nil, Err404 
    } else if statusCode == 429 {
        return statusCode, nil, Err429 
    }

    var info ApiInfoS
    if err := json.NewDecoder(resp.Body).Decode(&info); err == nil {
        return statusCode, &info, nil
    }
    return 0, nil, errors.New("Failed")
}

func Host(ip string, apiKey string, timeout int) (int, *HostS, error) {

    client := NewHttpClient(timeout)

    var url string
    url = fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ip, apiKey)

    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, nil, err
    }

    request.Header.Set("User-Agent", fmt.Sprintf("Godan-v%s", util.Version))

    resp, err := client.Do(request)
    if err != nil {
        return 0, nil, err
    }

    defer resp.Body.Close()

    var statusCode int 
    statusCode = resp.StatusCode
    if statusCode == 404 {
        return statusCode, nil, Err404 
    } else if statusCode == 401 {
        return statusCode, nil, Err401 
    } else if statusCode == 429 {
        return statusCode, nil, Err429
    }

    var host HostS
    if err := json.NewDecoder(resp.Body).Decode(&host); err == nil {
        return statusCode, &host, nil
    }
    return 0, nil, errors.New("Failed")
}

func Profile(apiKey string, timeout int) (int, *ProfileS, error) {

    client := NewHttpClient(timeout)

    var url string
    url = fmt.Sprintf("https://api.shodan.io/account/profile?key=%s",
        apiKey)

    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, nil, err
    }

    request.Header.Set("User-Agent", fmt.Sprintf("Godan-v%s", util.Version))

    resp, err := client.Do(request)
    if err != nil {
        return 0, nil, err
    }

    defer resp.Body.Close()

    var statusCode int
    statusCode = resp.StatusCode
    if statusCode == 404 {
        return statusCode, nil, Err404 
    } else if statusCode == 401 {
        return statusCode, nil, Err401 
    } else if statusCode == 429 {
        return statusCode, nil, Err429
    }

    var profile ProfileS
    if err := json.NewDecoder(resp.Body).Decode(&profile); err == nil {
        return statusCode, &profile, nil
    }
    return 0, nil, errors.New("Failed")
}