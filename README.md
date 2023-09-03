# Overview
This is the example of the API that is developed by Go. 

# What's inside
- API for initialization
- API for reservation
- API for cancelation

# How to run
Clone project `git clone `
```shell
$ git clone https://github.com/Khumnin/assignment.git
```

Run build
```shell
$ docker image build -t assignment -f Dockerfile .
```

Run project
```shell
$ docker container run -p 8080:8080 --rm assignment
```