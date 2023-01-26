package iface

import pb "github.com/andyzhou/tinysync/pb"

/*
 * interface for file
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IFile interface {
	RemoveFile(
			rootPath string,
			subDir string,
			fileName string,
		) bool
	SaveFile(
			rootPath string,
			req *pb.FileSyncReq,
		) bool
}