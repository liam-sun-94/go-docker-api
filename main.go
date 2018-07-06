package main

import (
	"github.com/henrylee2cn/faygo"
	"fmt"
	"go-restfulDocker/db"
	"encoding/json"
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
//Delete
type testPut struct {
	Data string	`param:"<in:query>"`
}
//Delete
type testDelete struct {
	Data string	`param:"<in:query>"`
}


//GET实现
func (i *testGet) Serve(ctx *faygo.Context) error {
	//if ctx.CookieParam("faygoID") == "" {
	//	ctx.SetCookie("faygoID", time.Now().String())
	//}
	result := db.Get("/db/database", i.Id)
	return ctx.String(200, result)
	//return ctx.JSON(200, i)
}
//POST实现
func(p *testPost) Serve(ctx *faygo.Context) error{
	//if ctx.CookieParam("faygoID") == "" {
	//	ctx.SetCookie("faygoID", time.Now().String())
	//}
	//t := p.Data["1"]
	//fmt.Println(t["name"])
	//fmt.Println(ctx.JSON(200,p))

	for problemId, problem := range p.Data{
		value, err1 :=problem.(map[string]interface{})
		if !err1 {
			return ctx.String(200, "数据格式有误")
		}
		mjson,_ :=json.Marshal(value)
		mString :=string(mjson)
		result, err := db.Save("/db/database", problemId, mString)

		if err != nil{
			return ctx.String(200,  problemId + result + "：" + err.Error())
		}
	}
	return ctx.String(200, "ok")
}
//PUT实现
func(p *testPut) Serve(ctx *faygo.Context) error{
	//if ctx.CookieParam("faygoID") == "" {
	//	ctx.SetCookie("faygoID", time.Now().String())
	//}
	fmt.Println(p.Data)
	fmt.Println(ctx.JSON(200,p))
	return ctx.JSON(200,p)
}
//DELETE实现
func (i *testDelete) Serve(ctx *faygo.Context) error {
	//if ctx.CookieParam("faygoID") == "" {
	//	ctx.SetCookie("faygoID", time.Now().String())
	//}
	return ctx.String(200, "这是Delete请求！！！" + i.Data)
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

	 //*/


	// Start the service
	faygo.Run()

}
