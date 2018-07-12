package docker
//package main

import (
	"github.com/docker/docker/client"
	"fmt"
	"github.com/docker/docker/api/types"
	"os"
	"context"
	"io/ioutil"
)

func BuildImage(name string, tarPath string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	options := types.ImageBuildOptions{
		SuppressOutput: true,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Tags:           []string{name},
	}

	BuildContext, err := os.Open(tarPath)
	buildResponse, err := cli.ImageBuild(ctx, BuildContext, options)
	fmt.Println(buildResponse)
	fmt.Println(buildResponse.Body)
	defer buildResponse.Body.Close()
	defer BuildContext.Close()




}


func SaveImage(ImageId string, path string){

	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil{
	fmt.Println(err)
	}
	read, err := cli.ImageSave(ctx, []string{ImageId})
	if err != nil{
	fmt.Println(err)
	}
	fmt.Println(read)
	res, err := ioutil.ReadAll(read)
	fmt.Println(string(res))

	file, err := os.Create("t.tar")
	if err != nil {
	fmt.Println(err)
	}
	file.Write(res)
	defer file.Close()


}