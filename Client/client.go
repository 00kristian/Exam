package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/00kristian/Exam/proto"
	"google.golang.org/grpc"
)


var port int32
var client proto.HashtableClient

func main() {
	port = 8500
	io := bufio.NewReader(os.Stdin)
	welcome()

	
	conn, err := grpc.Dial(toAddr(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}


	defer conn.Close()
	client = proto.NewHashtableClient(conn)

	for {
		text, _ := io.ReadString('\n')
		commands := strings.Fields(text)
		switch commands[0] {
		case "\\put":
			key, _ := strconv.ParseInt(commands[1], 10, 64)
			value, _ := strconv.ParseInt(commands[2], 10, 64)
			put(key, value)
		case "\\help":
			help()
		case "\\result":
			amount, _ := strconv.ParseInt(commands[1], 10, 64)
			get(amount)
		default:
			fmt.Println("Please use one of the commands. If you are unsure about them, type \\help.")
		}
		



	}

	



}


func put(key int64, value int64){
	offer := &proto.Keyvalue {
		Key: key,
		Value: value,
	}

	ack, err := client.Put(context.Background(), offer)

	if err != nil {
		log.Fatalf("Failed to connect to nodes please try again: %v", err)
		if port == 8300 {
			port = 8500
			
		} else {
			port = port - 100
			
		}
	}
	
	if ack.Result {
		fmt.Print("You have succesfully changed a key  ")
		log.Printf("You have succesfully changed a key  ")
	} else {
		fmt.Print("could not connect to server pls wait ")
		if port == 8300 {
			port = 8500
		} else {
			port = port - 100
		}
	

	}




}

func get(key int64){
	offer := proto.GetValue{Key: key}

	ack,err := client.Get(context.Background(), &offer) 

	if err != nil {
		log.Fatalf("Failed to connect to nodes please try again: %v", err)
		if port == 8300 {
			port = 8500
			return
		} else {
			port = port - 100
			return
		}
	}
	if ack.Key != 0{
		fmt.Printf("Here is your current value for the specific key %d", ack.Key)

	}else {
		fmt.Print("No value for that specific key please insert one or try a different key ")
	}


}





func welcome() {
	fmt.Println("____________________ Hashtable servide ____________________")
	fmt.Println("											")
	fmt.Println("Here you have the possiblity to put in different ")
	fmt.Println("a hashmap. To aqquire further assist on ")
	fmt.Println("the tool cosider using the command \\help.	")
	fmt.Println("_______________________________________________________")
	fmt.Println()
}

func help() {
	fmt.Println("____________________ Commands _________________________")
	fmt.Println("											")
	fmt.Println("Following commands are available:		")
	fmt.Println("\\put <Key> <Value>  ")
	fmt.Println("\\get <Key>             ")
	fmt.Println("_______________________________________________________")
	fmt.Println()
}

func toAddr(port int32) string {
	return fmt.Sprintf("localhost:%v", port)
}


