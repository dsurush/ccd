package app

import (
	"ccs/models"
	"ccs/token"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//LoginHandler is for login
func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//	fmt.Println("login\n")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
//	get := writer.Header().Get("Role")
//	fmt.Println("I am HEADER ROLE = ", get)
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
	response.Role = user.Role
	response.State = user.Status
	response.Name = user.Name
	response.Surname = user.Surname
//	writer.Header().Set(`Role`, user.Role)
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
// Add new User
func (server *MainServer) AddNewUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
func (server *MainServer) EditUser(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
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
func (server *MainServer) SetStateAndDate(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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

	return
}