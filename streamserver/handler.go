/*
  author='du'
  date='2020/2/17 15:23'
*/
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

func main() {

}

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
