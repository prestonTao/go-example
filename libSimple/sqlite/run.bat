set GOPATH=%~dp0
echo %GOPATH%
go install db
rem �����exe�ļ�
go build src/sqlite3main.go
sqlite3main.exe
pause