package file

import (
	"fmt"
	"os"
	"project/go-resetfulDocker/cmd"
)
/**
*  	临时的Dockerfile文件方法
*
*/
func CreateDockerFile(path string){

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//str1 := "FROM alpine \n" +
	//	"COPY sgerrand.rsa.pub /etc/apk/keys/sgerrand.rsa.pub \n" +
	//	"COPY glibc-2.27-r0.apk /  \n" +
	//	"RUN  apk add glibc-2.27-r0.apk \n" +
	//	"COPY DB /usr/local/go/go-contianerAPI/DB  \n" +
	//	"COPY go-containerAPI /usr/local/go/go-contianerAPI/go-containerAPI   \n" +
	//	`CMD ["/bin/sh", "-c", "/usr/local/go/go-contianerAPI/go-containerAPI"]`
	str2 := "FROM registry.cn-hangzhou.aliyuncs.com/alex-docker/alpine-go:v1.0 \n" +
	"COPY DB /usr/local/go/go-contianerAPI/DB \n" +
	"COPY go-containerAPI /usr/local/go/go-contianerAPI/go-containerAPI \n" +
	`CMD ["/bin/sh", "-c", "/usr/local/go/go-contianerAPI/go-containerAPI"]`

	//"RUN /usr/local/go/go-contianerAPI/go-containerAPI \n"
	//str1 := "FROM centos \n RUN ifconfig \n CMD \n" + test
	//buffStr := []byte(str1)
	//re, err := f.Write(buffStr)
    //
	f.WriteString(str2)
	if err != nil {
		panic(err)
	}
}

func CreateTarfile(pathT string, pathF string){
	//"test/test1.tar"
	//"/home/alex/go/src/test_cli/go_docker_api"
	fmt.Println(pathT)
	re := cmd.Lcmd("tar", []string{"-cvf", pathT, "-C", pathF,"."})
	fmt.Println(re)
}