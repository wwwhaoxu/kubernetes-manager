## kubernets-manager项目

## test


### Run
```shell
make build
 ./_output/km -c configs/km.yaml
```

### Add User
```shell
curl -XPOST -H"Content-Type: application/json" -d'{"username":"root","password":"root1234","role":"admin""nickname":"root","email":"root@qq.com","phone":"18888888xxxx"}' http://127.0.0.1:8080/v1/users
```
### Test User Login
```shell
curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"root","password":"root1234"}' http://127.0.0.1:8080/login

# Output
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJYLVVzZXJuYW1lIjoiYXV0aG50ZXN0IiwiZXhwIjoyMDMwMjk3NjUyLCJpYXQiOjE2NzAyOTc2NTIsIm5iZiI6MTY3MDI5NzY1Mn0.wzpMG6hOljfPjczAKvRjBRtMa-U6K2Vu9Pmd7t9QDrM"}
```
### 获取 User root 的详细信息
```shell
curl -XGET -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/users/root
```