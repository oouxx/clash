package vless

import (
	"net"

	"github.com/gofrs/uuid"
	"github.com/oouxx/clash/component/vmess"
)

const (
	XRO          = "xtls-rprx-origin"
	XRD          = "xtls-rprx-direct"
	XRS          = "xtls-rprx-splice"
	Version byte = 0 // protocol version. preview version is 0
)

// Client is vless connection generator
type Client struct {
	UUID   *uuid.UUID
	Addons *Addons
}

// StreamConn return a Conn with net.Conn and DstAddr
func (c *Client) StreamConn(conn net.Conn, dst *vmess.DstAddr) (net.Conn, error) {
	return newConn(conn, c, dst)
}

// NewClient return Client instance
func NewClient(uuidStr string, addons *Addons) (*Client, error) {
	uid, err := uuid.FromString(uuidStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		UUID:   &uid,
		Addons: addons,
	}, nil
}
