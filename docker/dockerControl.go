package docker
//package main

import (
	"github.com/docker/docker/client"
	"fmt"
	"github.com/docker/docker/api/types"
	"os"
	"context"
	"io/ioutil"
	"project/go-resetfulDocker/file"
	"project/go-resetfulDocker/db"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)


//image info
type ImageInfoList struct{
	Code 	int
	Msg		string
	Count   int
	Data	[]ImageListInfo
}
type ImageListInfo struct{
	Containers 	int64
	Created 	int64
	ID 			string
	Labels 		string
	//ParentID 	string
	RepoDigests []string
	RepoTags 	string
	Size 		int64
}

func getCtx()(context.Context, *client.Client, error){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
		return nil, nil, err
	}
	return ctx, cli, nil
}

/**
* 建立docker镜像
* 参数 name  新建镜像的名称  格式： xxx:xx  <REPOSITORY>:<TAG>
* 返回值  error 类型
* 返回 新建错误
 */

func BuildImage1(name string, tarPath string) error{
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
	buildResponse, err := cli.ImageBuild(ctx, BuildContext, options)
	if err != nil {
		panic(err)
		return err
	}
	body, err := ioutil.ReadAll(buildResponse.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	defer buildResponse.Body.Close()
	defer BuildContext.Close()
	return nil
}


//Image List
//no args
//return all Images info
func ImageList()(ImageInfoList, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
		return ImageInfoList{}, err
	}
	imageList, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
		return ImageInfoList{}, err
	}

	var IIL ImageInfoList
	fmt.Println(imageList)
	//var result InageInfoList
	IIL.Code = 0
	IIL.Count = len(imageList)
	IIL.Msg = ""
	for _, v := range imageList{
		var ILI ImageListInfo
		ILI.Containers = v.Containers
		ILI.Created = v.Created
		ILI.ID = v.ID
		ILI.Labels = v.Labels["maintainer"]
		//ILI.ParentID = v.ParentID
		ILI.RepoDigests = v.RepoDigests
		ILI.RepoTags = v.RepoTags[0]
		ILI.Size = v.Size
		IIL.Data = append(IIL.Data, ILI)
	}
	fmt.Println(IIL.Data)
	return IIL, nil
}


//Image one
//return a image info
func ImageOne(imageId string )(types.ImageInspect, error){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
		return types.ImageInspect{}, err
	}
	imageOne, _, err := cli.ImageInspectWithRaw(ctx, imageId)
	if err != nil {
		panic(err)
		return types.ImageInspect{}, err
	}
	//fmt.Println(imageList)
	//fmt.Println(t)
	//return imageList, nil
	return imageOne, nil
}


//container list  == docker ps -a
//return all container info
func ContainerList()([]types.Container, error){

	//ctx := context.Background()
	//cli, err := client.NewClientWithOpts()
	//if err != nil {
	//	panic(err)
	//	return nil, err
	//}
	ctx, cli, err := getCtx()
	containerList, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All:	true,  //docker ps -a
	})
	if err != nil {
		panic(err)
		return nil, err
	}
	fmt.Println(containerList)
	for _, v := range containerList{
		fmt.Println(v)
	}
	return containerList, nil
}


//container one
//return a container inf0
func ContainerOne(cID string)(types.ContainerJSON, error){
	ctx, cli, err := getCtx()
	if err != nil {
		panic(err)
		return types.ContainerJSON{}, err
	}
	containerList, err := cli.ContainerInspect(ctx, cID)
	if err != nil {
		panic(err)
		return types.ContainerJSON{}, err
	}

	return containerList, nil

}

//Start Container
//no return
func StartContainer(ContainerID string) error{
	ctx, cli, err := getCtx()
	if err != nil {
		panic(err)
		//return types.ContainerJSON{}, err
		return err
	}
	if err := cli.ContainerStart(ctx, ContainerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
		return err
	}
	return nil
}

