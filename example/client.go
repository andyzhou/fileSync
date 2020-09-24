package main

import (
	"fmt"
	"github.com/andyzhou/fileSync/face"
	"github.com/andyzhou/fileSync/iface"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
 * face for example client
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

const (
	RpcHost = "127.0.0.1"
	RpcPort = 6070 //service port
	RootPath = "/data/test"
	SubDir = "t1"

	//
	FileSrcPath = "/Users/diudiu8848/Downloads"
	FileName = "test1.gif"

)

func main() {
	var (
		wg sync.WaitGroup
	)

	//try catch signal
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Kill,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGABRT,
	)

	///signal snatch
	go func(wg *sync.WaitGroup) {
		var needQuit bool
		for {
			if needQuit {
				break
			}
			select {
			case s := <- c:
				log.Println("Get signal of ", s.String())
				wg.Done()
				needQuit = true
			}
		}
	}(&wg)

	//init service
	service := face.NewSync(RpcPort, RootPath)

	//add node
	rpcAddr := fmt.Sprintf("%s:%d", RpcHost, RpcPort)
	service.AddNode(rpcAddr)

	//start wait group
	wg.Add(1)
	fmt.Println("start example...")

	//testing
	go fileTesting(service)

	wg.Wait()
	service.Quit()
	fmt.Println("stop example...")
}

//file testing
func fileTesting(service iface.ISync) {
	//simple sync
	//simpleSync(service)

	//direct sync
	directSync(service)
}

//direct sync
func directSync(service iface.ISync) {
	fileFullPath := fmt.Sprintf("%s/%s", FileSrcPath, FileName)
	bRet := service.FileDirectSync(fileFullPath, SubDir, cb)
	fmt.Println("file directSync result:", bRet)
}

//simple sync
func simpleSync(service iface.ISync) {
	//open file
	fileFullPath := fmt.Sprintf("%s/%s", FileSrcPath, FileName)

	//simple sync
	fileSyncObj := service.ReadFile(fileFullPath)
	if fileSyncObj == nil {
		return
	}
	fileSyncObj.SubDir = SubDir

	//file sync
	bRet := service.FileSync(fileSyncObj, cb)
	fmt.Println("file simpleSync result:", bRet)
}

//set callback
func cb(subDir, fileName string) {
	fmt.Println("cb, subDir:", subDir, ", fileName:", fileName)
}