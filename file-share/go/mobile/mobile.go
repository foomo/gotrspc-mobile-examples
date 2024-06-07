package mobile

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/foomo/gotsrpc-mobile-examples/file-share/embeddedfrontend"
)

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
	l := log.Default()

	l.Println("starting server")

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
					http.ServeFile(w, r, s.files[len(s.files)-1])
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
