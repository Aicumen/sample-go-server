package api

import (
	"encoding/json"
	"net/http"
)

// UserCreatePost ...
func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed Creating user"))
	} else {
		err = AddUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			temp, _ := json.Marshal(err)
			w.Write(temp)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("User Created"))
		}
	}
}

// UserLoginPost ...
func UserLoginPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestUser User
	err := decoder.Decode(&requestUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed login"))
	} else {
		user, err := GetUser(requestUser.User)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			temp, _ := json.Marshal(err)
			w.Write(temp)
		} else {
			if requestUser.Password == user.Password && requestUser.User == requestUser.User {
				key, err := CreateAPIKey(user.User)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					temp, _ := json.Marshal(err)
					w.Write(temp)
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(key))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Incorrect password"))
			}
		}
	}
}

// UserLogoutGet ...
func UserLogoutGet(w http.ResponseWriter, r *http.Request) {

	apikey := r.URL.Query().Get(apikeyQueryParam)
	user, err := GetValueFromKey(apikey)
	if err != nil || user == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Probablty you are not logged in"))
	} else {
		DeleteKeyValuePair(apikey)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged out successfully"))
	}

}

// UserMypostsGet ...
func UserMypostsGet(w http.ResponseWriter, r *http.Request) {
	apikey := r.URL.Query().Get(apikeyQueryParam)
	user, err := GetValueFromKey(apikey)
	if err != nil || user == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Probablty you are not logged in"))
	} else {
		posts, err := UsersPosts(user)
		if err != nil || len(posts) < 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error or No Posts"))
		} else {
			w.WriteHeader(http.StatusOK)
			postsMarshall, err := json.Marshal(posts)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error. Marshal Error"))
			}
			w.Write(postsMarshall)
		}
	}
}
