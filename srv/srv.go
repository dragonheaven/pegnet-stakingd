package srv

import (
	"../config"
	"../node"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strings"

	jrpc "github.com/AdamSLevy/jsonrpc2/v13"
	"github.com/rs/cors"
)

type APIServer struct {
	Node   *node.Pegnetd
	Config *viper.Viper
}

func NewAPIServer(conf *viper.Viper, n *node.Pegnetd) *APIServer {
	s := new(APIServer)
	s.Node = n
	s.Config = conf

	return s
}

// Start the server in its own goroutine. If stop is closed, the server is
// closed and any goroutines will exit. The done channel is closed when the
// server exits for any reason. If the done channel is closed before the stop
// channel is closed, an error occurred. Errors are logged.
func (s *APIServer) Start(stop <-chan struct{}) (done <-chan struct{}) {
	// Set up JSON RPC 2.0 handler with correct headers.
	jrpc.DebugMethodFunc = true
	jrpcHandler := jrpc.HTTPRequestHandler(s.jrpcMethods(), nil)

	var handler http.Handler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			jrpcHandler(w, r)
		})

	// Set up server.
	srvMux := http.NewServeMux()

	srvMux.Handle("/", handler)
	srvMux.Handle("/v1", handler)

	cors := cors.New(cors.Options{AllowedOrigins: []string{"*"}})
	srv = http.Server{Handler: cors.Handler(srvMux)}

	if strings.Contains(s.Config.GetString(config.APIListen), ":") {
		// This means the use set the listen address rather than just the port
		srv.Addr = s.Config.GetString(config.APIListen)
	} else {
		// Set the full listen address from the port
		srv.Addr = fmt.Sprintf(":%d", s.Config.GetInt(config.APIListen))
	}

	// Start server.
	_done := make(chan struct{})
	log.Infof("Listening on %v...", srv.Addr)
	go func() {
		var err error
		// TODO: Renable tls
		//if flag.HasTLS {
		//	err = srv.ListenAndServeTLS(flag.TLSCertFile, flag.TLSKeyFile)
		//} else {
		err = srv.ListenAndServe()
		//}
		if err != http.ErrServerClosed {
			log.Errorf("srv.ListenAndServe(): %v", err)
		}
		close(_done)
	}()
	// Listen for stop signal.
	go func() {
		select {
		case <-stop:
			if err := srv.Shutdown(nil); err != nil {
				log.Errorf("srv.Shutdown(): %v", err)
			}
		case <-_done:
		}
	}()
	return _done
}
