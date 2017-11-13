package main

import (
		// "encoding/json"
		"log"
		"net/http"
		// "github.com/gorilla/mux"
		)

//json structure for returning fibonacci numbers 
type fiboutput struct {
	N string `json:"n"`
	Nthfibonacci string `json:"fibonacci(n)"`
}


/*
--- way to build mass_friend_list
var query = `https://api.vk.com/method/execute?access_token=${access_token}&code=${code}`;
API.friends.get ({"uid":123123)

--- way to get all friends details function fetchDetails(id) {

--- way to get all the likes, or some portion of likes from a sample

--- get cities by id - function fetchCityInfo(){
*/ 


// func fibHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	nString := vars["n"]
// 	n, fn := fib_processer(nString)
// 	fibout := fiboutput{ N: n, Nthfibonacci: fn}
// 	json.NewEncoder(w).Encode(fibout)
// }

func main() {
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/fib/{n}", fibHandler)

    router := NewRouter()
	log.Println("Listening.........")
	http.ListenAndServe(":8080",router)
}