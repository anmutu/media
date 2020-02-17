/*
  author='du'
  date='2020/2/17 15:23'
*/
package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("vid")
	vl := VIDEO_DIR + id
	video, err := os.Open(vl)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "内部错误")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "time", time.Now(), video)
	defer video.Close()
}

//从client上传文件到server端
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	//对文件做校验，查看文件是否过大
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "文件过大")
		return
	}

	//拿到file文件
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("获取文件失败:%s", err)
		SendErrorResponse(w, http.StatusInternalServerError, "内部错误")
		return
	}

	//将文件写到data的二进制里
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("y读取文件失败:%s", err)
		SendErrorResponse(w, http.StatusInternalServerError, "内部错误")
		return
	}

	//将data写入到我们要保存的文件里面
	filename := p.ByName("vid")
	err = ioutil.WriteFile(VIDEO_DIR+filename, data, 0666)
	if err != nil {
		log.Printf("写入文件失败:%s", err)
		SendErrorResponse(w, http.StatusInternalServerError, "内部错误")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "上传成功")
}
