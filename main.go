package main

import (
	"github.com/henrylee2cn/faygo"
	"project/go-resetfulDocker/db"
	"encoding/json"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
	//"go-resetfulDocker/docker"
	"project/go-resetfulDocker/docker"
)



//GET
//type testGet struct {
//	//   <range: 0:10> 限制数据大小[0,10)
//	Id        string      `param:"<in:query> <required>"`// <desc:ID>
//	//Title     string   `param:"<in:query> <nonzero>"`
//	//Paragraph []string `param:"<in:query> <name:p> <len: 1:10> <regexp: ^[\\w]*$>"`
//	//Cookie    string   `param:"<in:cookie> <name:faygoID>"`
//	// Picture         *multipart.FileHeader `param:"<in:formData> <name:pic> <maxmb:30>"`
//}
//
////PUT
//type testPut struct {
//	Id string `param:"<in:query> <required>"`
//}
////Delete
//type testDelete struct {
//	Id string `param:"<in:query> <required>"` // <desc:ID>
//}
//
////getInfo
//type getInfo struct {
//	Key string `param:"<in:query> <required>"` // <desc:ID>
//}

//测试页
//type testPage struct{}
//type test struct{}


////LevelDb api
//POST
//type SaveProblems struct {
//	Data map[string]interface{} 	`param:"<in:body>"`
//}

////docker api
//Image List
type ImageList struct {}

//Get
// Image one
type ImageOne struct{
	ImageId string  `param:"<in:query> <required>"`
}

//Get
//build image
type BuildImage struct{
	Data map[string]interface{} 	`param:"<in:body>"`
}

//Get
//build image info
type BII struct{
	ImgID	string	`param:"<in:query> <required>"`
}

//Get
//Image to tar
type ImageSave struct{}

//Get
//build container
type CreateContainer struct{
	Name  string `param:"<in:query> <required>"`
}

//Get
//build container info
type CCI struct{}

//Get
//run container
type StartContainer struct{
	CID		string `param:"<in:query> <required>"`
}

//Get
//container one
type ContainerOne struct{
	Id string `param:"<in:query> <required>"`
}

//Get
//container list
type ContainerList struct{}





//GET实现
//func (i *testGet) Serve(ctx *faygo.Context) error {
//
//	//test ImageList
//	//docker.GetAll()
//
//	result := db.Get("src/user1/dockerSrc/DB", i.Id)
//	return ctx.String(200, result)
//	//return ctx.JSON(200, i)
//}
//POST实现
//func(p *SaveProblems) Serve(ctx *faygo.Context) error{

	//将数据存储进levelDB
	//对收到的数据进行解读
	//for problemId, problem := range p.Data{
	//	value, err1 :=problem.(map[string]interface{})
	//	if !err1 {
	//		return ctx.String(200, "数据格式有误")
	//	}
	//	mjson,_ :=json.Marshal(value)
	//	mString :=string(mjson)
	//
	//	fmt.Println(p.Data)
	//	result, err := db.Save("src/user1/dockerSrc/DB", problemId, mString)
	//
	//	if err != nil{
	//		return ctx.String(200,  problemId + result + "：" + err.Error())
	//	}
	//}

	//
	//str, err := docker.CreateImage()

	//编写Dockerfile
	//file.CreateDockerFile("src/user1/dockerSrc/Dockerfile")
	//fmt.Println("Dockerfile end")
	//
	////生成tar包
	//file.CreateTarfile("src/user1/srcTar/temp.tar", "src/user1/dockerSrc")
	//fmt.Println("tar end")
	//
	////imagesbuild
	//docker.BuildImage("ttt:ttt", "src/user1/srcTar/temp.tar")
	//fmt.Println("imagesbuild end")
	//
	////将镜像保存成tar文件
	////docker.SaveImage(name, path)
	//err := docker.SaveImage("ttt:ttt", "src/imageTar/temp.tar")
	//if err != nil{
	//	return ctx.String(200, err.Error())
	//}
	//fmt.Println("SaveImage end")

	//返回可供下载dockeriamge.tar的地址
