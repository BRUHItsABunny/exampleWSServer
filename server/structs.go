package server

import (
	"exampleWSServer/database"
	"github.com/BRUHItsABunny/bunnlog"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type ExampleServer struct {
	HTTPServer *http.Server
	Database   *database.ExampleDatabase
	Log        *bunnlog.BunnyLog
}

type ExampleHandle struct {
	Database *database.ExampleDatabase
	Upgrade  websocket.Upgrader
	Log      *bunnlog.BunnyLog
}

type Command struct {
	Command   string        `json:"command"`
	Variables []interface{} `json:"variables"`
}

type Request struct {
	Command  string  `json:"command"`
	UserName *string `json:"username,omitempty"`
	UserAge  *int    `json:"userage,omitempty"`
	UserSex  *string `json:"usersex,omitempty"`
	Message  *string `json:"message,omitempty"`
}

type Connection struct {
	sync.Mutex
	Conn *websocket.Conn
}

type Response struct {
	BTError  *Error  `json:"error,omitempty"`
	BTResult *Result `json:"result,omitempty"`
}

type Error struct {
	Command string `json:"command"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Result struct {
	Command string `json:"command"`
	Message string `json:"message"`
}

func GetExampleHandle(db *database.ExampleDatabase, bLog *bunnlog.BunnyLog) ExampleHandle {
	return ExampleHandle{Database: db, Upgrade: websocket.Upgrader{CheckOrigin: helpCheckOrigin}, Log: bLog}
}

func GetExampleServer(bLog *bunnlog.BunnyLog) (*ExampleServer, error) {
	var err error
	var db *database.ExampleDatabase
	var mux *http.ServeMux
	var srv *http.Server
	var server *ExampleServer

	db, err = database.GetExampleDatabase(bLog)
	if err == nil {
		serverHandler := GetExampleHandle(db, bLog)
		mux = http.NewServeMux()
		mux.HandleFunc("/bunny", serverHandler.ServeHTTP)
		/*
			cfg := &tls.Config{
					MinVersion:               tls.VersionTLS12,
					CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
					PreferServerCipherSuites: true,
					CipherSuites: []uint16{
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
						tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_RSA_WITH_AES_256_CBC_SHA,
					},
				}
		*/
		srv = &http.Server{
			Addr:    ":80",
			Handler: mux,
		}
		server = &ExampleServer{HTTPServer: srv, Database: db, Log: bLog}
		return server, nil
	}
	return &ExampleServer{}, err
}
