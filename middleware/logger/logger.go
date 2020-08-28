package logger

import (
	"ccs/token"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dsurush/jwt/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func WriteToFile(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Panic("Failed to log to file", err)
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("Can't close file")
		}
	}()
	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalf("Can'")
	}
}

func Logger(prefix string) func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
			var GetToken = ""
			GetToken = request.Header.Get("Authorization")
			var PayloadFields token.Payload
			if GetToken == "" && request.URL.Path != `/api/login` {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			if GetToken != "" && request.URL.Path != `/api/login` {
				parts, err := jwt.SplitToken(GetToken)
				if err != nil {
					http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
				payloadEncoded := parts[1]
				payloadJSON, err := base64.RawURLEncoding.DecodeString(payloadEncoded)
				//var Pl token2.Payload
				if err != nil {
					http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
				err = json.Unmarshal(payloadJSON, &PayloadFields)
				if err != nil {
					http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
			}
			if request.URL.Path == `/api/login` {
				request.Header.Set(`ID`, string(PayloadFields.Id))
			}
			//get := request.Header.Get(`ID`)
			//fmt.Println(`I am Payloads fields = `, get)
			log.Printf(
				"%s: [%s] login: %s %s Method: %s, path: %s\n",
				time.Now().Format("2006-01-02 15:04:05"),
				strings.Split(request.RemoteAddr, ":")[0],
				PayloadFields.Login,
				prefix,
				request.Method,
				request.URL.Path,
			)
			data := fmt.Sprintf("%s: [%s] login: %s %s Method: %s, path: %s\n",
				time.Now().Format("2006-01-02 15:04:05"),
				strings.Split(request.RemoteAddr, ":")[0],
				PayloadFields.Login,
				prefix,
				request.Method,
				request.URL.Path,
			)
			WriteToFile(`logfile.log`, data)
			next(writer, request, pr)
		}
	}
}