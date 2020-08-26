package authorized

import (
	"ccs/token"
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Authorized(roles []string, payload func(ctx context.Context) interface{}) func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
			auth := payload(request.Context()).(*token.Payload)
			for _, role := range roles {
				if role == auth.Role {
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
