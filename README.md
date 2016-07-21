# Contiv-UI Session Manager 
### Constructed using gorilla/mux and gorilla/sessions
### Check the below link for details on the above packages

```
http://www.gorillatoolkit.org/pkg/mux
http://www.gorillatoolkit.org/pkg/sessions

```

### Things that are completed

* Currently the Contiv-UI application states are session managed.
* Session manager will be listening on port 8086 and right now supports 4 functions :
	* Login :(Post Request) Validates user name, password, creates a new session and stores userName, userRole and login time in it. The session cookie is encrypted and is sent to the client.
A response object is used to talk to the client regarding the state of the user.
	* Logout : (Delete Request) Terminates the session and delete the session cookie on the client.
	* ValidateAppState : (Get Request) This function would validate the session cookie, It first decrypts it and checks if the request is valid or not through the last modification time.
ValidateAppState will be called for every state change in the UI.


### Things that are pending
* Code refactoring is needed.
* Should Integrate the session manager with etcd for storing user details.
* I am currently adding the API calls of netplugin and volplugin to flow through the session manager proxy.
* Integration with active directory.
