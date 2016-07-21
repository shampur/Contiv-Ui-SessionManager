package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

//Global Declarations
var store = sessions.NewCookieStore([]byte("something-very-secret"))
var userObj UserClass

//Handler function for Logging
func loginUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Executing loginUser Handler")
	err := req.ParseForm()
	if err != nil {
		fmt.Println("Error in pasrsing form = ", err)
	}
	var userdetails UserDetails
	userdetails.Name = req.FormValue("name")
	userdetails.Pwd = req.FormValue("pwd")
	if userObj.validateCredentials(userdetails) == true {
		fmt.Println("User Credentials are valid")
		session, errses := store.New(req, "contiv-session")
		if errses != nil {
			fmt.Println("unable to fetch session")
			http.Error(w, errses.Error(), http.StatusInternalServerError)
			return
		}
		prepareSession(session)
		fmt.Println("Session Name = ", session.Name())
		fmt.Println("Session Values =", session.Values)
		fmt.Println("Saving Session")
		fmt.Println("Session = ", session)
		session.Save(req, w)
	}
	//io.WriteString(w, "Hello world!")
	//w.Write([]byte{1, 2, 3, 4})

	responseData, err := json.Marshal(userObj)
	if err != nil {
		fmt.Println("Json decode error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)

	/*
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, "Hello astaxie!")
	*/
}

//Handler Function to validate the application state
func validateAppState(w http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "contiv-session")
	if err != nil {
		fmt.Println("unable to fetch session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if validateSession(session) == true {
		userObj.setDetails(session.Values["userName"].(string), session.Values["role"].(string), "success", "", true, true)
		if err != nil {
			fmt.Println("Unable to decode JSON")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		userObj.setDetails("none", "none", "success", "none", false, false)
		session.Options.MaxAge = -1
	}
	session.Save(req, w)
	responseData, err := json.Marshal(userObj)
	if err != nil {
		fmt.Println("error while building responseObject")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

// Handler function for OPTIONS request
func preFlight(w http.ResponseWriter, req *http.Request) {
}

// Handler function for logging out
func logOutUser(w http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "contiv-session")
	if err != nil {
		fmt.Println("error while decoding session")
	}
	session.Options.MaxAge = -1
	session.Save(req, w)
	userObj.setDetails("none", "none", "success", "none", false, false)
	responseData, err := json.Marshal(userObj)
	if err != nil {
		fmt.Println("error while json encoding")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
