package router

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"screening/handler"
	"screening/utils"
)

// Router to create a user, db connection object as param
func UpdateRouter(conn *sql.DB) {
	logger := utils.Logger()
	var responseData utils.ResponseData
	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		// route only for PATCH method
		if r.Method == http.MethodPatch {
			body, err := ioutil.ReadAll(r.Body)
			logger.Error().Msg("Error reading request body")
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}
			var user utils.UpdateUserType
			err = json.Unmarshal(body, &user)
			if user.Id == 0 {
				responseData.Success = false
				responseData.Message = "id attribute missing"
			} else {
				// Calling the handler method
				err = handler.UpdateUser(conn, user.Name, user.Email, user.Id)
				if err != nil {
					responseData.Success = false
					responseData.Message = "User Updated Error!"
				} else {
					responseData.Success = true
					responseData.Message = "User Updated!"
				}
			}
		} else {
			responseData.Success = false
			responseData.Message = "Method not found!"
		}
		jsonData, jsonError := json.Marshal(responseData)
		if jsonError != nil {
			logger.Error().Msg("Error encoding JSON")
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
}
