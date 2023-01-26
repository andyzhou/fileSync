package iface

import (
	"context"
	pb "github.com/andyzhou/tinysync/pb"
)

/*
 * interface for rpc call back
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IRpcCB interface {
	DirSync(
		ctx context.Context,
		in *pb.DirSyncReq,
	) (*pb.SyncResp, error)
	FileRemove(
		ctx context.Context,
		in *pb.FileRemoveReq,
	) (*pb.SyncResp, error)
	FileSync(
		ctx context.Context,
		in *pb.FileSyncReq,
	) (*pb.SyncResp, error)
}

