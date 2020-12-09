package main

import (
    "context"
    "errors"
    "fmt"
    "golang.org/x/sync/errgroup"
    "net"
    "net/http"
    "os"
    "os/signal"
    "time"
)

func main() {
    var c chan os.Signal = make(chan os.Signal)
    signal.Notify(c, os.Interrupt)

    ctx, cancel := context.WithCancel(context.Background())
    eg, _ := errgroup.WithContext(ctx)

    // 收到中断信号则退出
    go func() {
        <-c
        cancel()
    }()

    eg.Go(func() error {
        return serve(ctx, cancel, ":8001")
    })
    eg.Go(func() error {
        return serve(ctx, cancel, ":8002")
    })

    if err := eg.Wait(); err != nil {
        fmt.Println("程序异常退出: " + err.Error())
    } else {
        fmt.Printf("程序正常退出")
    }
}

func serve(ctx context.Context, cancel context.CancelFunc, address string) error {
    l, _ := net.Listen("tcp", address)
    srv := &http.Server{}

    Go(func() {
        select {
        case <-ctx.Done():
            err := srv.Shutdown(context.TODO())
            if err != nil {
                fmt.Println(err)
            }
        }
    })
    var err error
    if address == ":8001" {
        // 模拟其中一个服务异常中断
        time.Sleep(time.Second * 2)
        err = errors.New("模拟异常")
    } else {
        err = srv.Serve(l)
    }
    fmt.Printf("服务端 %s 退出 %+v\n", address, err)
    if err == http.ErrServerClosed {
        return nil
    }

    cancel()
    return err
}

func Go(f func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println(r)
            }

            f()
        }()
    }()
}
