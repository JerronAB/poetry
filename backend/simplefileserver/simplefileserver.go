package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

type File struct {
	Filename      string
	LinkExtension string
	Content       []byte
	UploadDate    time.Time
	ExpireTime    int
	UploadedOn    string
	ExpiresOn     string
}

var (
	files = make(map[string]File)
	mu    sync.RWMutex
)

func (f *File) isExpired() bool {
	currentTime := time.Now()
	hrs := f.ExpireTime
	expireOn := f.UploadDate.Add(time.Duration(hrs) * time.Hour)
	isExpired := currentTime.After(expireOn)
	return isExpired
}

func (f *File) expiresOn() string {
	today := time.Now().Local().Day()
	hrs := f.ExpireTime
	expireOn := f.UploadDate.Add(time.Duration(hrs) * time.Hour)
	if expireOn.Local().Day() == today {
		return expireOn.Format("3:04 PM")
	} else {
		return expireOn.Format("Jan 2 at 3:04 PM")
	}
}

func (f *File) uploadedOn() string {
	today := time.Now().Local().Day()
	uDate := f.UploadDate.Local()
	if f.UploadDate.Local().Day() == today {
		return uDate.Format("3:04 PM")
	} else {
		return uDate.Format("Jan 2 at 3:04 PM")
	}
}

func generateLink() string {
	return strconv.Itoa(rand.Intn(1000) + 1000)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(files)
	go removeExpired()
	t, err := template.ParseFiles("sfs.html")
	if err != nil {
		fmt.Fprintf(w, "Error loading template: %v", err)
		return
	}
	t.Execute(w, files)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	request, _ := url.QueryUnescape(strings.TrimPrefix(r.URL.String(), "/download/"))
	log.Printf("Downloading: %s", request)
	mu.RLock()
	content := files[request].Content
	w.Header().Set("Content-Disposition", `attachment; filename="`+files[request].Filename+`"`)
	w.Header().Set("Content-Type", "application/octet-stream")
	mu.RUnlock()
	w.Write(content)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Error uploading file: %v", err)
		return
	}
	defer file.Close()

	r.ParseForm()
	expireHours := r.FormValue("expireHours")
	exp, _ := strconv.Atoi(expireHours)

	content, err := io.ReadAll(file)
	link := generateLink()
	mu.Lock()
	f := File{
		Filename:      header.Filename,
		LinkExtension: link,
		Content:       content,
		UploadDate:    time.Now(),
		ExpireTime:    exp,
	}
	f.ExpiresOn = f.expiresOn()
	f.UploadedOn = f.uploadedOn()
	files[link] = f
	mu.Unlock()

	fmt.Printf("Uploaded File: %s\n", header.Filename)
	fmt.Printf("File Size: %d\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)
	rootHandler(w, r)
}

func removeExpired() {
	for key, file := range files {
		if file.isExpired() {
			delete(files, key)
		}
	}
}

func main() {
	http.HandleFunc("/download/", downloadHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
