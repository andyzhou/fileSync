package iface

import fileSync "github.com/andyzhou/fileSync/pb"

/*
 * interface for inter manager
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IManager interface {
	Quit()

	//for file sync and remove
	FileRemove(subDir, fileName string) bool
	FileSync(req *fileSync.FileSyncReq) bool

	//for rpc node
	RemoveNode(addr string) bool
	AddNode(addr string) bool
}