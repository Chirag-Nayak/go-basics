package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Good Bye ! ! !")
	fmt.Fprint(w, "Good Bye from handlers.ServeHTTP ! ! !")
}
