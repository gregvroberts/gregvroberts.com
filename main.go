package main

import (
	"fmt"
	"net/http"
)

/**
 * @Description:
 * @param w
 * @param r
 */
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
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
