package main

import (
	"fmt"
	"net/http"
)

/**
 * @Description: handles requests
 * @param w
 * @param r
 */
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // set the header content
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@gregvroberts.com\">support@gregvroberts.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep being sent to an invalid page</p>")
	}
}

func main() {
	/**
	 * @Description: No matter what path they give us, send us to the handlerFunc
	 * @param "/" sets the file pattern
	 * @param handlerFunc call this method everytime this pattern is requested
	 */
	http.HandleFunc("/", handlerFunc)

	/**
	 * @Description: Start the server
	 * @param ":3000" Port/address
	 * @param nil use the built in serve functionality of the http framework
	 */
	http.ListenAndServe(":3000", nil)

}
