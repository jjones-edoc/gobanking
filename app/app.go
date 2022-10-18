package app

import "net/http"

func start() {
	//setup routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//start server
	http.ListenAndServe(":4000", nil)
}
