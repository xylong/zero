go-zero场景

### 目录结构
- [`apps`](./apps)  所有微服务
  - [`app`](./apps/app) BFF服务
  - [`cart`](./apps/cart) 购物车服务
  - [`order`](./apps/order) 订单服务
  - [`pay`](./apps/pay) 支付服务
  - [`product`](./apps/product) 商品服务
  - [`recommend`](./apps/recommend) 推荐服务
  - [`reply`](./apps/reply) 评论服务
  - [`user`](./apps/user) 用户服务
- [`pkg`](./pkg)  公共目录
    - [`etcdc`](./pkg/etcdc) etcd封装
    - [`stores`](./pkg/stores) 数据存储
      - [`rocketmq`](./pkg/stores/rocketmq) - apps  所有的微服务

### api生成
```shell
goctl api go -api apps/app/api/api.api -dir apps/app/api --style=goZero
```