//	return ctx.String(200, "localhost:8080/imageTar/temp.tar")
//}
//PUT实现
//func(p *testPut) Serve(ctx *faygo.Context) error{
//	//if ctx.CookieParam("faygoID") == "" {
//	//	ctx.SetCookie("faygoID", time.Now().String())
//	//}
//	//for problemId, value := range p.Data {
//	//	setValue, err1 := value.(map[string]interface{})
//	//	if !err1 {
//	//		return ctx.String(200, "数据格式有误")
//	//	}
//	//	mjson,_ :=json.Marshal(setValue)
//	//	mString :=string(mjson)
//	//	err := db.Put("/db/database", problemId, mString)
//	//	if !err {
//	//		return ctx.String(200, problemId+"更改错误")
//	//	}
//	//}
//	result := db.Put("/db/database", p.Id )
//
//	return ctx.String(200,  strconv.FormatBool(result) )
//	//return ctx.JSON(200,p)
//}
////DELETE实现
//func (i *testDelete) Serve(ctx *faygo.Context) error {
//	//if ctx.CookieParam("faygoID") == "" {
//	//	ctx.SetCookie("faygoID", time.Now().String())
//	//}
//	result := db.Delete("/db/database", i.Id)
//	return ctx.String(200, strconv.FormatBool(result))
//	//return ctx.JSON(200, i)
//}
//
////test轮询
//func (r *getInfo) Serve(ctx *faygo.Context) error {
//
//
//	re := db.Get("src/statusDB","ttt3")
//	fmt.Println(re)
//	//return ctx.JSON(200, r.Json, true)
//	return ctx.String(200, re)
//}
////testPage
//func (r *testPage) Serve(ctx *faygo.Context) error {
//	return ctx.Render(200, "view/index.html", faygo.Map{})
//}
//
//
////test
//func (r *test) Serve(ctx *faygo.Context) error {
//
//	go docker.CreateImage("user1", "ttt2", "ttt3")
//	return ctx.String(200, "test")
//}



////api Serve

//ImageList 7
func (IL *ImageList) Serve(ctx *faygo.Context)error  {

	//get all Image info
	imageList, err := docker.ImageList()
	if err != nil{
		panic(err)
	}
	//fmt.Println(imageList)
	//b,err := json.Marshal(imageList)
	//fmt.Println(string(b))
	return ctx.JSON(200, imageList)
	//return ctx.String(200, "ImageList test")
}

//Image One 7
func (BI *ImageOne) Serve(ctx *faygo.Context) error{

	//id := "sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869"
	imageOne, err := docker.ImageOne(BI.ImageId)
	//return ctx.String(200, "ImageOne")
	if err != nil{
		panic(err)
	}
	return ctx.JSON(200, imageOne)
}

//Post
//2 BuildImage 7   next =>>file path
func (BI *BuildImage) Serve(ctx *faygo.Context) error{
	id := BI.Data["uid"]
	imgName := BI.Data["imgName"]
	imgId := BI.Data["imgId"]
	Data := BI.Data["problem"]
	//
	imageName, err1 := imgName.(string)
	if !err1 {
		return ctx.String(200, "imgName数据格式有误")
	}
	imageId, err1 := imgId.(string)
	if !err1 {
		return ctx.String(200, "imgId数据格式有误")
	}
	uid, err1 := id.(string)
	if !err1 {
		return ctx.String(200, "uid数据格式有误")
	}


	fmt.Println(imageName,imageId)
	problemData, err := Data.(map[string]interface{})
	if !err{
		return ctx.String(200, "数据格式有误")
	}
	for problemId, problem := range problemData{
		value, err1 :=problem.(map[string]interface{})
		if !err1 {
			return ctx.String(200, "数据格式有误")
		}
		mjson,_ :=json.Marshal(value)
		mString :=string(mjson)

		//fmt.Println(BI.Data)


		result, err := db.Save(`src/` + uid +`/dockerSrc/DB`, problemId, mString)

		if err != nil{
			return ctx.String(200,  problemId + result + err.Error())
		}
	}
	go docker.CreateImage(uid, imageName, imageId)
	return ctx.String(200, "BuildImage")
}

//3 BII 7 notreal
func (BI *BII) Serve(ctx *faygo.Context) error{
	re := db.Get("src/statusDB",BI.ImgID)
	fmt.Println(re)
	//return ctx.JSON(200, r.Json, true)
	return ctx.String(200, re)
}

//4 ImageSave   no
func (BI *ImageSave) Serve(ctx *faygo.Context) error{

	//go docker.CreateImage("user1", "ttt2", "ttt3")
	return ctx.String(200, "ImageSave")
}

//5 CreateContainer 7
func (BI *CreateContainer) Serve(ctx *faygo.Context) error{
	docker.CreateContainer(BI.Name)
	return ctx.String(200, "buildContainer")
}

//6 CCI
func (BI *CCI) Serve(ctx *faygo.Context) error{
	return ctx.String(200, "CCI")
}

//7 StartContainer
func (BI *StartContainer) Serve(ctx *faygo.Context) error{
	err := docker.StartContainer(BI.CID)
	if err != nil{
		return ctx.String(200, err.Error())
	}
	return ctx.String(200, "OK")
}

