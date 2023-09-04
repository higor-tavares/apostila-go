package main
import (
	"fmt"
	"net/http"
	"time"
)
func main() {
	http.HandleFunc(
		"/time",
		func (writer http.ResponseWriter, r *http.Request) {
		//the method Format() use this date to analyze e recognize the format, this does'nt smell good
			fmt.Fprintf(writer, "<h1>Hello! There is %s</h1>", time.Now().Format("02/01/2006 15:04:05"))
		})
	fmt.Printf("The server is running on http://localhost:9090/time \nhit ˆC to exit\n")
	http.ListenAndServe(":9090", nil)
}