//Create Container
//no return
func CreateContainer(name string){
	ctx, cli, err := getCtx()
	if err != nil {
		panic(err)
		//return types.ContainerJSON{}, err
	}
	//config := container.Config{
	//	Image: "",
	//}
	config := container.Config{
		Image: "alpine",
		//Cmd:   []string{"/bin/sh", "-c", "echo test"},
		AttachStdin: true,
		OpenStdin  : true,
		ExposedPorts: nat.PortSet{
			"8080/tcp":  {},
			"3306/tcp":  {},
		},
	}
	hostConfig := container.HostConfig{
		AutoRemove: false,
		PortBindings: nat.PortMap{
			"3306/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "3306"},
			},
		},
	}

	networkingConfig := network.NetworkingConfig{}
	resp, err := cli.ContainerCreate(ctx, &config, &hostConfig, &networkingConfig, name)
	if err != nil {
		panic(err)
		fmt.Println(err)
		//return types.ContainerJSON{}, err
	}
	fmt.Println(1)
	//it can run the contianer if no ContainerStart() the container status always is created
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Println(2)

	//if _, err := cli.ContainerWait(ctx, resp.ID,""); err != nil {
	//	panic(err)
	//}
	fmt.Println(resp)
	//out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(5)
	//io.Copy(os.Stdout, out)

	//ctx := context.Background()
	//cli, err := client.NewClientWithOpts()
	//if err != nil {
	//	panic(err)
	//}
	//
	////_, err = cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	////if err != nil {
	////	panic(err)
	////}
	//fmt.Println(1)
	//resp, err := cli.ContainerCreate(ctx, &container.Config{
	//	Image: "alpine",
	//	Cmd:   []string{"echo", "hello world"},
	//}, nil, nil, name)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(2)
	//if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
	//	panic(err)
	//}
	//fmt.Println(3)
	//if _, err1 := cli.ContainerWait(ctx, resp.ID, ""); err1 != nil {
	//	panic(err1)
	//}
	//fmt.Println(4)

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



//Create Image
//args uid    imgname    imageId
//no return
func CreateImage(uid string, imgName string, imgId string){

	src := "src"
	xLine := "/"
	dockerFile := "dockerSrc/Dockerfile"
	tarFile := "srcTar/tt.tar"
	dockerSrc := "dockerSrc"
	tag := ":dockertest"
	imgPath := "imageTar"
	imgTar := ".tar"

	//fmt.Println(time.Now().Unix())

	//编写Dockerfile
	//file.CreateDockerFile("src/user2/dockerSrc/Dockerfile", "temp")
	//fmt.Println("Dockerfile end")
	file.CreateDockerFile(src + xLine + uid + xLine + dockerFile)
	//fmt.Println(src + xLine + uid + xLine + dockerFile)
	result, err := db.Save("src/statusDB", imgId, "Dockerfile end ")
	fmt.Println("Dockerfile end :    " + result)


	//生成tar包
	//file.CreateTarfile("src/user2/srcTar/temp.tar", "src/user2/dockerSrc")
	//fmt.Println("tar end")
	file.CreateTarfile(src + xLine + uid + xLine + tarFile, src + xLine + uid + xLine +dockerSrc)
	//fmt.Println(src + xLine + uid + xLine + tarFile)
	//fmt.Println(src + xLine + uid + xLine +dockerSrc)
	result, err = db.Save("src/statusDB", imgId, "tar end")
	fmt.Println("tar end")

	//imagesbuild
	//BuildImage("ttt:ttt", "src/user2/srcTar/temp.tar")
	//fmt.Println("imagesbuild end")
	BuildImage1(imgName + tag, src + xLine + uid + xLine + tarFile)
	//fmt.Println(imgName + tag)
	//fmt.Println(src + xLine + uid + xLine + tarFile)
	result, err = db.Save("src/statusDB", imgId, "imagesbuild end")
	fmt.Println("imagesbuild end")

	//将镜像保存成tar文件
	//docker.SaveImage(name, path)
	//err := SaveImage("ttt:ttt", "src/imageTar/temp.tar")
	err = SaveImage(imgName + tag, src + xLine + imgPath + xLine + imgId + imgTar)
	if err != nil{
		result, err = db.Save("src/statusDB", imgId, "SaveImage end fail" + err.Error())
		fmt.Println("imagesbuild end" + err.Error())
	}
	//fmt.Println(imgName + tag)
	//fmt.Println(src + xLine + imgPath + xLine + imgId + imgTar)
	result, err = db.Save("src/statusDB", imgId, "SaveImage end")
	fmt.Println("SaveImage end")

	pathd := `localhost:8080/` + src + xLine + imgPath + xLine + imgId + imgTar
	result, err = db.Save("src/statusDB", imgId, pathd)
	fmt.Println(pathd)

	//fmt.Println(time.Now().Unix())
}

