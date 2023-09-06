### 1. "获取用户信息"

1. route definition

- Url: /user/info
- Method: POST
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. request definition



```golang
type UserInfoReq struct {
	UserID int `json:"userId"`
}
```


3. response definition



```golang
type UserInfoResp struct {
	UserID int `json:"userId"`
	Nickname string `json:"nickname"`
}
```

