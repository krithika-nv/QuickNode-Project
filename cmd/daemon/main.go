package main
import (
	"fmt" 
	"net/http"
	"os"
)

var port string

func init() {
	port = os.Getenv("DAEMON_LISTEN_PORT")
	if port == "" {
		port = "8080"
	}
}
func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { 
		fmt.Fprintf(w, "329d4feb-c5c0-4de5-b10c-701b44fbec4f") 
	}) 
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
