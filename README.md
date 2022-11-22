# Casbin demo

1. 以gin建立http server
2. 建帳密檢查函數
3. 建產生jwt函數
4. 建`login` api, 通過帳密檢查後產生jwt並返回
5. 建token解析middleware, 通過返回帳號
6. 建casbin middleware, 確認帳號權限
7. 建get post api, 解析token成功,權限檢查通過後返回業務data與新token