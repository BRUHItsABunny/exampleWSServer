package main

import (
	"examplesocket/server"
	"github.com/BRUHItsABunny/bunnlog"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bLog := bunnlog.GetBunnLog(bunnlog.VerbosityDEBUG, log.Ldate|log.Ltime)
	bLog.SetOutputTerminal()
	srv, err := server.GetExampleServer(&bLog)
	if err == nil {
		bLog.Infoln("Server running on 127.0.0.1:80")
		go func() {
			_ = srv.HTTPServer.ListenAndServe()
		}()
	} else {
		bLog.Errorln(err.Error())
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	bLog.Infoln("Shutting down now...")
}
