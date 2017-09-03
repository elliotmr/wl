package wl

import (
	"net"
	"os"
	"github.com/pkg/errors"
	"sync"
	"sync/atomic"
)

var atomicIDCounter uint32

func GetNewID() ObjectID {
	// subtraction of one required to start id at 0
	return ObjectID(atomic.AddUint32(&atomicIDCounter, 1) - 1)
}

type ObjectID uint32

func (oid ObjectID) ID() uint32 {
	return uint32(oid)
}

type Object interface {
	ID() uint32
}

type Client struct {
	conn net.Conn
	display *Display
	mutex *sync.Mutex
	cond *sync.Cond
	objects map[ObjectID]Object
}

func (c *Client) Connect(sockName string) error {
	// TODO(mde): Add  support for connecting to an open file descriptor
	if sockName ==  "" {
		sockName = os.Getenv("WAYLAND_DISPLAY")
	}
	if sockName == "" {
		sockName = "wayland-0"
	}
	addr, err := net.ResolveUnixAddr("unix", sockName)
	if err != nil {
		return errors.Wrapf(err, "unable to resolve unix socket address (%s)", sockName)
	}
	c.conn, err = net.DialUnix("unix", nil, addr)
	if err != nil {
		return errors.Wrapf(err, "unable to connect to wayland server at (%s)", sockName)
	}

	c.display = &Display{ObjectID: GetNewID()}
	c.mutex = &sync.Mutex{}
	c.cond = sync.NewCond(&sync.Mutex{})
	c.objects = make(map[ObjectID]Object)
	c.objects[c.display.ObjectID] = c.display

	return nil
}

