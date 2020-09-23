package iface

import fileSync "github.com/andyzhou/fileSync/pb"

/*
 * interface for sync, main entry
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type ISync interface {
	FileRemove(subDir string, fileName string) bool
	FileSync(req *fileSync.FileSyncReq) bool
	AddNode(addr string) bool
	GetManager() IManager
}