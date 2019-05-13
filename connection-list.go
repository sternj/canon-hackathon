package main

import "sync"

type connectionList struct {
	mut sync.Mutex
	cons []*rtpConnection
}

func (c *connectionList) addConnection(connection *rtpConnection) {
	c.mut.Lock()
	defer c.mut.Unlock()
	newCons := append(c.cons, connection)
	c.cons = newCons
}

func (c *connectionList) getElem(index int) *rtpConnection {
	c.mut.Lock()
	defer c.mut.Unlock()
	if index >= len(c.cons) {
		return nil
	}
	return c.cons[index]
}

func newConnectionList() *connectionList {
	return &connectionList{sync.Mutex{}, make([]*rtpConnection, 0)}
}