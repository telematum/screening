package router

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"screening/handler"
	"screening/utils"
)

func CreateRoute(conn *sql.DB) {
	var responseData utils.ResponseData
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}
			var user utils.CreateUserType
			err = json.Unmarshal(body, &user)
			if user.Name == "" {
				responseData.Success = false
				responseData.Message = "Name attribute missing"
			} else if user.Email == "" {
				responseData.Success = false
				responseData.Message = "Email attribute missing"
			} else {
				err = handler.CreateUser(conn, user.Name, user.Email)
				if err != nil {
					responseData.Success = false
					responseData.Message = "User Create Error!"
				} else {
					responseData.Success = true
					responseData.Message = "User Created!"
				}
			}
		} else {
			responseData.Success = false
			responseData.Message = "Method not found!"
		}
		jsonData, jsonError := json.Marshal(responseData)
		if jsonError != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
}