package app

import (
	"ccs/models"
	"ccs/token"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"time"
)

//LoginHandler is for login
func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody token.RequestDTO
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}

	log.Printf("login = %s, pass = %s\n", requestBody.Username, requestBody.Password)
	response, err := server.tokenSvc.Generate(request.Context(), &requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
		if err != nil {
			log.Print(err)
		}
		return
	}
	user, err := server.tokenSvc.FindUserForPassCheck(requestBody.Username)
	//log.Println(response)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
		if err != nil {
			log.Print(err)
		}
		return
	}
	ok := models.CheckStatusLine(user.StatusLine)
	if ok == false && user.Role == "user"{
		writer.WriteHeader(http.StatusIMUsed)
		return
	}
	const StatusLine = true
	err = server.svc.SetStatusLine(requestBody.Username, StatusLine)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
		if err != nil {
			log.Print(err)
		}
		return
	}
	err = server.svc.SetLoginTime(user.Id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
		if err != nil {
			log.Print(err)
		}
		return
	}
	err = server.svc.SetVisitTime(user.Id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.can't fix Visit times", err.Error()})
		if err != nil {
			log.Print(err)
		}
	}
	response.Role = user.Role
	response.State = user.Status
	response.Name = user.Name 	
	response.Surname = user.Surname
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//GetUser
func (server *MainServer) GetUserByIdHandler(writer http.ResponseWriter, _ *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("I found client By number id")
	id := pr.ByName(`id`)
	//fmt.Println(id)

	response, err := server.svc.GetUserById(id)
	if err != nil {
		fmt.Println("can't take from db")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(response)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//Get users
func (server *MainServer) GetUsersHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	response, err := server.svc.GetUsers()
	if err != nil {
		fmt.Println("can't take from db")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//Get users with work time
func (server *MainServer) GetUsersWithWorkTimeHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	//TODO: REFACTOR
	type UsersTimeResponse struct {
		Users []models.UserWithWorkTimeDTO `json:"users"`
		TimeUnix int64 `json:"time_unix"`
	}
	var res UsersTimeResponse
	res.TimeUnix = time.Now().Unix()
	//
	response, err := server.svc.GetUsersWithWorkTime()
	res.Users = response
	if err != nil {
		fmt.Println("can't take from db")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(writer).Encode(&res)
	if err != nil {
		log.Print(err)
	}
}
// Add new User
func (server *MainServer) AddNewUserHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody models.SaveUser
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = server.svc.AddNewUser(requestBody)
	if err != nil {
		fmt.Println("Err to add new user")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := requestBody
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
// Edit new User
func (server *MainServer) EditUserHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody models.SaveUser
	id := pr.ByName(`id`)
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = server.svc.EditUser(requestBody, id)
	if err != nil {
		fmt.Println("Err to add new user")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := requestBody
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
// Set Status and Date
func (server *MainServer) SetStateAndDateHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody models.StatesDTO
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	ID := request.Header.Get(`ID`)
	fmt.Println("im id in handler", ID)
	err = server.svc.SetStateAndDate(requestBody, ID)
	if err != nil {
	//	fmt.Println("Err to add new user")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	atoi, err := strconv.Atoi(ID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.can't conver id from string", err.Error()})
		if err != nil {
			log.Print(err)
		}
	}
	err = server.svc.SetVisitTime(int64(atoi))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.can't fix Visit times", err.Error()})
		if err != nil {
			log.Print(err)
		}
	}

	return
}
// Get User Stats
func (server *MainServer) GetUserStatsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	type Time struct {
		Time int64 `json:"time"`
	}
	var requestBody Time
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	ID := request.Header.Get(`ID`)
	requestBody.Time /= 1000
	fmt.Println("IIIIDDDD = is ", ID)
	fmt.Println("TIMEEE + IS ", requestBody.Time)
	response, err := server.svc.GetUserStats(ID, requestBody.Time)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
// Get Users Stats
func (server *MainServer) GetUsersStatsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var interval models.TimeInterval
	interval.From = time.Now().Unix() - 3240000
	interval.To = time.Now().Unix()

	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
	from /= 1000
	interval.From = int64(from)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
	to /= 1000
	interval.To = int64(to)
	}

	fmt.Println(interval.To)
	response, err := server.svc.GetUsersStats(interval)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
// Get User Stat for Admin by from/to
func (server *MainServer) GetUserStatsForAdminHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var interval models.TimeInterval
	startOfDay := models.GetUnixTimeStartOfDay(time.Now())
	interval.From = startOfDay
	interval.To = time.Now().Unix()
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		from /= 1000
		interval.From = int64(from)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		to /= 1000
		interval.To = int64(to)
	}

	id := pr.ByName(`id`)

	response, err := server.svc.GetUserStatsForAdmin(id, interval)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
///
func (server *MainServer) SetNewPassHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody models.ChangePassword
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	ID := request.Header.Get(`ID`)
	fmt.Println("im id in handler", ID)
	err = server.svc.ChangePassword(ID, requestBody.Password, requestBody.NewPassword)
	if err != nil {
	//	fmt.Println("Err to add new user")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (server *MainServer) ExitClickHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	//
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody models.StatesDTO
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println("json_invalie")
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	ID := request.Header.Get(`ID`)
	fmt.Println("im id in handler", ID)

	//
	/// ---- ////

	err = server.svc.ExitClick(ID, requestBody)
	if err != nil {
		//	fmt.Println("Err to add new user")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	atoi, err := strconv.Atoi(ID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.can't conver id from string", err.Error()})
		if err != nil {
			log.Print(err)
		}
	}
	err = server.svc.SetVisitTime(int64(atoi))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.can't fix Visit times", err.Error()})
		if err != nil {
			log.Print(err)
		}
	}
	return
}
//
func (server *MainServer) ReportHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	//TODO: REFACTOR
	type ReportResponse struct {
		Reports []models.Report `json:"reports"`
		TimeUnix int64 `json:"time_unix"`
	}
	var res ReportResponse
	res.TimeUnix = time.Now().Unix()
	//
	from := request.URL.Query().Get(`from`)
	to := request.URL.Query().Get(`to`)

	response, err := server.svc.GetReport(from, to)
	res.Reports = response
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&res)
	if err != nil {
		log.Print(err)
		return
	}
}
func (server *MainServer) StatusConfirmHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody models.StatusConfirm
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	fmt.Println("im id in handler", requestBody)
	//TODO INSERT AND UPDATE
	ID := request.Header.Get(`ID`)
	fmt.Println("im id in handler", ID)
	userId, err := strconv.Atoi(ID)

	err = server.svc.SetActivities(int64(userId), requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.server_connection"})
		log.Print(err)
		return
	}
	return
}
