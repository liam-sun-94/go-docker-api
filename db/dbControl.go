package db

import (
	"time"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
)

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

func contect(path string) (db *leveldb.DB, err error ){
	db, err = leveldb.OpenFile(path,nil)
	if err != nil {
		 return nil, err
	}
	return db,nil
}

func Get(path string,Id string) string{

	db, err := contect(path)
	defer db.Close()

	data, err  := db.Get([]byte(Id),nil)
	if err  != nil {
		return err .Error()
	}
	return string(data)

}
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
