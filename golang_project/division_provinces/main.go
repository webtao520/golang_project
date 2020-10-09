package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "sync"
)

// 按照34个省划分数据
// 1.创建34个省份，创建34个数据管道
// 2.读优质数据，写入对应省份管道
// 3.把省份管道写道对应文件，开34个协程

// 抽象出一个省份对象
type Province struct {
    // Id 身份证前2位
    Id string
    // 省份名
    Name string
    // 该省对应的文件，例如 北京.txt
    File *os.File
    // 本省文件的数据管道
    chanData chan string
}

// 声明等待组
var wg sync.WaitGroup

func main() {
    // 声明个map，存放所有省市的
    pMap := make(map[string]*Province)
    ps := []string{"北京市11", "天津市12", "河北省13",
        "山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",
        "黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34",
        "福建省35", "江西省36", "山东省37", "河南省41", "湖北省42",
        "湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
        "重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54",
        "陕西省61", "甘肃省62", "青海省63", "宁夏回族自治区64", "新疆维吾尔自治区65",
        "香港特别行政区81", "澳门特别行政区82", "台湾省83"}
    // 遍历所有省市，创建实例，省份管道创建
    for _, p := range ps {
        name := p[:len(p)-2]
        id := p[len(p)-2:]
        // 创建对象
        province := Province{Id: id, Name: name}
        // 添加进map
        pMap[id] = &province
        // 为当前省份开一个文件
        file, _ := os.OpenFile("./"+province.Name+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        province.File = file
        defer file.Close()
        // 创建当前省份管道
        province.chanData = make(chan string, 1024)
        fmt.Println(name, "管道已创建")
    }
    // 遍历34个省份，开34个对应文件写数据
    for _, province := range pMap {
        wg.Add(1)
        // 写入数据，这里是map中的地址
        go writeFile(province)
    }
    // 读优质文件，写入对应的省份管道
    file, _ := os.Open("./kaifang_good.txt")
    defer file.Close()
    // 缓冲读取
    reader := bufio.NewReader(file)
    // 逐行读
    for {
        lineBytes, _, err := reader.ReadLine()
        if err == io.EOF {
            for _, province := range pMap {
                close(province.chanData)
                fmt.Println(province.Name, "管道已经关闭")
            }
            break
        }
        // 转str，转utf
        lineStr := string(lineBytes)
        // 逗号切分
        fieldsSlice := strings.Split(lineStr, ",")
        id := fieldsSlice[1][0:2]
        // 对号入座，写入对应管道
        if provice, ok := pMap[id]; ok {
            provice.chanData <- (lineStr + "\n")
        } else {
            fmt.Println("未知的省", id)
        }
    }
    wg.Wait()
}

// 向文件写数据
func writeFile(province *Province) {
    for lineStr := range province.chanData {
        province.File.WriteString(lineStr)
        fmt.Print(province.Name, "写入", lineStr)
    }
    wg.Done()
}