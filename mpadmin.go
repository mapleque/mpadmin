package mpadmin

import (
	"fmt"

	"github.com/mapleque/kelp/http"
	"github.com/mapleque/kelp/logger"
	"github.com/mapleque/kelp/mysql"
)

type Server struct {
	logger logger.Loggerer
	conn   mysql.Connector
	token  string
}

func New(log logger.Loggerer, conn mysql.Connector) *Server {
	return &Server{
		logger: log,
		conn:   conn,
	}
}

func (this *Server) Run(host, token string) {
	this.token = token
	server := http.New(host)
	server.Use(http.LogHandler)
	server.Use(http.RecoveryHandler)

	this.initRouter(server)
	fmt.Println("http server listen on", host)
	server.Run()
}

func (this *Server) initRouter(root *http.Server) {
	// init your router here
	root.Use(this.AuthToken)
	root.Use(this.AuthUser)
	wxapp := root.Group("/wxapp")
	{
		wxapp.Handle("添加app", "/create", this.WXAppCreate)
		wxapp.Handle("修改app", "/update", this.WXAppUpdate)
		wxapp.Handle("删除app", "/delete", this.WXAppDelete)
		wxapp.Handle("查询app", "/retrieve", this.WXAppRetrieve)
	}
}
