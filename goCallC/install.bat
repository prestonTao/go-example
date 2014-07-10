set GOPATH=%~dp0
echo %GOPATH%
go install hover
go install prints
go install inc
go run src/main.go
pause