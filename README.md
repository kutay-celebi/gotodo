<p align="center">
<img width="250" src="https://raw.githubusercontent.com/kutay-celebi/gotodo/master/img/img-todoapp.png">
</p>

[![Build Status](https://travis-ci.com/kutay-celebi/gotodo.svg?branch=master)](https://travis-ci.com/kutay-celebi/gotodo)
[![codecov](https://codecov.io/gh/kutay-celebi/gotodo/branch/master/graph/badge.svg?token=fsnUvnu3Fx)](https://codecov.io/gh/kutay-celebi/gotodo)
![GitHub](https://img.shields.io/github/license/kutay-celebi/gotodo)
![GitHub last commit](https://img.shields.io/github/last-commit/kutay-celebi/gotodo)
[![Demo](https://img.shields.io/badge/-Demo-blue?style=flat)](http://34.142.84.33)


# Introduction

This is a simple todo application built by `Go` and `Vue`. `gin` and `gorm` which are Go library are used in backend development. There is
no UI Framework is used in Vue. The developments are made with pure CSS. Frontend developments have been made with TypeScript.

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

# About Todo App

This application, which is a simple TODO application, has been worked with new technologies.

Development processes are based on TDD. Both unit tests and E2E tests are available for the UI. E2E tests were performed with `cypress` and
unit tests with `jest`. On the GO side, an ORM tool `gorm` is used. In addition, the web framework `gin` was used on the GO application.

Pipelines are located on `travis`. The application is distributed with the help of `kubernetes`.

# Screenshot

<a href="https://raw.githubusercontent.com/kutay-celebi/gotodo/master/img/ss-1.png">
<img width="600" src="https://raw.githubusercontent.com/kutay-celebi/gotodo/master/img/ss-1.png">
</a>


<a href="https://raw.githubusercontent.com/kutay-celebi/gotodo/master/img/ss-2.png">
<img width="600" src="https://raw.githubusercontent.com/kutay-celebi/gotodo/master/img/ss-2.png">
</a>

# Run Application on Test Environment

At the root directory of project,

```shell
docker-compose -f docker-compose-test.yml up -d 
```

After that you can reach application `localhost:3000`

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
go mod download
```

4. Run Development Databse

```shell
docker-compose -f docker-compose-dev.yml up -d
```

4. Run backend

```shell
cd api
DB_HOST=localhost DB_USER=todo DB_PASS=todo go run .
```

5. Run Frontend

```shell
cd frontend
yarn serve
```
