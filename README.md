# Admin-app

## Introduction

Admin-app is a Go-based administration panel application.

## Getting Started

### Environmental Requirements

* Go 1.22.0
* Nodejs: v14.16.0
* Npm: 6.14.11

### Development Directory Creation

### Get the Code
git clone ```[https://github.com/go-admin-team/go-admin.git]```

```bash
mkdir goadmin
cd goadmin
```


Startup Instructions

Server Startup Instructions

cd ./go-admin
go mod tidy
go build
./go-admin server -c config/settings.yml

or for Windows:
cd ./go-admin
go mod tidy
go build
go-admin.exe server -c config/settings.yml

Use Docker to Compile and Start
docker build -t go-admin .
docker run --name go-admin -p 8000:8000 -v /config/settings.yml:/config/settings.yml -d go-admin-server

UI Interactive Terminal Startup Instructions
npm install
npm run dev

Features
Fast: build a production admin panel app in ten minutes
Theming: beautiful ui themes supported (default adminlte, more themes are coming)
Plugins: many plugins to use (more useful and powerful plugins are coming)
Rbac: out of box rbac auth system
Frameworks: support most of the go web frameworks
Translation
We need your help: https://github.com/ADEMOLA200/Admin-App.git

Who is using
Comment the issue to tell us

How to
Following three steps to run it.

$ mkdir new_project && cd new_project
$ go install github.com/GoAdminGroup/adm@latest
$ adm init web

Example
Quick fork and start your admin panel app.

License
The Admin-app is licensed under the MIT License. See LICENSE for details.

Copyright (c) 2024 Admin-app Authors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

undefined
