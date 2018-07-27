@ECHO OFF

SET srcPath=main.go
SET buildPath=runtime/hlj-comet

IF %1 == linux (
    SET GOOS=linux
    SET GOARCH=amd64

    go build -o %buildPath% %srcPath%
    ECHO success

    REM 恢复环境变量
    SET GOOS=windows
    SET GOARCH=amd64
) ELSE IF %1 == windows (
    SET GOOS=windows
    SET GOARCH=amd64

    go build -o %buildPath%.exe %srcPath%
    ECHO success
) ELSE (
    ECHO Usage: build {windows/linux}
)
