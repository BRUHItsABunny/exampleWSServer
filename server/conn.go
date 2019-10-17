package server

func (conn *Connection) WriteJSON(msg interface{}) error {
	var err error

	conn.Lock()
	err = conn.Conn.WriteJSON(msg)
	conn.Unlock()

	return err
}

func (conn *Connection) MakeError(cmd Command, code int, message string) Response {
	return Response{BTError: &Error{Command: cmd.Command, Code: code, Message: message}}
}

func (conn *Connection) MakeSuccess(cmd Command, msg string) Response {
	return Response{BTResult: &Result{Command: cmd.Command, Message: msg}}
}

func (conn *Connection) MakeErrorV2(cmd Request, code int, message string) Response {
	return Response{BTError: &Error{Command: cmd.Command, Code: code, Message: message}}
}

func (conn *Connection) MakeSuccessV2(cmd Request, msg string) Response {
	return Response{BTResult: &Result{Command: cmd.Command, Message: msg}}
}
