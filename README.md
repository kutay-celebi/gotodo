# Introduction

This is a simple todo application built by `Go` and `Vue`. `gin` and `gorm` which are Go library are used in backend development. There is no UI Framework is used in Vue. The developments are made with pure CSS. Frontend developments have been made with TypeScript.

# Requirements

- Docker
- Make sure you have `Go 1.16.6`
- NodeJS

# File Structure
```
/api                // backend
    /todo           // business logic of todo
    /util           // useful tool&functions 
/frontend           // frontend
    /src            
    /components     // ui components
    /layouts        // router layouts
    /model          // api response models.
    /router         // router configs
    /util           // useful tools&functions
    /views          // pages to be rendered via router
```

# Development

1. Clone Project

```shell
git clone https://github.com/kutay-celebi/gotodo.git
```

2. Install frontend dependencies

``` shell
cd frontend
yarn install
#or 
npm install
```

3. Install backend dependencies

```shell
cd api
go get
```

4. Run backend
```shell
cd api
go build
```
5. Run Frontend
```shell
cd frontend
yarn serve
```




