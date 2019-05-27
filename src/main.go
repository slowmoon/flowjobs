package main

import (
	"encoding/binary"
	"fmt"
	"game-server/jobs"
)

func main(){
   fmt.Println("run...")
   person:= jobs.NewPerson("hello")
   person.Run()

   buf := make([]byte, 20)
   n:= binary.PutVarint(buf, 129)
   n += binary.PutVarint(buf[n:],  127)
   fmt.Println(n, buf)
   r,t :=  binary.Varint(buf[:n])
   fmt.Println(r, t)
   r,t =  binary.Varint(buf[t:n])
   fmt.Println(r, t)
}