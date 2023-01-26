package iface

import fileSync "github.com/andyzhou/tinysync/pb"

/*
 * interface for rpc client
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IClient interface {
	Quit()
	DirSync(subDir, newSubDir string, isRemove bool) bool
	FileRemove(subDir, fileName string) bool
	FileSync(req *fileSync.FileSyncReq) bool
}

