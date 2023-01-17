package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var srcPath = flag.String("s", "", "source directory")
var desPath = flag.String("d", "", "destination directory")

func main() {
	//创建一个监控对象
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()
	//添加要监控的对象，文件或文件夹
	flag.Parse()
	UserSrc := *srcPath
	UserDes := *desPath

	if UserSrc == "" {
		log.Fatal("Src can not be null")
		os.Exit(0)
	}
	if UserDes == "" {
		log.Fatal("Destination can not be null")
		os.Exit(0)
	}
	err = watch.Add(UserSrc)
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(UserSrc, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			watch.Add(path)
		} else {
			oldPath := UserSrc
			newPath := UserDes
			inPath := strings.Replace(path, oldPath, "", 1)
			newFilePath := newPath + inPath
			copy(path, newFilePath)
		}

		log.Println(path, info.Size())
		return nil
	})
	log.Println(err)

	//我们另启一个goroutine来处理监控对象的事件
	go func(watcher *fsnotify.Watcher) {
		for {
			select {
			case ev := <-watch.Events:
				{
					oldPath := UserSrc
					newPath := UserDes
					inPath := strings.Replace(ev.Name, oldPath, "", 1)

					newFilePath := newPath + inPath

					// fmt.Println(newFilePath)
					//判断事件发生的类型，如下5种
					// Create
					// Write
					// Remove
					// Rename
					// Chmod
					if ev.Op&fsnotify.Create == fsnotify.Create {
						if IsDir(ev.Name) {
							log.Println("Create Directory: ", ev.Name)
							watcher.Add(ev.Name)
						} else {
							log.Println("Create File: ", ev.Name)
							copy(ev.Name, newFilePath)
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("Write : ", ev.Name)
						copy(ev.Name, newFilePath)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("Remove : ", ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("Rename : ", ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						// log.Println("Chmod : ", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}(watch)

	//循环
	select {}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	fileDirPath := path.Dir(dst)
	fmt.Println("fileDirPath", fileDirPath)
	err = os.MkdirAll(fileDirPath, 0777)
	fmt.Println("error[mkdir]: ", err)

	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	fmt.Println("写入成功")
	return nBytes, err
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false
	}
	return s.IsDir()

}
