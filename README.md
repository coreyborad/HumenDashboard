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

### Database

- [mysql]() - Database to save normal record
- [mongodb]() - Save stock daily record

## Start

```
docker-compose up -d --build mysql nginx php-fpm go-stock mongo
```
