package face

import (
	"fmt"
	pb "github.com/andyzhou/tinysync/pb"
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
	port int //rpc service port
	addr string //rpc address
	rootPath string
	listener *net.Listener
	service *grpc.Server
}

//construct
func NewRpc(
			port int,
			rootPath string,
		) *Rpc {
	//self init
	this := &Rpc{
		port:port,
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
	if f.port <= 0 {
		return
	}

	//listen tcp port
	listen, err := net.Listen("tcp", f.addr)
	if err != nil {
		log.Println("Rpc::createService failed, err:", err.Error())
		panic(err.Error())
	}

	//create rpc server
	f.service = grpc.NewServer()

	//register call back
	pb.RegisterFileSyncServiceServer(
			f.service,
			NewIRpcCB(f.rootPath),
		)
	f.listener = &listen

	//start service
	f.start()
}

