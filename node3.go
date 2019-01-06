// package main

// import (
//   "fmt"
//   "net"
//   // "os"
//   "bufio"
//   "time"
//   "./pow"
// )

// var listOfNodes=[...]int{8081,8082}
// var myPort int=8082
// var ip string= pow.GetOutboundIP().String()
// var publicKey string="32132132132132132113132123"
// var identity="0x00" 
// var m map[int]string 

// func main() {
//   nonce, result:=pow.Pow(ip,publicKey,0)
//   fmt.Println("Hello World")
//   fmt.Println(nonce)
//   fmt.Println(result)
//   m[myPort]=result
//   communicatePow()
// }


// func communicatePow()  {
//    ln, _ := net.Listen("tcp", ":"+string(myPort))

//   // accept connection on port
//   conn, _ := ln.Accept()

//   var startTime = time.Now()
  
  
//   for time.Since(startTime)!=60000 { 
//     // will listen for message to process ending in newline (\n)
//     powIdentity, _ := bufio.NewReader(conn).ReadString('\n')
//     // output message received
//     fmt.Println("Message Received:", string(powIdentity))
//     // sample process for string received
       
//   }

//   for i := 0; i < len(listOfNodes); i++ {
//       fmt.Println("Current Node: "+string(listOfNodes[i]))
//     if listOfNodes[i]!=myPort {
//       conn, _ := net.Dial("tcp", "127.0.0.1:"+string(listOfNodes[i]))
//       reader := bufio.NewReader(os.Stdin)
        
//         // send to socket
//         fmt.Fprintf(conn, identity + "\n")
//     }
//   }

// }
package main

import (
  "fmt"
  "net"
  // /"os"
   "bufio"
  "strconv"
   "time"
   "strings"
  "./pow"
)

var listOfNodes=[...]int{8081,8082,8083} //list of nodes in network
var myPort int=8083 //my address
var ip string= pow.GetOutboundIP().String()
var publicKey string="5213213213213221298"
var identity="0x00" 
var m map[int]string
var sentMap map[int]bool
var numNodes int=len(listOfNodes)
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
  
  go receive()
  go send()

  

  // go func2(){
  //     for i := 0; i < len(listOfNodes); i++ {
  //     fmt.Println("Current Node: "+strconv.Itoa(listOfNodes[i]))
  //       if listOfNodes[i]!=myPort {
  //         conn, _ := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(listOfNodes[i]))
  //         //reader := bufio.NewReader(os.Stdin)
        
  //         // send to socket
  //         fmt.Fprintf(conn, m[myPort]+" "+strconv.Itoa(myPort)+ "\n")
  //       }
  //     }
  // }()
  time.Sleep(10*time.Second)
  fmt.Println(m)
}


func send() {

  for len(m)!=numNodes{

  for i := 0; i < numNodes; i++ {
      fmt.Println("Current Node: "+strconv.Itoa(listOfNodes[i]))
        if listOfNodes[i]!=myPort {

          if sentMap[listOfNodes[i]]!=true {
                conn, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(listOfNodes[i]))
                //reader := bufio.NewReader(os.Stdin)
                if err!=nil || sentMap[listOfNodes[i]]!=true{
                  fmt.Println("Already Sent to node "+ strconv.Itoa(listOfNodes[i]))
                  fmt.Println("Error")
          }else{
                sentMap[listOfNodes[i]]=true
                fmt.Fprintf(conn, m[myPort]+" "+strconv.Itoa(myPort) + "\n")  
          }  
          }
        }
      }
  }
}


func receive() {
      fmt.Println(strconv.Itoa(myPort))
  ln, _ := net.Listen("tcp", ":"+strconv.Itoa(myPort))
  //fmt.Println(ln)
  //fmt.Println("Hello 1")
  // accept connection on port
  conn, _ := ln.Accept()
  //fmt.Println("Hello 2")
  var startTime = time.Now()
  fmt.Println("Starting the for loop")
  
  for time.Since(startTime)<=6000000 { 
    // will listen for message to process ending in newline (\n)
    // fmt.Println("Time since start now is")
    // fmt.Println(time.Since(startTime))

    // powIdentity, _ := bufio.NewReader(conn).ReadString('\n')
    // // output message received
    // fmt.Println("Message Received:", string(powIdentity))
    // // sample process for string received
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
    }
  }
  
}