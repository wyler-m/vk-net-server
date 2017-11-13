package main

import (
    // "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func details(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    userId := vars["userId"]
    
 //    reader := strings.NewReader(`{"user_id":"325842020"}`)
	// request, err := http.NewRequest("POST", "https://api.vk.com/method/friends.get", reader)
	// // TODO: check err
	// client := &http.Client{}
	// resp, err := client.Do(request)
    url := "https://api.vk.com/method/friends.get"
	resp, err := http.PostForm("https://api.vk.com/method/friends.get",url.Values{"userId": userId})
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

    fmt.Fprintln(w, body)
}