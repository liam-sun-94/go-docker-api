package db

import (
	"time"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"encoding/json"
	"os"
	"go-restfulDocker/cmd"
)

type Problem struct{
	Name string
	Type string
	Problem string
	Answer string
	CreateTime string
	Author string
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


/**
* leveldb数据库打开
* 参数 path 数据库地址
* 返回值 *level.DB 类型  error 类型
* 返回数据库的操作对象
 */

func contect(path string) (db *leveldb.DB, err error ){
	db, err = leveldb.OpenFile(path,nil)
	if err != nil {
		 return nil, err
	}
	return db,nil
}


/**
*	levelDB 数据库读取方法
*   参数 path 数据库地址
*        key  存储关键字  string类型
*	返回值 为string类型
*/
func Get(path string,key string) string{

	db, err := contect(path)
	defer db.Close()

	data, err  := db.Get([]byte(key),nil)
	if err  != nil {
		return err .Error()
	}
	return string(data)

}

/**
*	levelDB 数据库存储方法
*   参数 path 数据库地址
*        key  存储关键字  string类型
*        value 存储内容   string类型
*	返回值  为string类型
*/

func Save(path string, key string, value string) (string, error){
	db, err := contect(path)
	if err != nil {
		return "错误：", err
	}
	defer db.Close()

	err = db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return   "存储失败", err
	}
	return   "存储成功", nil
}

/**
*	levelDB 数据库删除方法
*   参数 path 数据库地址
*        key  存储关键字  string类型
*	返回值 为string类型
*/
func Delete(path string,key string) bool{

	db, err := contect(path)
	defer db.Close()

	err = db.Delete([]byte(key),nil)
	if err  != nil {
		return false //err .Error()
	}
	return true
}

/**
*	levelDB 数据库修改方法
*   参数 path 数据库地址
*        key  题目号  string类型
*        value  修改内容
*	返回值 为string类型
*/
func Put(path string, key string) bool{

	db, err := contect(path)
	defer db.Close()
    data, err := db.Get([]byte(key), nil)
	if err  != nil {
		return false
	}
    jsonStr := []byte(data)
    var jsonData Problem
    err = json.Unmarshal(jsonStr,&jsonData)
	if err != nil {
		fmt.Println("error:", err)
	}
	jsonData.Name = jsonData.Name + "1"

	mjson,_ :=json.Marshal(jsonData)
	mString :=string(mjson)

	err = db.Put([]byte(key), []byte(mString), nil)
	if err  != nil {
		return false //err .Error()
	}
	return true
}


/**
*  	临时的Dockerfile文件方法
*
*/
func CreateDockerFile(path string, test string){
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	str1 := "FROM ubuntu \n" +
	"RUN mkdir " + test + "\n"

	//str1 := "FROM centos \n RUN ifconfig \n CMD \n" + test
	//buffStr := []byte(str1)
	//re, err := f.Write(buffStr)

	re, err := f.WriteString(str1)
	if err != nil {
		panic(err)
	}
	fmt.Println(re)
	fmt.Println(path)
	//d2 := []byte{115, 111, 109, 101, 10}
	//n2, err := f.Write(d2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("wrote %d bytes\n", n2)
	//
	//n3, err := f.WriteString("writes\n")
	//fmt.Printf("wrote %d bytes\n", n3)
	//
	//
	//f.Sync()
	//
	//
	//w := bufio.NewWriter(f)
	//n4, err := w.WriteString("buffered\n")
	//fmt.Printf("wrote %d bytes\n", n4)
	//
	//
	//str1 := "FROM centos \n RUN ifconfig \n CMD"
	//buffStr := []byte(str1)
	//n5, err := f.Write(buffStr)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(n5)
	//
	//w.Flush()

}

func CreateTarfile(pathT string, pathF string){
	//"test/test1.tar"
	//"/home/alex/go/src/test_cli/go_docker_api"
	re := cmd.Lcmd("tar", []string{"-cvf", pathT, "-C", pathF,"."})
	fmt.Println(re)
}