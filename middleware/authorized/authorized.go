package authorized

import (
	"ccs/token"
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func Authorized(roles []string, payload func(ctx context.Context) interface{}) func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
			auth := payload(request.Context()).(*token.Payload)
			request.Header.Set(`ID`, strconv.Itoa(int(auth.Id)))
			//fmt.Println("AUTH ID = ", auth.Id)
			for _, role := range roles {
				if role == auth.Role {
			//		fmt.Printf("auth role = %s\n", auth.Role)
					//	log.Printf("access granted %v %v", roles, auth)
					next(writer, request, pr)
					return
				}
			}
			http.Error(writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
	}
}
