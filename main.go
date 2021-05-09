package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"

    "flowpipe-server/config"
    "flowpipe-server/router"
    "flowpipe-server/router/middleware/logs"
    "flowpipe-server/util"
    "flowpipe-server/core/model"
    // db "flowpipe-server/core/db"

    "github.com/DeanThompson/ginpprof"
    "github.com/gin-gonic/gin"
)

func setup() (e *gin.Engine, err error) {
    c := config.Cfg()

    var level string
    if c.API.Debug {
        gin.SetMode("debug")
        level = "debug"
    } else {
        gin.SetMode("release")
        level = c.Logs.Level
    }

    if err = util.InitLogs(c.Logs.LogPath, level); err != nil {
        return nil, err
    }

    //启动日志记录
    logMid, err := logs.Logs(c.Logs.LogPath)
    if err != nil {
        return nil, err
    }

    // 启动路由
    e, err = router.SetupRouter(
        logMid,
        // 跨域处理
    )

    if err != nil {
        return nil, err
    }

    // 数据库迁移
    model.Migrate()

    ginpprof.Wrap(e)
    return
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    e, err := setup()
    if err != nil {
        log.Fatal(err)
    }

    // 启动数据库连接
    // model.InitDB()

    // 启动http服务
    srv := &http.Server{
        Addr:    fmt.Sprintf(":%d", config.Cfg().API.Port),
        Handler: e,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c)
    s := <-c
    fmt.Println("Got signal:", s)
    util.Logs().Info("Shutdown Server...")

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }

    // defer db.WfDb.Close()
    util.Logs().Info("Server exiting")
}

