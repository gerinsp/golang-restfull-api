package simple

import "fmt"

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Close Connection", c.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	conn := &Connection{File: file}
	return conn, func() {
		conn.Close()
	}
}
