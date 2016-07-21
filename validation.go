package main

import (
	"fmt"
	"time"

	"github.com/gorilla/sessions"
)

func validateSession(session *sessions.Session) bool {
	fmt.Println("Validate Session")
	if session.IsNew != true {
		lastLoginTime := session.Values["lastLoginTime"].(string)
		fmt.Println("lastLoginTime = ", lastLoginTime)
		parsedTime, err := time.Parse(time.RFC3339, lastLoginTime)
		if err != nil {
			fmt.Println("Error while parsing time")
		}
		duration := time.Since(parsedTime)
		fmt.Println("currentTime = ", time.Now().Format(time.RFC3339))
		fmt.Println("Past time = ", parsedTime)
		fmt.Println("Duration eloped = ", duration.Minutes())
		minutesPassed := duration.Minutes()
		if minutesPassed < 0 || minutesPassed > 0.3 {
			return false
		}
		session.Values["lastLoginTime"] = time.Now().Format(time.RFC3339)
		fmt.Println("Session Valid")
		return true
	}
	return false
}

func prepareSession(ses *(sessions.Session)) {
	ses.Values["userName"] = userObj.UserName
	ses.Values["role"] = userObj.UserRole
	ses.Values["lastLoginTime"] = time.Now().Format(time.RFC3339)
}

/*
    return false
		fmt.Println("The session is already present")
		fmt.Println("username = ", session.Values["username"])
		fmt.Println("status = ", session.Values["status"])
		fmt.Println("The session name is =", session.Name())
		fmt.Println("The session ID is =", session.ID)
		fmt.Println("Session Options path = ", session.Options.Path)
	}
  else{
    session.Values["username"] = "charan"
    session.Values["status"] = "Logged In"
  }


	// Set some session values.

	// Save it before we write to the response/return from the handler.
	err = session.Save(req, w)
	if err != nil {
		fmt.Println("saving failed = ", err)
	}
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "charan", Value: "user", Expires: expiration}
	http.SetCookie(w, &cookie)
	io.WriteString(w, "Hello world!")
*/
