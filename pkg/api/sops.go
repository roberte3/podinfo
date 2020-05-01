//sops.go
package api

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) fetchSops(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fetchSops")
	secretString := os.Getenv("superSecret")
	responseString := fmt.Sprintf("secret: %v", secretString)
	fmt.Fprintf(w, responseString)
}
