package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // set the header content
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Frequently Asked Questions</h1><p>Here is a list of questions that our users commonly ask.</p>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // set the header content
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@gregvroberts.com\">support@gregvroberts.com</a>.")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/contact", contact)

	/**
	 * @Description: Start the server
	 * @param ":3000" Port/address
	 * @param nil use the built in serve functionality of the http framework
	 */
	http.ListenAndServe(":3000", r)

}
