package transport

import (
	"context"
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net/http"
	"regexp"
)

func InitRestService(address, port string, grpcClientCon *grpc.ClientConn) error {
	ctx := context.Background()
	rMux := runtime.NewServeMux(runtime.WithHealthEndpointAt(grpc_health_v1.NewHealthClient(grpcClientCon), "/health"))
	if err := library.RegisterBookServiceHandler(ctx, rMux, grpcClientCon); err != nil {
		return err
	}
	if err := library.RegisterUserServiceHandler(ctx, rMux, grpcClientCon); err != nil {
		return err
	}

	mux := http.NewServeMux()
	handlers(mux, rMux)

	log.Printf("gateway server ran on %s:%s", address, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux); err != nil {
		return err
	}
	return nil
}

func handlers(mux *http.ServeMux, rMux *runtime.ServeMux) {
	mux.Handle("/", rMux)
	mux.HandleFunc("/swagger.json", serveSwagger)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(http.Dir("api/swagger/swagger-ui"))))
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/swagger/library.swagger.json")
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}
