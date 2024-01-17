git add .
git commit -m "Ultimo Commit"
git push
set GOOS=linux
set GOARCH=amd64
go build ./src/ main.go
del ./src/main.zip
tar.exe -a -cf main.zip main