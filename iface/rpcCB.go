package iface

import (
	"context"
	fileSync "github.com/andyzhou/fileSync/pb"
)

/*
 * interface for rpc call back
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IRpcCB interface {
	FileRemove(
		ctx context.Context,
		in *fileSync.FileRemoveReq,
	) (*fileSync.FileSyncResp, error)
	FileSync(
		ctx context.Context,
		in *fileSync.FileSyncReq,
	) (*fileSync.FileSyncResp, error)
}

