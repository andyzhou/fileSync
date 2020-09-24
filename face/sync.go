package face

import (
	"github.com/andyzhou/fileSync/iface"
	fileSync "github.com/andyzhou/fileSync/pb"
	"io/ioutil"
	"log"
	"strings"
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

//dir sync for batch node
func (f *Sync) DirSync(
				subDir string,
				newSubDir string,
				isRemove bool,
				cb func(subDir string, isRemove bool),
			) bool {
	return f.manager.DirSync(subDir, newSubDir, isRemove, cb)
}

//file remove from batch node
func (f *Sync) FileRemove(
					subDir string,
					fileName string,
					cb func(subDir, fileName string),
				) bool {
	return f.manager.FileRemove(subDir, fileName, cb)
}

//file direct sync into batch node
func (f *Sync) FileDirectSync(
					orgFile string,
					destSubDir string,
					cb func(subDir, fileName string),
				) bool {
	//basic check
	if orgFile == "" {
		return false
	}

	//try read file
	fileSyncObj := f.ReadFile(orgFile)
	if fileSyncObj == nil {
		return false
	}

	//set dest sub dir
	fileSyncObj.SubDir = destSubDir

	return f.manager.FileSync(fileSyncObj, cb)
}

//file sync into batch node
//cb for success
func (f *Sync) FileSync(
					req *fileSync.FileSyncReq,
					cb func(subDir, fileName string),
				) bool {
	return f.manager.FileSync(req, cb)
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

//read original file
func (f *Sync) ReadFile(
					filePath string,
				) *fileSync.FileSyncReq {
	//basic check
	if filePath == "" {
		return nil
	}

	//get file name
	tempSlice := strings.Split(filePath, "/")
	if tempSlice == nil || len(tempSlice) <= 0 {
		return nil
	}

	tempLen := len(tempSlice)
	fileName := tempSlice[tempLen - 1]
	if fileName == "" {
		return nil
	}

	//try read file
	byteData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("readFile failed, err:", err.Error())
		return nil
	}

	//format result
	result := &fileSync.FileSyncReq{
		FileName:fileName,
		FileSize:int64(len(byteData)),
		FileContent:byteData,
	}
	return result
}