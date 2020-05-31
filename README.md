# 我的web项目

1. sqlx
2. gin
3. jwt
4. zap
5. xxxxxxxxxxxxx.



├── Makefile
├── README.md
├── build
├── go.mod
├── go.sum
├── logs
│   ├── app.log
│   └── app.log.err
├── make.sh
└── src
    ├── config
    │   ├── README.md
    │   ├── config.go
    │   ├── logger.go
    │   └── viper.go
    ├── config.yaml
    ├── dao
    │   ├── user.go
    │   └── user_test.go
    ├── handler
    │   ├── handler.go
    │   ├── indexHandler
    │   │   └── index.go
    │   └── userHandler
    │       └── user.go
    ├── main.go
    ├── models
    │   ├── db.go
    │   ├── request
    │   │   └── login.go
    │   └── tables
    │       └── user.go
    ├── pkg
    │   ├── auth
    │   │   └── auth.go
    │   ├── errno
    │   │   ├── code.go
    │   │   └── errno.go
    │   ├── jwt
    │   │   └── jwt.go
    │   └── snowflake
    │       └── snowflake.go
    ├── routers
    │   ├── approuter
    │   │   └── approuter.go
    │   ├── middlewares
    │   │   ├── auth.go
    │   │   ├── gin_log.go
    │   │   └── header.go
    │   └── routers.go
    └── service
        └── user.go
