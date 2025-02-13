package httpReceiver

import (
	"errors"
	"net"
)


var extensionFile = map[string]string{
	"html":  "text/html; charset=UTF-8",
	"json":  "application/json; charset=utf-8",
	"js":    "text/javascript; charset=UTF-8",
	"css":   "text/css; charset=UTF-8",
	"mapa":  "",
	"jsx":   "text/javascript; charset=UTF-8",
	"wasm":  "application/wasm",

	// Extensiones de contenido imagen
	"glb":  "model/gltf-binary",
	"gif":  "image/gif",
	"svg":  "image/svg+xml",
	"svgr": "",
	"png":  "image/png",
	"jpg":  "image/jpeg",
	"webp": "image/webp",
	"bmp":  "image/bmp",
	"ico":  "image/x-icon",

	// Extensiones de audio
	"mp3": "audio/mpeg",

	// Extensiones de contenido video
	"mp4":  "video/mp4",
	"webm": "video/webm",

	// Extensiones de contenido de documentos
	"pdf":  "application/pdf",
	"txt":  "text/plain; charset=UTF-8",
	"woff":  "font/woff",
	"woff2": "font/woff2",

	// Nuevas extensiones
	"onnx": "",
	"hash": "",
}

func IsExtensionContent(extName string) (value string, found bool) {

	value, found = extensionFile[extName]
	if found {
		return 
	}

	return 
}



func IsLocalIP(ip string)(bool , error) {

	
    addr := net.ParseIP(ip)
    if addr == nil {
        return false, errors.New("error al parsear la ip")
    }

    if ipv4 := addr.To4(); ipv4 != nil {

        // Check if it's a private IPv4 address
        return ipv4[0] == 10 || // Check if it's in the 10.x.x.x range
            (ipv4[0] == 172 && ipv4[1] >= 16 && ipv4[1] <= 31) || // Check if it's in the 172.16.x.x to 172.31.x.x range
            (ipv4[0] == 192 && ipv4[1] == 168) || // Check if it's in the 192.168.x.x range
            ipv4.IsLoopback(), nil // Check if it's a loopback address (e.g. 127.0.0.1) 

    } else if ipv6 := addr.To16(); ipv6 != nil {
        
        // Check if it's a unique local IPv6 address
        return ipv6[0] == 0xfc || // Check if it's in the fc00::/7 range
            ipv6[0] == 0xfd, nil // Check if it's in the fd00::/8 range
    }

    return false , nil
}

func (SK HttpReceiver) GetIp()(ip string , port string , err error){

    ip, port , err = net.SplitHostPort(SK.request.RemoteAddr)
    if err != nil {
        return
    }
    return
}

func (SK HttpReceiver)IsLocalIP() (ok bool , err error) {

	ip , _  , err := SK.GetIp()
    if err != nil {
		return 
    }

	ok , err = IsLocalIP(ip)
    if err != nil {
		return
    }
    
    return
}

