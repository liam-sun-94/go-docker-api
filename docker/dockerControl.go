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

/**
* 建立docker镜像
* 参数 name  新建镜像的名称  格式： xxx:xx  <REPOSITORY>:<TAG>
* 返回值  error 类型
* 返回 新建错误
 */

func BuildImage(name string, tarPath string) error{
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
		return err
	}
	options := types.ImageBuildOptions{
		SuppressOutput: true,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Tags:           []string{name},
	}

	BuildContext, err := os.Open(tarPath)
	fmt.Println(tarPath)

	buildResponse, err := cli.ImageBuild(ctx, BuildContext, options)

	if err != nil {
		panic(err)
		return err
	}

	fmt.Println(buildResponse)
	fmt.Println(buildResponse.OSType)
	fmt.Println(buildResponse.Body)
	////buildResponse.Body
	//fmt.Println(buildResponse.Body)

	defer buildResponse.Body.Close()
	defer BuildContext.Close()

	return nil
}


/**
* 获取所有docker镜像
* 参数
* 返回值
* 测试使用
 */
func GetAll(){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	imageList, err := cli.ImageList(ctx, types.ImageListOptions{})
	var Id string
	for  x, v := range imageList{
		if x>0{
			break
		}
		fmt.Println(x)
		Id = v.ID
		fmt.Println(v)
		fmt.Println(v.ID)
		fmt.Println(v.RepoTags[0])
		fmt.Println(v.Containers)

		//fmt.Println(x[0]
	}
	SaveImage(Id, "tttt")
}


/**
* 保存特定镜像
* 参数 name string类型   格式： nginx:latest  <REPOSITORY>:<TAG>
* 	  path  镜像的存储路径
* 返回值 error  返回新建错误
*
 */
func SaveImage(name string, path string) error{
	fmt.Println("save")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil{
		fmt.Println(err)
		return err
	}
	read, err := cli.ImageSave(ctx, []string{name})
	if err != nil{
		fmt.Println(err)
		return err
	}
	//fmt.Println(read)
	res, err := ioutil.ReadAll(read)
	//fmt.Println(string(res))

	file, err := os.Create(path)
	//fmt.Println(path)

	if err != nil {
		fmt.Println(err)
		return err
	}
	file.Write(res)
	defer file.Close()

	return nil

}