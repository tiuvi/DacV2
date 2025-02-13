package dan
import (
    "net"
)



func MyLocalIp()string {

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if v.IP.IsLoopback() {
					continue
				}
				if v.IP.To4() != nil {
					return v.IP.String()
				}
			}
		}
	}
	return ""
}





