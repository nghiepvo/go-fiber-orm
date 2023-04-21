# Golang + Fiber + GORM simple

A first look on Fiber and ORM on Golang  

Note: Update on .zshrc or .bashrc  

```conf
export PATH=$PATH:/usr/local/go/bin
```

```shell
go mod init github.com/nghiepvo/go-fiber-orm
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

go run main.go
```

## Auto reload project

Set up <https://github.com/cosmtrek/air>  

Note: Update on .zshrc or .bashrc  

```conf
alias air='/home/nv/go/bin/air'
```

```shell
air init
air server --port 3000

# this project was setup with command line
air -c ./cmd/.air.toml server --port 3000   
```

## Update on .gitignore

```conf
# Ignore Auto reload project
tmp/
.air.toml

#Ignore DB file
*.db
```

## Add swagger

```shell
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```

Note: Update on .zshrc or .bashrc  

```conf
export PATH=$(go env GOPATH)/bin:$PATH
```

```shell
source ~/.zshrc
# or
source ~/.bashrc
swag init
go get -u github.com/gofiber/swagger
```

```shell

```

## Refactor

- Hexagonal Architecture.
