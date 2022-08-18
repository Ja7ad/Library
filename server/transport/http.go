package transport

import (
	"context"
	"expvar"
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/library"
	_ "github.com/Ja7ad/library/server/statik"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net/http"
	"net/http/pprof"
	"regexp"
)

func InitRestService(address, port string, grpcClientCon *grpc.ClientConn) error {
	ctx := context.Background()
	rMux := runtime.NewServeMux(
		runtime.WithHealthEndpointAt(grpc_health_v1.NewHealthClient(grpcClientCon), "/health"),
		runtime.WithIncomingHeaderMatcher(headers),
	)
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

func handlers(mux *http.ServeMux, rMux *runtime.ServeMux) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	mux.Handle("/", cors(rMux))
	mux.HandleFunc("/swagger.json", serveSwagger)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(statikFS)))
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())
	return nil
}

func headers(key string) (string, bool) {
	switch key {
	case "X-Custom-header1":
		return key, true
	case "Header2":
		return key, true
	case "Header3":
		return key, true
	default:
		return key, false
	}
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/swagger.json")
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
