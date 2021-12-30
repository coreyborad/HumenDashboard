## System Architecture
![image](https://github.com/coreyborad/HumenDashboard/blob/master/doc/struct.png?raw=true)

## Use Package

### Frontend

- [vuejs/vue](https://github.com/vuejs/vue) - Frontend framework
- [PanJiaChen/vue-element-admin](https://github.com/PanJiaChen/vue-element-admin) - Dashboard tempalte
- [vuejs/vuex](https://github.com/vuejs/vuex) - State management
- [vuejs/vue-router](https://github.com/vuejs/vue-router) - Route in frontend

### Backend - Stock

- [golang/go](https://github.com/golang/go) - Backend language
- [gin-gonic/gin](https://github.com/gin-gonic/gin) - Webserver framework
- [go-gorm/gorm](https://github.com/go-gorm/gorm) - ORM framework

### Backend - Makeup

- [php/php-src](https://github.com/php/php-src) - Backend language
- [laravel/laravel](https://github.com/laravel/laravel) - MVC and ORM framework

### PWA - Xlsx

- [facebook/react](https://github.com/facebook/react) - Create PWA with React
- [remix-run/react-router](https://github.com/remix-run/react-router) - This project use `V6` version

Notice: This part's backend is use Stock service, for quick develop

### Database

- [mysql]() - Database to save normal record
- [mongodb]() - Save stock daily record

### Other

- [googleapis/google-api-go-client](https://github.com/googleapis/google-api-go-client) - For use google sheet api on Golang

## Start

```
docker-compose up -d --build mysql nginx php-fpm go-stock mongo
```

## Build

```
GOOS=linux GOARCH=amd64 go build main.go
```