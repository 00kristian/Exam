package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/00kristian/Exam/proto"
	"google.golang.org/grpc"
)

type Server struct{
	Name int32 
	Port int32 
	Peerport int32
}

var hashmap map[int64] int64
var lock sync.Mutex

func main () {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Arguments required: <name> <port> <peer address>")
		os.Exit(1)
	}
	name, err := strconv.Atoi(args[0])
	listenAddr, err1 := strconv.Atoi(args[1])
	peerPort, err2 := strconv.Atoi(args[2])

	if err != nil || err1 != nil || err2 != nil {
		fmt.Println("Arguments has to be Integers")
		os.Exit(1)
	}
	hashmap = make(map[int64]int64)

	s := Server{Name: int32(name), Port: int32(listenAddr), Peerport: int32(peerPort)}

	s.Start()

	for {
		s.Dial()
		time.Sleep(5 * time.Second)
	}


	


}

func (ser *Server) Start(){
	lis, err := net.Listen("tcp", toAddr(ser.Peerport))
	if err != nil {
		fmt.Printf("kommer keg ehr")
		log.Fatalf("failed to listen: %v", err)

	}
	s := grpc.NewServer()
	proto.RegisterHashtableServer(s, ser)
	fmt.Printf("kommer her ey")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}




}


func (s *Server) Put( ctx context.Context, in *proto.Keyvalue) (*proto.Result, error){
	connection, err := grpc.Dial(toAddr(s.Peerport), grpc.WithInsecure())
	

	if err != nil {
		log.Printf("could not get Access: %v", err)
		if s.Peerport == 8300 {
			s.Peerport = 8500
			return &proto.Result{Result: false}, nil 
		} else {
			s.Peerport = s.Peerport - 100
			return &proto.Result{Result: false}, nil
		}
	}
	defer connection.Close()

	client := proto.NewHashtableClient(connection)

	lock.Lock()
	hashmap[int64(in.Key)]= int64(in.Value)
	lock.Unlock()

	req := proto.Keyvalue{Key: in.Key, Value: in.Value}
	client.Put(ctx, &req)

	

	return &proto.Result{Result: true}, nil


}

func (s *Server) Get(ctx context.Context, in *proto.GetValue) (proto.GetValue, error){
	connection, err := grpc.Dial(toAddr(s.Peerport), grpc.WithInsecure())
		

		if err != nil {
			log.Printf("could not get Access: %v", err)
			if s.Peerport == 8300 {
				s.Peerport = 8500
				return proto.GetValue{Key: 0}, nil
				} else {
				s.Peerport = s.Peerport - 100
				return proto.GetValue{Key: 0}, nil
			}
		}
		defer connection.Close()

		lock.Lock()

		var value = hashmap[int64(in.Key)]
		
		lock.Unlock()

		client := proto.NewHashtableClient(connection)
		req := proto.GetValue{Key: int64(value)}
		client.Get(ctx, &req)

		return proto.GetValue{Key: value}, nil

}

	




func toAddr(port int32) string {
	return fmt.Sprintf("localhost:%v", port)
}





func (ser *Server) Dial() {
	connection, err := grpc.Dial(toAddr(ser.Peerport), grpc.WithInsecure())
	
	if err != nil {
		log.Printf("could not get Access: %v", err)
		fmt.Print("hjjee")
		if ser.Peerport == 8300 {
			ser.Peerport = 8500
			return
		} else {
			ser.Peerport = ser.Peerport - 100
			return
		}
	}

	fmt.Printf("hej")
	defer connection.Close()


	ctx := context.Background()
	client := proto.NewHashtableClient(connection)
	req := proto.GetValue{Key: 1}

	for {
		log.Printf("Checking if servers are synced\n")
		_, err := client.Get(ctx, &req)
		time.Sleep(2 * time.Second)

		log.Printf("The number on hash 1 is %d", hashmap[1])


		if err != nil {
			log.Printf("Lost connection. Dialed to new server")
			if ser.Peerport == 8300 {
				ser.Peerport = 8500
				return
			} else {
				ser.Peerport = ser.Peerport - 100
				return
			}
		}
		time.Sleep(5 * time.Second)

	}
}