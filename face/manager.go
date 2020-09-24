package face

import (
	fileSync "github.com/andyzhou/fileSync/pb"
	"sync"
)

/*
 * face for inter manager
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type Manager struct {
	clients *sync.Map
}

//construct
func NewManager() *Manager{
	//self init
	this := &Manager{
		clients:new(sync.Map),
	}
	return this
}

//quit
func (f *Manager) Quit() {
	if f.clients == nil {
		return
	}
	sf := func(_, v interface{}) bool {
		client, ok := v.(*Client)
		if !ok {
			return false
		}
		client.Quit()
		return true
	}
	f.clients.Range(sf)
}

//file remove from all clients
func (f *Manager) FileRemove(
					subDir string,
					fileName string,
					cb func(subDir, fileName string),
				) bool {
	//basic check
	if subDir == "" || fileName == "" {
		return false
	}

	if f.clients == nil {
		return false
	}

	//do doc sync on all clients
	sf := func(k, v interface{}) bool {
		client, ok := v.(*Client)
		if !ok {
			return false
		}
		bRet := client.FileRemove(subDir, fileName)
		if !bRet {
			return false
		}
		//run callback
		if cb != nil {
			cb(subDir, fileName)
		}
		return true
	}
	f.clients.Range(sf)

	return true
}

//file sync to all clients
func (f *Manager) FileSync(
					req *fileSync.FileSyncReq,
					cb func(subDir, fileName string),
				) bool {
	//basic check
	if req == nil {
		return false
	}

	if f.clients == nil {
		return false
	}

	//do doc sync on all clients
	sf := func(k, v interface{}) bool {
		client, ok := v.(*Client)
		if !ok {
			return false
		}
		bRet := client.FileSync(req)
		if !bRet {
			return false
		}
		//run call back
		if cb != nil {
			cb(req.SubDir, req.FileName)
		}
		return true
	}
	f.clients.Range(sf)

	return true
}

//remove client node
func (f *Manager) RemoveNode(
					addr string,
				) bool {
	//basic check
	if addr == "" || f.clients == nil {
		return false
	}

	//remove
	f.clients.Delete(addr)

	return true
}

//add client node
func (f *Manager) AddNode(
					addr string,
				) bool {
	//basic check
	if addr == "" || f.clients == nil {
		return false
	}

	//check record
	_, ok := f.clients.Load(addr)
	if ok {
		return false
	}

	//init new client
	client := NewClient(addr)

	//sync into map
	f.clients.Store(addr, client)

	return true
}
