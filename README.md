## go 后端开发 ——> 新结构

### 包结构
----
>>>
└── user
    ├── cmd
    │   └── myapp
    │       └── main.go
    ├── service
    │   ├── repository
    │   │   ├── mongodb
    │   │   │   └── user.go
    │   │   ├── mysql
    │   │   │   └── user.go
    │   │   └── redis
    │   │       └── user.go
    │   └── user.go
    ├── transport
    │   ├── grpc
    │   │   └── user.go
    │   └── restful
    │       └── user.go
    └── user.go

---- 
- L1：cmd、service、transport、user.user
    - cmd 提供了app入口、进行初始化
    - service 包含app的逻辑和库的交互，两者之间松解耦
        - L2：repository 多方数据库交互
        - L2：service.user 实现逻辑
    - transport 根据需求实现不同的接口
        - L2：对不同的接口进行封装
    - user.go 包含提供给外部调用的结构体和方法