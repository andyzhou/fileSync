package iface

import fileSync "github.com/andyzhou/tinySync/pb"

/*
 * interface for sync, main entry
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type ISync interface {
	DirSync(subDir, newSubDir string, isRemove bool, cb func(subDir, newSubDir string, isRemove bool)) bool
	FileRemove(subDir string, fileName string, cb func(subDir, fileName string)) bool
	FileSync(req *fileSync.FileSyncReq, cb func(subDir, fileName string)) bool
	FileDirectSync(orgFile string, destSubDir string, cb func(subDir, fileName string)) bool
	ReadFile(filePath string) *fileSync.FileSyncReq
	AddNode(addr string) bool
	GetManager() IManager
}