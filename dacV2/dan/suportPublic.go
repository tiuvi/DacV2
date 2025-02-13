package dan

import (
	"net"
	"net/http"
	"io"
	"strings"

)

func IsNotSuchHost(errorUnwrap string)bool{

	return strings.Contains(strings.ToLower(errorUnwrap), "no such host")

}

func IsExistDomain(domain string)(err error) {

	_, err = net.LookupHost(domain)
	if err != nil {
		return 
	}
	return nil
}

func IsDomainPointingToMyIP(domain string) (bool, error) {

	miIp , err := GetPublicIP()
	if err != nil{
		return false, err
	}

    ips, err := net.LookupIP(domain)
    if err != nil {
        return false, nil
    }

    for _, resolvedIP := range ips {
        println(resolvedIP.String() ,net.ParseIP(miIp).String() )
        
        if resolvedIP.Equal(net.ParseIP(miIp)) {
            return true , nil
        }

    }

    return false, nil
}

func GetPublicIP() (string, error) {

    resp, err := http.Get("https://api.ipify.org/?format=txt")
    if err != nil {
        return "", err
    }

    body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
    if err != nil {
        return "", err
    }
    
    return strings.TrimSpace(string(body)), nil
}



func IsCell(domain string)bool{

    ok := strings.HasPrefix(domain, "cell")
    if !ok {
        return false
    }


    return true
}

func IsDomain(domain string) bool {

	count := strings.Count(domain, ".")
    if count == 1{
        return true
    }
    return false
}