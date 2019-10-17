package server

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func (srv *ExampleHandle) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	rsp.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	conn, err := srv.Upgrade.Upgrade(rsp, req, nil)
	if err == nil {
		srv.Accept(conn)
	} else {
		srv.Log.Errorln(err)
	}
	/*
		if req.Header.Get("User-Agent") == "BUNNYBOI" {
				conn, _ := server.Upgrade.Upgrade(rsp, req, nil)
				server.Accept(conn)
			} else {
				rsp.WriteHeader(500)
			}
	*/
}

func (srv *ExampleHandle) Accept(ws *websocket.Conn) {
	var err error
	var wErr *websocket.CloseError
	var ok bool
	var cmd Request

	dictActions := map[string]func(Request, *Connection){
		"echo":   srv.HandleEcho,
		"mirror": srv.HandleMirror,
	}
	bunCon := &Connection{Conn: ws}
	defer srv.cleanup(bunCon)
	for {
		err = ws.ReadJSON(&cmd)
		if err != nil {
			srv.Log.Errorln(err)
			if wErr, ok = err.(*websocket.CloseError); ok && wErr.Code == websocket.CloseAbnormalClosure {
				return
			} else {
				result := bunCon.MakeErrorV2(cmd, 2, err.Error())
				err := bunCon.WriteJSON(result)
				if err != nil {
					srv.Log.Errorln(err)
				}
			}
			continue
		}
		if val, ok := dictActions[cmd.Command]; ok {
			val(cmd, bunCon)
		} else {
			result := bunCon.MakeErrorV2(cmd, 1, "No valid command")
			err := bunCon.WriteJSON(result)
			if err != nil {
				srv.Log.Errorln(err)
			}
		}
	}
}

func (srv *ExampleHandle) cleanup(conn *Connection) {
	conn.Lock()
	err := conn.Conn.Close()
	if err != nil {
		srv.Log.Errorln(err)
	}
	conn.Unlock()
}
