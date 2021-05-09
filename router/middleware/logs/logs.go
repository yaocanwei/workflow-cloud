package logs

import (
    "log"
    "os"
    "path"
    "github.com/gin-gonic/gin"
)

func  Logs(logPath string) (h gin.HandlerFunc, err error) {
    f, err := os.OpenFile(path.Join(logPath, "access.log"), os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        log.Fatal(err)
    }

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }

    logCfg := gin.LoggerConfig{
        Output: f,
    }

    return gin.LoggerWithConfig(logCfg), err
}
