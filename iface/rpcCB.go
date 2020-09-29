package iface

import (
	"context"
	fileSync "github.com/andyzhou/tinySync/pb"
)

/*
 * interface for rpc call back
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IRpcCB interface {
	DirSync(
		ctx context.Context,
		in *fileSync.DirSyncReq,
	) (*fileSync.SyncResp, error)
	FileRemove(
		ctx context.Context,
		in *fileSync.FileRemoveReq,
	) (*fileSync.SyncResp, error)
	FileSync(
		ctx context.Context,
		in *fileSync.FileSyncReq,
	) (*fileSync.SyncResp, error)
}

