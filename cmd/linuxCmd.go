package cmd

import (
"os/exec"
"log"
)

func Lcmd(command string ,args []string) string {
	//args:=[]string{"network","ls"}
	out ,err:= exec.Command(command,args...).Output()
	if err!=nil{
		log.Fatal(err)
	}
	//fmt.Printf("%s",string(out))
	//fmt.Println(string(out))

	return string(out)
}
