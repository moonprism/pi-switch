SET GOOS=linux
SET GOARCH=arm
go fmt
go build main.go
cd web
call npm install
call npm run build
cd ..
ftp -i -s:"./ftp.txt"