//8 Containerone 7
func (BI *ContainerOne) Serve(ctx *faygo.Context) error{
	//id := "63b763ba1531ab210daa275242bf839ddd9517d3650a92b1321dbec4a6148fd5"
	COne, err := docker.ContainerOne(BI.Id)
	if err != nil{
		fmt.Println(err)
	}

	return ctx.JSON(200, COne)
}

//9 ContainerList 7  == docker ps -a
func (BI *ContainerList) Serve(ctx *faygo.Context) error{
	CList, err := docker.ContainerList()
	if err != nil{
		panic(err)
	}
	fmt.Println(CList)
	return ctx.JSON(200, CList)
}









// 补充API文档信息
//Get
//func (a *testGet) Doc() faygo.Doc {
//	return faygo.Doc{
//		// API接口说明
//		Note: "testGet",
//		// 响应说明或示例
//		Return: "读取请求，返回文本格式的处理结果",
//		// 补充参数定义
//		//Params: []faygo.ParamInfo{
//		//	{
//		//		Name:     "other",
//		//		In:       "query",
//		//		Required: true,
//		//		Model:    int(-1),
//		//		Desc:     "other plus number",
//		//	},
//		//},
//	}
//}
////Post
//func (a *BuildImage) Doc() faygo.Doc {
//	str := "添加操作，数据格式为json格式，返回文本格式的处理结果<br>"
//	//+
//	//	`数据格式：
//	//    {1:{<br>`+
//	//	`"name":"第一题",<br>`+
//	//	`"type":"选择题",<br>`+
//	//	`"problem":"1+1=？",<br>`+
//	//	`"answer":"A.1&B.2&C.3&D.4@A",<br>`+
//	//	`"createTime":"2018-7-6 16:34",<br>`+
//	//	`"author":"slq"}}`
//	// {1:{
//	//	     "name":"第一题",
//	//	     "type":"选择题",
//	//	     "problem":"1+1=？",
//	//	     "answer":"A.1&B.2&C.3&D.4@A",
//	//	     "createTime":"2018-7-6 16:34",
//	//	     "author":"slq"
//	//    }
//	// }
//
//	return faygo.Doc{
//		// API接口说明
//		Note: "testPost",
//		// 响应说明或示例
//		Return: str,
//	}
//}
////Put
//func (a *testPut) Doc() faygo.Doc {
//	return faygo.Doc{
//		// API接口说明
//		Note: "testPut",
//		// 响应说明或示例
//		Return: "返回文本格式的处理结果",
//	}
//}
////Delete
//func (a *testDelete) Doc() faygo.Doc {
//	return faygo.Doc{
//		// API接口说明
//		Note: "testDelete",
//		// 响应说明或示例
//		Return: "返回文本格式的处理结果",
//	}
//
//}

//Image One
func (IO *ImageOne) Doc()faygo.Doc{
	return faygo.Doc{
		Note:"ImageId for test ",
		Return:"sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869",
	}
}
//Container one
func (IO *ContainerOne) Doc() faygo.Doc{
	return faygo.Doc{
		Note:"Container ID",
		Return:"63b763ba1531ab210daa275242bf839ddd9517d3650a92b1321dbec4a6148fd5",
	}
}

//cpu 100%
func tet(){
	for {
		tt :=0
		tt++
	}
}

func main() {
	//go tet()
	//go tet()
	//go tet()
	//go tet()

	app := faygo.New("myapp", "0.1")

	// Register the route in a tree style
	//app.Route(
	//	//静态路由
	//	app.NewStatic("imageTar", "src/imageTar"),
	//	app.NewStatic("Plugin", "view/Plugin"),
	//
	//)


	// Register the route in a chain style
	////static file route
	app.Static("src", "src")
	app.Static("Plugin", "view/Plugin")

	//api route
	app.GET("/ImageList", new(ImageList))   //
	app.GET("/ImageOne", new(ImageOne))     //
	app.POST("/BuildImage", new(BuildImage))
	app.GET("/BII", new(BII))
	app.GET("/ImageSave", new(ImageSave))
	app.GET("/CreateContainer", new(CreateContainer))
	app.GET("/BCI", new(CCI))
	app.GET("/ContainerStart", new(StartContainer))
	app.GET("/ContainerOne", new(ContainerOne))
	app.GET("/ContainerList", new(ContainerList))


	//app.GET("/index", new(testGet))
	//app.POST("/index", new(testPost))
	//app.DELETE("/index", new(testDelete))
	//app.PUT("/index", new(testPut))


	//app.GET("/getInfo", new(getInfo))
	//app.GET("/testPage", new(testPage))

	//app.GET("/test", new(test))




	// Start the service
	faygo.Run()

}
