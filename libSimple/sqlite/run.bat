set GOPATH=%~dp0
echo %GOPATH%
go install db
rem ±àÒë³ÉexeÎÄ¼ş
go build src/sqlite3main.go
sqlite3main.exe
pause