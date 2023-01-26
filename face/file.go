package face

import (
	"fmt"
	pb "github.com/andyzhou/tinysync/pb"
	"io/ioutil"
	"log"
	"os"
)

/*
 * face for file storage
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type File struct {
}

//construct
func NewFile() *File {
	//self init
	this := &File{
	}
	return this
}

//remove file from local
func (f *File) RemoveFile(
				rootPath string,
				subDir string,
				fileName string,
			) bool {
	//basic check
	if rootPath == "" || subDir == "" || fileName == "" {
		return false
	}

	//format full path
	filePath := fmt.Sprintf("%s/%s/%s", rootPath, subDir, fileName)

	//remove file
	err := os.Remove(filePath)
	if err != nil {
		log.Println("File::RemoveFile failed, err:", err.Error())
		return false
	}

	return true
}

//save file into local
func (f *File) SaveFile(
				rootPath string,
				req *pb.FileSyncReq,
			) bool {
	//basic check
	if rootPath == "" || req == nil {
		return false
	}

	//format full path
	fileDir := fmt.Sprintf("%s/%s", rootPath, req.SubDir)
	filePath := fmt.Sprintf("%s/%s", fileDir, req.FileName)

	//check sub dir
	err := f.checkOrCreateDir(fileDir)
	if err != nil {
		log.Println("File::SaveFile failed, err:", err.Error())
		return false
	}

	//save file content
	err = ioutil.WriteFile(filePath, req.FileContent, os.ModePerm)
	if err != nil {
		log.Println("File::SaveFile failed, err:", err.Error())
		return false
	}

	return true
}

/////////////////
//private func
/////////////////

//check or create dir
func (f *File) checkOrCreateDir(
					dir string,
				) error {
	_, err := os.Stat(dir)
	if err == nil {
		return err
	}
	bRet := os.IsExist(err)
	if bRet {
		return nil
	}
	err = os.Mkdir(dir, 0777)
	return err
}