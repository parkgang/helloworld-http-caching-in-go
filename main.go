package main

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var root = flag.String("root", "./assets", "file system path")

func main() {
	http.Handle("/", http.FileServer(http.Dir(*root)))
	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/black", blackHandler)

	log.Println("Listening on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	filename := "./assets/note.txt"

	// 마지막 수정 시간 가져오기
	file, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
	}
	modifiedtime := file.ModTime()

	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedtime.String())))
	w.Header().Set("Etag", etag)
	w.Header().Set("Cache-Control", "max-age=5")

	// etag가 변하지 않았다면 304 응답
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			fmt.Println("etag가 변하지 않았음으로 304 응답")
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	// 파일을 읽은 후 응답
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("파일 응답")
	fmt.Fprint(w, string(dat))
}

func blackHandler(w http.ResponseWriter, r *http.Request) {
	// 캐시 로직
	key := "black"
	e := `"` + key + `"`
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	fmt.Println("캐시되지 않아서 요청이 꽂이고 있습니다.")
	// 검은색 이미지 생성
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

// writeImage는 jpeg 형식으로 이미지를 인코딩하고 ResponseWriter에 writes 합니다.
func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
