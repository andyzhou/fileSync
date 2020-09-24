package iface

import fileSync "github.com/andyzhou/fileSync/pb"

/*
 * interface for rpc client
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IClient interface {
	Quit()
	DirSync(subDir string, isRemove bool) bool
	FileRemove(subDir, fileName string) bool
	FileSync(req *fileSync.FileSyncReq) bool
}

