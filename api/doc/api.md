### 1. "登录"

1. route definition

- Url: /v1/user/login
- Method: POST
- Request: `reqLogin`
- Response: `response`

2. request definition



```golang
type ReqLogin struct {
	Version string `json:"version" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	RequestTime int64 `json:"requestTime,string" binding:"required"`
	AppId string `json:"appId" binding:"required"`
}

type Request struct {
	Version string `json:"version" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	RequestTime int64 `json:"requestTime,string" binding:"required"`
	AppId string `json:"appId" binding:"required"`
}
```


3. response definition



```golang
type Response struct {
	Modult string `json:"modult"`
	Code string `json:"code"`
	Message string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Data interface{} `json:"data"`
}
```

### 2. "用户列表"

1. route definition

- Url: /v1/users/list
- Method: POST
- Request: `reqList`
- Response: `response`

2. request definition



```golang
type ReqList struct {
	Version string `json:"version" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	RequestTime int64 `json:"requestTime,string" binding:"required"`
	AppId string `json:"appId" binding:"required"`
	Offset int `json:"offset,string" binding:"required,min=0,max=20"`
	Limit int `json:"limit,string" binding:"required,min=0,max=20"`
}

type Request struct {
	Version string `json:"version" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	RequestTime int64 `json:"requestTime,string" binding:"required"`
	AppId string `json:"appId" binding:"required"`
}
```


3. response definition



```golang
type Response struct {
	Modult string `json:"modult"`
	Code string `json:"code"`
	Message string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Data interface{} `json:"data"`
}
```

### 3. "登出"

1. route definition

- Url: /v1/users/out
- Method: POST
- Request: `reqOut`
- Response: `response`

2. request definition



```golang
type ReqOut struct {
	Version string `json:"version" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	RequestTime int64 `json:"requestTime,string" binding:"required"`
	AppId string `json:"appId" binding:"required"`
}

type Request struct {
	Version string `json:"version" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	RequestTime int64 `json:"requestTime,string" binding:"required"`
	AppId string `json:"appId" binding:"required"`
}
```


3. response definition



```golang
type Response struct {
	Modult string `json:"modult"`
	Code string `json:"code"`
	Message string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Data interface{} `json:"data"`
}
```

