package tinysync

import (
	"github.com/andyzhou/tinysync/face"
	"github.com/andyzhou/tinysync/iface"
	pb "github.com/andyzhou/tinysync/pb"
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
			port int, //rpc service port
			rootPath string,
		) *Sync {
	//self init
	this := &Sync{
		manager: face.NewManager(),
	}
	if port > 0 {
		//init rpc service
		this.rpc = face.NewRpc(port, rootPath)
	}
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
				cb func(subDir, newSubDir string, isRemove bool),
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
					req *pb.FileSyncReq,
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
				) *pb.FileSyncReq {
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
	result := &pb.FileSyncReq{
		FileName:fileName,
		FileSize:int64(len(byteData)),
		FileContent:byteData,
	}
	return result
}