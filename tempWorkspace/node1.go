// package main

// import (
// 	"fmt"
// 	"net"
// 	// "os"
// 	"bufio"
// 	"time"
// 	"./pow"
// )

// var listOfNodes=[...]int{8081,8082}
// var myPort int=8081
// var ip string= pow.GetOutboundIP().String()
// var publicKey string="32132132132132132113132123"

// func main() {
// 	nonce, result:=pow.Pow(ip,publicKey,0)
// 	fmt.Println("Hello World")
// 	fmt.Println(nonce)
// 	fmt.Println(result)

// 	communicatePow()


// }


// func communicatePow() bool {
//    ln, _ := net.Listen("tcp", ":"+string(myPort))

//   // accept connection on port
//   conn, _ := ln.Accept()

//   var startTime = time.Now()
  
//   // run loop forever (or until ctrl-c)
//   for time.Since(startTime)!=600 { 
//     // will listen for message to process ending in newline (\n)
//     powIdentity, _ := bufio.NewReader(conn).ReadString('\n')
//     // output message received
//     fmt.Println("Message Received:", string(powIdentity))
//     // sample process for string received
    
    
//   }

//   for i := 0; i < len(listOfNodes); i++ {
//   		fmt.Println("Current Node: "+string(listOfNodes[i]))
// 		if listOfNodes[i]!=myPort {
// 			conn, _ := net.Dial("tcp", "127.0.0.1:"+string(listOfNodes[i]))
// 			reader := bufio.NewReader(os.Stdin)
// 			fmt.Print("Text to send: ")
//     		text, _ := reader.ReadString('\n')
//     		// send to socket
//     		fmt.Fprintf(conn, text + "\n")
// 		}
// 	}

// }
package main

import (
  "fmt"
  "net"
  // /"os"
  "bufio"
  "strconv"
  "time"
  "./pow"
  "strings"
)

var listOfNodes=[...]int{8081,8082,8083} //list of nodes in network
var myPort int=8081 //my address
var ip string= pow.GetOutboundIP().String()
var publicKey string="32132132132132132113132123"
var identity="0x00" 
var m map[int]string
var sentMap map[int]bool



func main() {
  _, result:=pow.Pow(ip,publicKey,0)
  fmt.Println()
  fmt.Println()
  fmt.Println()
  fmt.Println("POW done now lets communicate it")
  m=make(map[int]string)
  m[myPort]=result
  sentMap=make(map[int]bool)
  sentMap[myPort]=true
  fmt.Println(m[myPort])
  communicatePow()
}


func communicatePow()  {
  fmt.Println("Starting Communication")
  fmt.Println(strconv.Itoa(myPort))



//   go func(){
//   	fmt.Println("In go func")
//   		ln, _ := net.Listen("tcp", ":"+strconv.Itoa(myPort))
//   // fmt.Println(ln)
//   // fmt.Println("Hello 1")
//   // // accept connection on port
//   conn, _ := ln.Accept()
//   // fmt.Println("Hello 2")
//   var startTime = time.Now()
//   // fmt.Println("Starting the for loop")
  
//   for time.Since(startTime)<=6000000 { 
//     // will listen for message to process ending in newline (\n)
//     // fmt.Println("Time since start now is")
//     // fmt.Println(time.Since(startTime))

//     data, _ := bufio.NewReader(conn).ReadString('\n')
//     if data!="" {
//     	fmt.Println("Message Received:" + data)
    	
//     	splitted:=strings.Split(strings.TrimSpace(data)," ")
    	
//     	powIdentity:= splitted[0]
//     	sender:= splitted[1]
//     	fmt.Println("POW:"+powIdentity)
//     	fmt.Println(sender)
//     	senderPort,err:=strconv.Atoi(sender)
//     	if err!=nil{

//     	}
//     	m[senderPort]=powIdentity
//     }
//     // output message received
    
//     // sample process for string received
//   }
// }()

  // go func(){
  // 		for i := 0; i < len(listOfNodes); i++ {
  //     fmt.Println("Current Node: "+strconv.Itoa(listOfNodes[i]))
  //   if listOfNodes[i]!=myPort {
  //     conn, _ := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(listOfNodes[i]))
  //     //reader := bufio.NewReader(os.Stdin)
        
  //       // send to socket
  //       fmt.Fprintf(conn, m[myPort] + "\n")
  //   }
  // }

  // 	}()

	go send()
	go receive()
  	
   time.Sleep(60 * time.Second)
   fmt.Println(m)
}

func send(){

	for len(m)!=len(listOfNodes){
		fmt.Println(len(m))
		fmt.Println("Sending")

		for i := 0; i < len(listOfNodes); i++ {
      		if sentMap[listOfNodes[i]]!=true{
      				fmt.Println("Current Node: "+strconv.Itoa(listOfNodes[i]))
    				if listOfNodes[i]!=myPort {
      				conn, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(listOfNodes[i]))
      				//reader := bufio.NewReader(os.Stdin)
        			if err!=nil && sentMap[listOfNodes[i]]!=true{
        	
        			}else{
        				sentMap[listOfNodes[i]]=true
        				fmt.Fprintf(conn, m[myPort]+" "+strconv.Itoa(myPort) + "\n")	
        			}
    				}
      		}
      	}
  		time.Sleep(10*time.Second)
  		receive()	
	}
}


func receive() {
	

	fmt.Println("Receiving")
  		ln, _ := net.Listen("tcp", ":"+strconv.Itoa(myPort))
  // fmt.Println(ln)
  // fmt.Println("Hello 1")
  // // accept connection on port
  conn, err := ln.Accept()
  if err!=nil{
  	time.Sleep(10*time.Second)
  	send()
  }
  // fmt.Println("Hello 2")
  var startTime = time.Now()
  // fmt.Println("Starting the for loop")
  
  for time.Since(startTime)<=6000000 { 
    // will listen for message to process ending in newline (\n)
    // fmt.Println("Time since start now is")
    // fmt.Println(time.Since(startTime))
  	fmt.Println(time.Since(startTime))
    data, _ := bufio.NewReader(conn).ReadString('\n')
    if data!="" {
    	fmt.Println("Message Received:" + data)
    	
    	splitted:=strings.Split(strings.TrimSpace(data)," ")
    	
    	powIdentity:= splitted[0]
    	sender:= splitted[1]
    	fmt.Println("POW:"+powIdentity)
    	fmt.Println(sender)
    	senderPort,err:=strconv.Atoi(sender)
    	if err!=nil{

    	}
    	
    	m[senderPort]=powIdentity
    	fmt.Println(m)
    }
    // output message received
    
    // sample process for string received
  }
	
}