package dan

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

func generateSelfSignedCert(host string) (tls.Certificate, error) {

    // 1. Generar clave privada RSA
    privKey, err := rsa.GenerateKey(rand.Reader, 4096)
    if err != nil {
        return tls.Certificate{}, err
    }

    // 2. Crear plantilla de certificado
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
		DNSNames:             []string{host}, 
        Subject: pkix.Name{
            CommonName:   host,
        },
        NotBefore: time.Now(),
        NotAfter:  time.Now().Add(365 * 24 * time.Hour),
        KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
        ExtKeyUsage: []x509.ExtKeyUsage{
            x509.ExtKeyUsageServerAuth, // Uso para servidor
        },
        BasicConstraintsValid: true,
   
    }

    // 3. Crear certificado autofirmado
    derBytes, err := x509.CreateCertificate(
        rand.Reader,
        &template,
        &template,
        &privKey.PublicKey,
        privKey,
    )
    if err != nil {
        return tls.Certificate{}, err
    }

    // 4. Codificar a PEM
    certPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "CERTIFICATE",
        Bytes: derBytes,
    })
    
    keyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(privKey),
    })

    // 5. Cargar certificado en formato TLS
    cert, err := tls.X509KeyPair(certPEM, keyPEM)
    if err != nil {
        return tls.Certificate{}, err
    }

    return cert, nil
}


func NewDanSpaceAutocert(host string , port uint16) (dan *DanSpace, err error) {

	mux := http.NewServeMux()

	cert , err := generateSelfSignedCert(host)
	if err != nil {
		return
	}

	dan = &DanSpace{

		Mux: mux,
		Server: &http.Server{

			Addr: ":" + strconv.FormatUint(uint64(port), 10),

					//TimeoutHandler devuelve un Handler que ejecuta h con el límite de tiempo dado.
			Handler: http.TimeoutHandler(mux, 30*time.Second, "Timeout!\n"),

			//ReadTimeout es la duración máxima para leer la solicitud completa,
			ReadTimeout: 30 * time.Second,

			//ReadHeaderTimeout es la cantidad de tiempo permitido para leer los encabezados de solicitud.
			ReadHeaderTimeout: 30 * time.Second,

			//WriteTimeout es la duración máxima antes de que se agote el tiempo de escritura de la respuesta.
			WriteTimeout: 30 * time.Second,

			// IdleTimeout es la cantidad máxima de tiempo para esperar la próxima solicitud cuando se habilita la función Keep-Alives. Si IdleTimeout es cero, se utiliza el valor de ReadTimeout.
			IdleTimeout: 60 * time.Second,

			// MaxHeaderBytes controla la cantidad máxima de bytes que el servidor leerá al analizar las claves y valores del encabezado de la solicitud, incluida la línea de solicitud.
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
			
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS12,
			},
		},
	}

	return dan , nil
}

func (dan *DanSpace) InitDanAutocert() {

	err := dan.Server.ListenAndServeTLS("", "")
	if err != nil {
		println(err.Error())
	}
}