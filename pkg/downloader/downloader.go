package downloader

import (
	"log"
	"path"
	"sync"

	"github.com/code-scan/Goal/Ghttp"
)

type Task struct {
	Name string
	URL  string
}

var task = make(chan Task, 10)
var Lock = sync.WaitGroup{}
var clock = sync.RWMutex{}
var Count int

func AddCount() {
	clock.Lock()
	Count++
	clock.Unlock()
}
func Push(t Task) {
	task <- t
}
func Downloader(url, save string) {
	var http Ghttp.Http
	http.New("GET", url)
	http.SetTimeOut(300)
	http.Execute()
	if _, err := http.SaveToFile(save); err != nil {
		log.Println("[!] Download Error : ", err)
		log.Println("[!] URL: ", url)
		return
	}
	log.Println("[*] Download Over => ", save)
}
func Worker(savePath string) {
	for {
		t := <-task
		Lock.Add(1)
		save := path.Join(savePath, t.Name)
		log.Println("[*] Start Download => ", t.Name)
		Downloader(t.URL, save)
		AddCount()
		Lock.Done()

	}
}
func Start(savePath string, max int) {
	for i := 0; i < max; i++ {
		log.Printf("[*] Create Worker_%d\n", i)
		go Worker(savePath)
	}
}
