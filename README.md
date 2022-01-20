## Description

...Todo

## Table of Contents

- [Description](#description)
- [Table of Contents](#table-of-contents)
- [Features](#features)
- [System Architecture](#system-architecture)
    - [Makeup and Stock Management](#makeup-and-stock-management)
    - [PWA for Google Sheet append form](#pwa-for-google-sheet-append-form)
    - [Schedule](#schedule)
- [Use Package](#use-package)
    - [Frontend](#frontend)
    - [Backend on Stock Service](#backend-on-stock-service)
    - [Backend on MakeUp Service](#backend-on-makeup-service)
    - [PWA on GoogleSheet form](#pwa-on-googlesheet-form)
    - [Database](#database)
- [Command for memo](#command-for-memo)
    - [Start Docker](#start-docker)
    - [Build Golang](#build-golang)

## Features

...Todo

## System Architecture

### Makeup and Stock Management
<img src="./doc/MakeupAndStockManagement.png?"  width="75%">

### PWA for Google Sheet append form
<img src="./doc/PWAforGoogleSheetAppendForm.png?"  width="75%">

### Schedule
<img src="./doc/Schedule.png?"  width="75%">

## Use Package

### Frontend

- [vuejs/vue](https://github.com/vuejs/vue) - Frontend framework
- [PanJiaChen/vue-element-admin](https://github.com/PanJiaChen/vue-element-admin) - Dashboard tempalte
- [vuejs/vuex](https://github.com/vuejs/vuex) - State management
- [vuejs/vue-router](https://github.com/vuejs/vue-router) - Route in frontend

### Backend on Stock Service

- [golang/go](https://github.com/golang/go) - Backend language
- [gin-gonic/gin](https://github.com/gin-gonic/gin) - Webserver framework
- [go-gorm/gorm](https://github.com/go-gorm/gorm) - ORM framework

### Backend on MakeUp Service

- [php/php-src](https://github.com/php/php-src) - Backend language
- [laravel/laravel](https://github.com/laravel/laravel) - MVC and ORM framework

### PWA on GoogleSheet form

- [facebook/react](https://github.com/facebook/react) - Create PWA with React
- [remix-run/react-router](https://github.com/remix-run/react-router) - This project use `V6` version
- [googleapis/google-api-go-client](https://github.com/googleapis/google-api-go-client) - For use google sheet api on Golang

```Notice: This part's backend is use Stock service, for quick develop```

### Database

- [mysql]() - Database to save normal record
- [mongodb]() - Save stock history

## Command for memo

### Start Docker
```sh
docker-compose up -d --build mysql nginx php-fpm go-stock mongo
```

### Build Golang

```sh
GOOS=linux GOARCH=amd64 go build main.go
```