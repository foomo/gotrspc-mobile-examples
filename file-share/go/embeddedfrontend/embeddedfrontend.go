package embeddedfrontend

import (
	"embed"
	"net/http"
)

//go:embed nextjs nextjs/_next*
var next embed.FS

func GetEmbeddedNextJSFrontendHandler() (http.Handler, error) {
	// dir, err := fs.Sub(next, "nextjs-bin")
	// if err != nil {
	// 	return nil, err
	// }
	return http.FileServer(http.FS(next)), nil
}
