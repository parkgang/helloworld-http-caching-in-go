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
	http.HandleFunc("/image", imageHandler)

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
		log.Println(err)
	}
	modifiedtime := file.ModTime()

	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedtime.String())))
	w.Header().Set("Etag", etag)
	w.Header().Set("Cache-Control", "max-age=5")

	// etag가 변하지 않았다면 304 응답
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			log.Println("etag가 변하지 않았음으로 304 응답")
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	// 파일을 읽은 후 응답
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	log.Println("파일 응답")
	fmt.Fprint(w, string(dat))
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	// 검은색 이미지 생성
	tempImage := color.RGBA{0, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{tempImage}, image.Point{}, draw.Src)

	// jpeg 형식으로 이미지를 인코딩하고 ResponseWriter에 writes 합니다.
	var img image.Image = m
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Cache-Control", "max-age=5")
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
