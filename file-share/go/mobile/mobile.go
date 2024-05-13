package mobile

import (
	"io"
	"net/http"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/foomo/gotsrpc-mobile-examples/file-share/embeddedfrontend"
)

// derived from : https://raw.githubusercontent.com/golang/go/master/src/crypto/tls/generate_cert.go
// https://raw.githubusercontent.com/foomo/webgrapple/main/pkg/server/selfsign.go
func selfsign(hosts []string, certFile, keyFile string) error {

	const timeFormat = "Jan 2 15:04:05 2006"
	var (
		validFrom = time.Now().Format(timeFormat)
		validFor  = 365 * 24 * time.Hour
		isCA      = false
	)

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return fmt.Errorf("Failed to generate private key: %w", err)
	}

	var notBefore time.Time
	if len(validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse(timeFormat, validFrom)
		if err != nil {
			return fmt.Errorf("Failed to parse creation date: %w", err)
		}
	}

	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("Failed to generate serial number: %w", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		if h == "" {
			// maybe check for /etc/hosts
			// log.Fatalf("Missing required --host parameter")
		}
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return fmt.Errorf("Failed to create certificate: %w", err)
	}

	certOut, err := os.Create(certFile)
	if err != nil {
		return fmt.Errorf("Failed to open cert.pem for writing: %w", err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return fmt.Errorf("Failed to write data to cert.pem: %w", err)
	}
	if err := certOut.Close(); err != nil {
		return fmt.Errorf("Error closing cert.pem: %w", err)
	}

	keyOut, err := os.OpenFile(keyFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Failed to open key.pem for writing: %w", err)
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return fmt.Errorf("Unable to marshal private key: %w", err)
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return fmt.Errorf("Failed to write data to key.pem: %w", err)
	}
	if err := keyOut.Close(); err != nil {
		return fmt.Errorf("Error closing key.pem: %w", err)
	}
	return nil
}

// func RunServer(addr string) {
// 	embeddedFrontendHandler, err := embeddedfrontend.GetEmbeddedNextJSFrontendHandler()
// 	if err != nil {
// 		panic(err)
// 	}
// 	go func() {
// 		for {
// 			fmt.Println("starting server on addr", addr)
// 			err := http.ListenAndServe(addr, embeddedFrontendHandler)
// 			fmt.Println("-----------> server error", err)
// 			time.Sleep(time.Second)
// 		}
// 	}()
// }

// func RunTLSServer(path, addr string) {
// 	// pinning
// 	// 		- certificate pinning "... Pinning cannot loosen the trust requirements of your app — it can only tighten them ..."
// 	// 		- https://developer.apple.com/news/?id=g9ejcf8y
// 	// Creating an Identity for Local Network TLS
// 	// 		- https://developer.apple.com/documentation/network/creating_an_identity_for_local_network_tls
// 	path = strings.TrimPrefix(path, "file://")
// 	embeddedFrontendHandler, err := embeddedfrontend.GetEmbeddedNextJSFrontendHandler()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("creating", path, os.MkdirAll(path, 0777))
// 	certFile, keyFile := filepath.Join(path, "cert.pem"), filepath.Join(path, "key.pem")
// 	fmt.Println("self signinging", selfsign([]string{"localhost"}, certFile, keyFile))
// 	chanCrash := make(chan string)
// 	log, err := os.OpenFile(filepath.Join(path, "go.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	go func() {
// 		errs := []string{}
// 		for {

// 			select {
// 			case nextErr := <-chanCrash:
// 				errs = append(errs, nextErr)
// 			case <-time.After(10 * time.Second):

// 				fmt.Fprintln(log, "all good")
// 			}
// 			fmt.Fprintln(log, "num crashes", len(errs))
// 			for i, err := range errs {
// 				fmt.Fprintln(log, "err", i, err)
// 			}
// 		}
// 	}()
// 	go func() {
// 		fmt.Println("starting server on addr", addr)
// 		err := http.ListenAndServeTLS(addr, certFile, keyFile, embeddedFrontendHandler)
// 		chanCrash <- err.Error()
// 		fmt.Println("-----------> server error", err)
// 		fmt.Fprintln(log, "-----------> server error", err)
// 		time.Sleep(time.Second)
// 	}()
// }

type Server struct {
	handler http.Handler
	files   []string
}

func NewNextJSHandler(addr string) *Server {
	embeddedFrontendHandler, err := embeddedfrontend.GetEmbeddedNextJSFrontendHandler()
	if err != nil {
		panic(err)
	}
	s := &Server{
		handler: embeddedFrontendHandler,
	}
	go func() {
		for {
			fmt.Println("starting server on addr", addr)
			err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("serving:", r.URL.Path)
				switch r.URL.Path {
				case "/serve":
					f, err := os.Open(s.files[0])
					if err != nil {
						http.Error(w, "file thing:"+err.Error(), http.StatusInternalServerError)
						return
					}
					defer f.Close()
					io.Copy(w, f)
				case "/debug":
					w.Write([]byte(spew.Sdump(s)))
				default:
					embeddedFrontendHandler.ServeHTTP(w, r)
				}
			}))
			fmt.Println("-----------> server error", err)
			fmt.Println("files", s.files)
			time.Sleep(time.Second)
		}
	}()
	return s
}

func (s *Server) ExposeFile(file string) {
	s.files = append(s.files, file)
}
