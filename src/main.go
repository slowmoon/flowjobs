package main

import (
	"fmt"
	"game-server/jobs"
)

func main(){
   fmt.Println("run...")   
   person:= jobs.NewPerson("hello")
   person.Run()
}