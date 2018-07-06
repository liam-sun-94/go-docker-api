package main

import (
	"github.com/henrylee2cn/faygo"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"time"
	"strconv"
)
//GET
type testGet struct {
	//   <range: 0:10> 限制数据大小[0,10)
	Id        int      `param:"<in:query> <required>"`// <desc:ID>
	Title     string   `param:"<in:query> <nonzero>"`
	Paragraph []string `param:"<in:query> <name:p> <len: 1:10> <regexp: ^[\\w]*$>"`
	//Cookie    string   `param:"<in:cookie> <name:faygoID>"`
	// Picture         *multipart.FileHeader `param:"<in:formData> <name:pic> <maxmb:30>"`
}
//POST
type testPost struct {
	Data string	`param:"<in:body>"`
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
	return ctx.String(200, "这是Get请求！！！")
	//return ctx.JSON(200, i)
}

//POST实现
func(p *testPost) Serve(ctx *faygo.Context) error{
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

//test goleveldb
func testDb(){
	fmt.Println("test LevelDB")

	db, err := leveldb.OpenFile("./db",nil)
	if err != nil {
		fmt.Println(err)
	}

	//err = db.Put([]byte("testKey3"), []byte("testValue"), nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = db.Put([]byte("testKey4"), []byte("testValue2"), nil)
	//if err != nil {
	//	fmt.Println(err)
	//}

	datetime1 := time.Now().UnixNano() / 1000000
	datetime1s := time.Now().Second()
	for i := 1; i < 200000; i++{
		key := "testKey" + strconv.Itoa(i)
		value := "testValue" + strconv.Itoa(i)
		fmt.Println(key + "       " + value)
		err = db.Put([]byte(key), []byte(value), nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	//err = db.Delete([]byte("testKey"), nil)
	//if err != nil {
	//	fmt.Println(err)
	//}


	//for i := 0; i < 10000; i++ {
	//	temp := rand.Intn(6000000)
	//	key := "testKey" + strconv.Itoa(temp)
	//	data, err := db.Get([]byte(key),nil)
	//	fmt.Println(key + "     " + string(data))
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
	//for i := 0; i < 10; i++ {
	//	temp := 100000 + rand.Intn(100000)
	//	key := "testKey" + strconv.Itoa(temp)
	//	data, err = db.Get([]byte(key),nil)
	//	fmt.Println(key + "     " + string(data))
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
	//for i := 0; i < 10; i++ {
	//	temp := 200000 + rand.Intn(60000)
	//	key := "testKey" + strconv.Itoa(temp)
	//	data, err = db.Get([]byte(key),nil)
	//	fmt.Println(key + "     " + string(data))
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
	datetime2 := time.Now().UnixNano() / 1000000
	datetime2s := time.Now().Second()
	fmt.Println(datetime2)
	fmt.Println(datetime1)
	fmt.Println(datetime2s - datetime1s)
	tt := float64(datetime2 - datetime1) / 1000.0
	fmt.Println(tt)
	defer db.Close()
}



func main() {

	//testDb()

	app := faygo.New("myapp", "0.1")

	// Register the route in a chain style
	app.GET("/index", new(testGet))
	app.POST("/index", new(testPost))
	app.DELETE("/index", new(testDelete))

	// Register the route in a tree style
	// app.Route(
	//     app.NewGET("/index/:id", new(Index)),
	// )



	// Start the service
	faygo.Run()

}
