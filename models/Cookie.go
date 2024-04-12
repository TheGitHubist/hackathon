package models

import "net/http"

func CheckSession(w http.ResponseWriter, r *http.Request, data PageData) (PageData, User) {

	cookie, err := r.Cookie("session_token")
	if err != nil {
		data.IsConnected = false
	}
	value := cookie.Value
	if err != nil {
		data.IsConnected = false
	}

	user := LoadUserApi(value)

	return data, user
}
