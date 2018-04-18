package main

import (
	"context"
	"net"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/gotcp"
)

type SessionNode struct {
	common.SessionIntranet
}

func (this *SessionNode) Init(conn net.Conn, root context.Context, derived gotcp.ISession) {
	this.SessionIntranet.Init(conn, root, derived)

	// init cmd
}
