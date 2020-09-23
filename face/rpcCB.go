package face

import (
	"context"
	"errors"
	"github.com/andyzhou/fileSync/iface"
	fileSync "github.com/andyzhou/fileSync/pb"
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

//file remove
func (f *IRpcCB) FileRemove(
					ctx context.Context,
					in *fileSync.FileRemoveReq,
				) (*fileSync.FileSyncResp, error) {
	//check input value
	if in == nil {
		return nil, errors.New("invalid parameter")
	}

	//file remove from local
	f.file.RemoveFile(f.rootPath, in.SubDir, in.FileName)

	//format result
	result := &fileSync.FileSyncResp{
		Success:true,
	}
	return result, nil
}

//file sync
func (f *IRpcCB) FileSync(
					ctx context.Context,
					in *fileSync.FileSyncReq,
				) (*fileSync.FileSyncResp, error) {
	//check input value
	if in == nil {
		return nil, errors.New("invalid parameter")
	}

	//file save into local
	f.file.SaveFile(f.rootPath, in)

	//format result
	result := &fileSync.FileSyncResp{
		Success:true,
	}
	return result, nil
}
