@echo OFF
go build -ldflags -H=windowsgui -o .\bin\msedge-launcher.exe .\main.go