# Rate limiter

## Run
```bash
git clone git@github.com:aaronjen/rate-limiter.git
cd rate-limiter
go run .
```

## Package Used

- [fiber](https://docs.gofiber.io/)
- [testify](https://github.com/stretchr/testify)

## Storage
使用原生的 map 作為 in-memory storage。

會使用原生的 map 的原因是：
1. 沒有任何 server 流量的要求，所以希望能簡單處理，開發時間越快越好。
2. 只有一個 server 在 run，如果是需要開多台 server 做 loading balance 的話(部署上 k8s)，就會考慮使用 [redis](https://redis.io/) 作為 storage。

## Test

```bash
go test -run Test_Get_Handler -race -v ./handler
```
