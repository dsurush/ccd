package jwt

import (
	"context"
	jwtcore "github.com/dsurush/jwt/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type contextKey string

var payloadContextKey = contextKey("jwt")

func JWT(payloadType reflect.Type, secret jwtcore.Secret) func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
			header := request.Header.Get("Authorization")
			if header == "" {
				next(writer, request, param)
				return
			}

			if !strings.HasPrefix(header, "Bearer ") {
				next(writer, request, param)
				return
			}

			token := header[len("Bearer "):]

			ok, err := jwtcore.Verify(token, secret)
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			if !ok {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			payload := reflect.New(payloadType).Interface()

			err = jwtcore.Decode(token, payload)
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			ok, err = jwtcore.IsNotExpired(payload, time.Now())
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			if !ok {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			log.Print(payload)

			ctx := context.WithValue(request.Context(), payloadContextKey, payload)
			next(writer, request.WithContext(ctx), param)
		}
	}
}

func FromContext(ctx context.Context) (payload interface{}) {
	payload = ctx.Value(payloadContextKey)
	return
}

func IsContextNonEmpty(ctx context.Context) bool {
	return nil != ctx.Value(payloadContextKey)
}
