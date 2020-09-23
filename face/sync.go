package face

import (
	"github.com/andyzhou/fileSync/iface"
	fileSync "github.com/andyzhou/fileSync/pb"
)

/*
 * face for sync service, main entry
 * this is main entry
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type Sync struct {
	manager iface.IManager
	rpc iface.IRpc
}

//construct
func NewSync(
			port int,
			rootPath string,
		) *Sync {
	//self init
	this := &Sync{
		manager:NewManager(),
	}
	//init rpc
	this.rpc = NewRpc(port, rootPath)
	return this
}

//quit
func (f *Sync) Quit() {
	f.rpc.Stop()
}

//file remove from batch node
func (f *Sync) FileRemove(
					subDir string,
					fileName string,
				) bool {
	return f.manager.FileRemove(subDir, fileName)
}

//file sync into batch node
func (f *Sync) FileSync(
					req *fileSync.FileSyncReq,
				) bool {
	return f.manager.FileSync(req)
}

//add rpc node
func (f *Sync) AddNode(
					addr string,
				) bool {
	return f.manager.AddNode(addr)
}

//get manager face
func (f *Sync) GetManager() iface.IManager {
	return f.manager
}