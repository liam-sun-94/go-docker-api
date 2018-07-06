## go-resetfulDocker

#### 通过resetful API的方式来创建docker镜像，生成镜像

目前完成 post 和 get api的测试
实现 post数据的leveldb存储和get的读取

启动main.go后   

测试地址
> http://localhost:8080/apidoc

post的数据格式
> {
> "8":{ 
>         "name":"第8题", 
>         "type":"选择题", 
>         "problem":"1+1=？", 
>         "answer":"A.1&B.2&C.3&D.4@A", 
>         "createTime":"2018-7-6 16:34", 
>         "author":"slq"
>    },
> "9":{ 
>        "name":"第9题", 
>         "type":"选择题", 
>         "problem":"1+1=？", 
>         "answer":"A.1&B.2&C.3&D.4@A", 
>         "createTime":"2018-7-6 16:34", 
>         "author":"slq"
>    },
> "10":{ 
>         "name":"第10题", 
>         "type":"选择题", 
>         "problem":"1+1=？", 
>         "answer":"A.1&B.2&C.3&D.4@A", 
>         "createTime":"2018-7-6 16:34", 
>         "author":"slq"
>    }             
> } 

get 的id值为数字
例： 8 
