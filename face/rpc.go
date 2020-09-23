package face

import (
	"fmt"
	fileSync "github.com/andyzhou/fileSync/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
 * face for rpc service
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type Rpc struct {
	addr string //rpc address
	rootPath string
	listener *net.Listener
	service *grpc.Server //rpc service
}

//construct
func NewRpc(
			port int,
			rootPath string,
		) *Rpc {
	//self init
	this := &Rpc{
		addr:fmt.Sprintf(":%d", port),
		rootPath:rootPath,
	}
	//create service
	this.createService()
	return this
}

//stop service
func (f *Rpc) Stop() {
	if f.service != nil {
		f.service.Stop()
	}
	if f.listener != nil {
		(*f.listener).Close()
	}
}

/////////////////
//private func
/////////////////

//start service
func (f *Rpc) start() {
	//basic check
	if f.listener == nil || f.service == nil {
		return
	}

	//service rpc
	go f.beginService()
}

//begin rpc service
func (f *Rpc) beginService() {
	if f.listener == nil {
		return
	}

	//service listen
	err := f.service.Serve(*f.listener)
	if err != nil {
		log.Println("Rpc::beginService failed, err:", err.Error())
		panic(err)
	}
}

//create rpc service
func (f *Rpc) createService() {
	//listen tcp port
	listen, err := net.Listen("tcp", f.addr)
	if err != nil {
		log.Println("Rpc::createService failed, err:", err.Error())
		panic(err.Error())
	}

	//create rpc server
	f.service = grpc.NewServer()

	//register call back
	fileSync.RegisterFileSyncServiceServer(
			f.service,
			NewIRpcCB(f.rootPath),
		)
	f.listener = &listen

	//start service
	f.start()
}

