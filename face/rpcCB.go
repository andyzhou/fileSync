package face

import (
	"context"
	"errors"
	"fmt"
	"github.com/andyzhou/tinySync/iface"
	fileSync "github.com/andyzhou/tinySync/pb"
	"log"
	"os"
	"strings"
)

/*
 * face for rpc call back
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type IRpcCB struct {
	rootPath string
	file iface.IFile
}

//construct
func NewIRpcCB(
			rootPath string,
		) *IRpcCB {
	//self init
	this := &IRpcCB{
		rootPath:rootPath,
		file:NewFile(),
	}
	return this
}

//////////////////////
//call backs for rpc
/////////////////////

//dir sync
func (f *IRpcCB) DirSync(
					ctx context.Context,
					in *fileSync.DirSyncReq,
				) (*fileSync.SyncResp, error) {
	var (
		err error
		bRet bool
	)

	//check input value
	if in == nil {
		return nil, errors.New("invalid parameter")
	}

	//dir opt for local
	subDirFull := fmt.Sprintf("%s/%s", f.rootPath, in.SubDir)
	subDirFull = strings.ReplaceAll(subDirFull, "//", "/")
	if in.IsRemove {
		//remove
		_, err = os.Stat(subDirFull)
		if err == nil {
			err = os.Remove(subDirFull)
		}else{
			isExists := os.IsExist(err)
			if isExists {
				err = os.Remove(subDirFull)
			}else{
				err = nil
			}
		}

	}else{
		//add or rename
		if in.NewSubDir != "" {
			//rename
			newDirFull := fmt.Sprintf("%s/%s", f.rootPath, in.NewSubDir)
			err = os.Rename(subDirFull, newDirFull)
		}else{
			//add new
			err = os.Mkdir(subDirFull, 0777)
		}
	}

	if err != nil {
		log.Println("IRpcCB::DirSync failed, err:", err.Error())
		bRet = false
	}else{
		bRet = true
	}

	//format result
	result := &fileSync.SyncResp{
		Success:bRet,
	}
	return result, nil
}

//file remove
func (f *IRpcCB) FileRemove(
					ctx context.Context,
					in *fileSync.FileRemoveReq,
				) (*fileSync.SyncResp, error) {
	//check input value
	if in == nil {
		return nil, errors.New("invalid parameter")
	}

	//file remove from local
	f.file.RemoveFile(f.rootPath, in.SubDir, in.FileName)

	//format result
	result := &fileSync.SyncResp{
		Success:true,
	}
	return result, nil
}

//file sync
func (f *IRpcCB) FileSync(
					ctx context.Context,
					in *fileSync.FileSyncReq,
				) (*fileSync.SyncResp, error) {
	//check input value
	if in == nil {
		return nil, errors.New("invalid parameter")
	}

	//file save into local
	f.file.SaveFile(f.rootPath, in)

	//format result
	result := &fileSync.SyncResp{
		Success:true,
	}
	return result, nil
}
