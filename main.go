package main

import (
	"github.com/henrylee2cn/faygo"
	"go-resetfulDocker/db"
	"encoding/json"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
	"go-resetfulDocker/docker"
)
//GET
type testGet struct {
	//   <range: 0:10> 限制数据大小[0,10)
	Id        string      `param:"<in:query> <required>"`// <desc:ID>
	//Title     string   `param:"<in:query> <nonzero>"`
	//Paragraph []string `param:"<in:query> <name:p> <len: 1:10> <regexp: ^[\\w]*$>"`
	//Cookie    string   `param:"<in:cookie> <name:faygoID>"`
	// Picture         *multipart.FileHeader `param:"<in:formData> <name:pic> <maxmb:30>"`
}
//POST
type testPost struct {
	Data map[string]interface{} 	`param:"<in:body>"`
}
//PUT
type testPut struct {
	Id string `param:"<in:query> <required>"`
}
//Delete
type testDelete struct {
	Id string `param:"<in:query> <required>"` // <desc:ID>
}


//GET实现
func (i *testGet) Serve(ctx *faygo.Context) error {

	//test ImageList
	docker.GetAll()

	result := db.Get("src/user1/dockerSrc/DB", i.Id)
	return ctx.String(200, result)
	//return ctx.JSON(200, i)
}
//POST实现
func(p *testPost) Serve(ctx *faygo.Context) error{
	//将数据存储进levelDB
	for problemId, problem := range p.Data{
		value, err1 :=problem.(map[string]interface{})
		if !err1 {
			return ctx.String(200, "数据格式有误")
		}
		mjson,_ :=json.Marshal(value)
		mString :=string(mjson)

		fmt.Println(p.Data)
		result, err := db.Save("src/user1/dockerSrc/DB", problemId, mString)

		if err != nil{
			return ctx.String(200,  problemId + result + "：" + err.Error())
		}
	}

	//编写Dockerfile
	db.CreateDockerFile("src/user1/dockerSrc/Dockerfile", "temp")

	//生成tar包
	db.CreateTarfile("src/user1/srcTar/test.tar", "src/user1/dockerSrc")

	//imagesbuild
	docker.BuildImage("test:v1", "src/user1/srcTar/test.tar")

	//将镜像保存成tar文件
	//docker.SaveImage(name, path)
	err := docker.SaveImage("test:v1", "src/imageTar/temp.tar")
	if err != nil{
		return ctx.String(200, err.Error())
	}

	//返回可供下载dockeriamge.tar的地址
	return ctx.String(200, "localhost:8080/imageTar/temp.tar")
}
//PUT实现
func(p *testPut) Serve(ctx *faygo.Context) error{
	//if ctx.CookieParam("faygoID") == "" {
	//	ctx.SetCookie("faygoID", time.Now().String())
	//}
	//for problemId, value := range p.Data {
	//	setValue, err1 := value.(map[string]interface{})
	//	if !err1 {
	//		return ctx.String(200, "数据格式有误")
	//	}
	//	mjson,_ :=json.Marshal(setValue)
	//	mString :=string(mjson)
	//	err := db.Put("/db/database", problemId, mString)
	//	if !err {
	//		return ctx.String(200, problemId+"更改错误")
	//	}
	//}
	result := db.Put("/db/database", p.Id )

	return ctx.String(200,  strconv.FormatBool(result) )
	//return ctx.JSON(200,p)
}
//DELETE实现
func (i *testDelete) Serve(ctx *faygo.Context) error {
	//if ctx.CookieParam("faygoID") == "" {
	//	ctx.SetCookie("faygoID", time.Now().String())
	//}
	result := db.Delete("/db/database", i.Id)
	return ctx.String(200, strconv.FormatBool(result))
	//return ctx.JSON(200, i)
}


// 补充API文档信息

//Get
func (a *testGet) Doc() faygo.Doc {
	return faygo.Doc{
		// API接口说明
		Note: "testGet",
		// 响应说明或示例
		Return: "读取请求，返回文本格式的处理结果",
		// 补充参数定义
		//Params: []faygo.ParamInfo{
		//	{
		//		Name:     "other",
		//		In:       "query",
		//		Required: true,
		//		Model:    int(-1),
		//		Desc:     "other plus number",
		//	},
		//},
	}
}
//Post
func (a *testPost) Doc() faygo.Doc {
	str := "添加操作，数据格式为json格式，返回文本格式的处理结果<br>"
	//+
	//	`数据格式：
	//    {1:{<br>`+
	//	`"name":"第一题",<br>`+
	//	`"type":"选择题",<br>`+
	//	`"problem":"1+1=？",<br>`+
	//	`"answer":"A.1&B.2&C.3&D.4@A",<br>`+
	//	`"createTime":"2018-7-6 16:34",<br>`+
	//	`"author":"slq"}}`
	// {1:{
	//	     "name":"第一题",
	//	     "type":"选择题",
	//	     "problem":"1+1=？",
	//	     "answer":"A.1&B.2&C.3&D.4@A",
	//	     "createTime":"2018-7-6 16:34",
	//	     "author":"slq"
	//    }
	// }
	
	return faygo.Doc{
		// API接口说明
		Note: "testPost",
		// 响应说明或示例
		Return: str,
	}
}
//Put
func (a *testPut) Doc() faygo.Doc {
	return faygo.Doc{
		// API接口说明
		Note: "testPut",
		// 响应说明或示例
		Return: "返回文本格式的处理结果",
	}
}
//Delete
func (a *testDelete) Doc() faygo.Doc {
	return faygo.Doc{
		// API接口说明
		Note: "testDelete",
		// 响应说明或示例
		Return: "返回文本格式的处理结果",
	}

}



func main() {

	//testDb()

	app := faygo.New("myapp", "0.1")

	// Register the route in a chain style
	app.GET("/index", new(testGet))
	app.POST("/index", new(testPost))
	app.DELETE("/index", new(testDelete))
	app.PUT("/index", new(testPut))




	// Register the route in a tree style
	// app.Route(
	//     app.NewGET("/index/:id", new(Index)),
	// )
	//**

	//静态路由
	app.Route(
		app.NewStatic("imageTar", "src/imageTar"),
	)

	//test connect sql server 2008


	// Start the service
	faygo.Run()

}
