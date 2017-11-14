package main

import (
    // "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "github.com/gorilla/mux"
    "io/ioutil"
)

func get_friend_list(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    userId := vars["userId"]
	resp,err := http.PostForm("https://api.vk.com/method/friends.get",url.Values{"user_id":{userId}})
	if err != nil {	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {	}
	bodyString := string(body)
    fmt.Fprintln(w, bodyString)
}

func friends_of_friends(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    userId := vars["userId"]
	resp,err := http.PostForm("https://api.vk.com/method/friends.get",url.Values{"user_id":{userId}})
	if err != nil {	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {	}
	bodyString := string(body)
    fmt.Fprintln(w, bodyString)
}