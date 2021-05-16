package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"encoding/json"
	"os"
	"strconv"
)

func main()  {
	fmt.Print("listen on http://127.0.0.1:9090");
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/file", FileDownload)
	http.HandleFunc("/upload/file", UploadHandler)
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	var html = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<form action="/upload/file" method="post" enctype="multipart/form-data">
			<input name="file" type="file" />
			<button type="submit">提交</button>
		</form>
	</body>
	</html>
	`
	io.WriteString(w, html)
}

// UploadHandler 上传接口
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // 接收文件流
        file, fileHeader, err := r.FormFile("file")
        if err != nil {
            log.Printf("UploadHandler: 上传文件出错 -> {%s}", err)
            _, _ = io.WriteString(w, "上传文件出错！")
            return
        }
        defer file.Close()

        headerByte, _ := json.Marshal(fileHeader.Header)
        log.Printf("当前文件：Filename - >{%s}, Size -> {%v}, FileHeader -> {%s}", fileHeader.Filename, fileHeader.Size, string(headerByte))

        newFile, err := os.Create("./" + fileHeader.Filename)
        if err != nil {
            log.Printf("UploadHandler: 创建文件失败！")
			log.Printf("err %v", err)
            _, _ = io.WriteString(w, "服务器错误！")
            return
        }
        defer newFile.Close()

        // 复制文件到目标目录
        _, errCopy := io.Copy(newFile, file)
        if errCopy != nil {
            log.Printf("UploadHandler: 文件复制失败！ -> {%s}", err)
            _, _ = io.WriteString(w, "服务器错误！")
            return
        }

        // 成功响应

        // 重定向到这个请求
        http.Redirect(w, r, "/", http.StatusFound)

    }
}

func FileDownload(w http.ResponseWriter, r *http.Request) {
    filename := "E:/project/go-download-server/test.log"
	log.Println("sending %s", filename);

    file, _ := os.Open(filename)
    defer file.Close()

    fileHeader := make([]byte, 512)
    file.Read(fileHeader)

    fileStat, _ := file.Stat()

    w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
    w.Header().Set("Content-Type", http.DetectContentType(fileHeader))
    w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

    file.Seek(0, 0)
    io.Copy(w, file)

    return
}
