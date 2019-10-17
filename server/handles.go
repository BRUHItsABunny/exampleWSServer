package server

import "strconv"

func (srv *ExampleHandle) HandleEcho(cmd Request, conn *Connection) {
	if cmd.Message != nil {
		_ = conn.WriteJSON(conn.MakeSuccessV2(cmd, *cmd.Message))
	} else {
		_ = conn.WriteJSON(conn.MakeErrorV2(cmd, 69, "No variables/ not enough variables supplied"))
	}
}

func (srv *ExampleHandle) HandleEchoOld(cmd Command, conn *Connection) {
	if len(cmd.Variables) != 0 {
		_ = conn.WriteJSON(conn.MakeSuccess(cmd, cmd.Variables[0].(string)))
	} else {
		_ = conn.WriteJSON(conn.MakeError(cmd, 69, "No variables/ not enough variables supplied"))
	}
}

func (srv *ExampleHandle) HandleMirror(cmd Request, conn *Connection) {
	if cmd.UserName != nil && cmd.UserAge != nil && cmd.UserSex != nil {
		ageStr := strconv.Itoa(*cmd.UserAge)
		result := "HI, " + *cmd.UserName + "(" + ageStr + ":" + *cmd.UserSex + ")" + "!"
		_ = conn.WriteJSON(conn.MakeSuccessV2(cmd, result))
	} else {
		_ = conn.WriteJSON(conn.MakeErrorV2(cmd, 69, "No variables/ not enough variables supplied"))
	}
}

func (srv *ExampleHandle) HandleMirrorOld(cmd Command, conn *Connection) {
	if len(cmd.Variables) >= 3 {
		/*
			0 -> name string
			1 -> age int
			2 -> sex string
		*/
		// {"command": "echo", "variables": ["johnny", 19, "male"]}
		if !helpCheckString(cmd.Variables[1]) {
			age := int(cmd.Variables[1].(float64))
			name := cmd.Variables[0].(string)
			sex := cmd.Variables[2].(string)

			ageStr := strconv.Itoa(age)
			result := "HI, " + name + "(" + ageStr + ":" + sex + ")" + "!"
			//HI, johnny(19:male)!
			_ = conn.WriteJSON(conn.MakeSuccess(cmd, result))
		} else {
			_ = conn.WriteJSON(conn.MakeError(cmd, 69, "Age is not an integer"))
		}
	} else {
		_ = conn.WriteJSON(conn.MakeError(cmd, 69, "No variables/ not enough variables supplied"))
	}
}
