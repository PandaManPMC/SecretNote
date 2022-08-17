package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	address0 = flag.String("a0", ":7080", "Zero address to bind to.")
	now      = time.Now()
)

func main() {
	fmt.Println("服务启动")
	flag.Parse()
	gracehttp.Serve(
		&http.Server{Addr: *address0, Handler: newHandler("Zero  ")},
	)
	fmt.Println("父进程退出")
}

func newHandler(name string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		fmt.Println("收到请求2")
		time.Sleep(7 * time.Second)
		w.Write([]byte("test2"))
		fmt.Println("请求处理结束2")
		fmt.Fprintf(
			w,
			"%s started at %s slept from pid %d.\n",
			name,
			now,
			os.Getpid(),
		)
	})
	return mux
}

func main3() {

	fmt.Println("服务启动")

	maxHeaderBytes := 1024 * 1024 * 20
	server := http.Server{
		Addr: ":7080",
		//Handler: mvc,
		ReadTimeout:    time.Second * 60,
		WriteTimeout:   time.Second * 60,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: maxHeaderBytes,
	}
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		fmt.Println("收到请求")
		time.Sleep(7 * time.Second)
		writer.Write([]byte("test1"))
		fmt.Println("请求处理结束")
	})

	sigs := make(chan os.Signal, 1)
	exit := make(chan struct{}, 1)

	//signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGHUP)
	signal.Notify(sigs, syscall.SIGHUP)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		// 关闭旧服务
		fmt.Println("退出服务")
		timeoutCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		server.Shutdown(timeoutCtx)
		// fork 子进程

		// 退出父进程
		//time.Sleep(5 * time.Second)
		exit <- struct{}{}
	}()

	fmt.Println(server.ListenAndServe())
	<-exit
	fmt.Println("父进程退出")
}

// fork创建子进程，并在新进程之前更新覆盖全局父进程id
//func spawnChild() error {
//	// 获取当前启动传入可执行文件参数, 如./main
//	argv0, err := exec.LookPath(os.Args[0])
//	if err != nil {
//		return err
//	}
//
//	wd, err := os.Getwd()
//	if err != nil {
//		return err
//	}
//
//	files := make([]*os.File, 0)
//	files = append(files, os.Stdin, os.Stdout, os.Stderr)
//
//	// 存下当前进程, 这个id会在新进程启动之后kill掉
//	ppid := os.Getpid()
//	os.Setenv("APP_PPID", strconv.Itoa(ppid))
//
//	// 启动新进程
//	os.StartProcess(argv0, os.Args, &os.ProcAttr{
//		Dir:   wd,
//		Env:   os.Environ(),
//		Files: files,
//		Sys:   &syscall.SysProcAttr{},
//	})
//
//	return nil
//}

func main2() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGHUP)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
