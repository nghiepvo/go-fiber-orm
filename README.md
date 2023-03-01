# go-fiber-orm

```shell
go mod init github.com/nghiepvo/go-fiber-orm
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

Working with Sqlite3, please run first  

```shell
GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go run main.go
```

Set up <https://github.com/cosmtrek/air>  

```shell
air init
air server --port 3000
```

Update on gitignore

```conf
# Ignore Auto reload project
tmp/
.air.toml

#Ignore DB file
*.db
```