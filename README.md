# Casbin demo

1. 以gin建立http server
2. call POST API`login`取得token, body`{
   "account": "thomas",
   "password": "123456"
   }`
3. call其他API加入`Bearer Token`
4. rate limit
5. 使用jwt
6. 使用casbin管理權限

### login
```shell
curl --location --request POST '127.0.0.1:9109/v1/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account": "thomas",
    "password": "123456"
}'
```

