backend
=======

BattleBoats API

Servers
========
[Prod](104.131.157.78)

Build Instructions
==================

* Install [Go](http://golang.org/doc/install)
  
  ```
  brew install go
  ```

* Create working dir

  ```
  mkdir backend
  cd backend
  mkdir src bin pkg
  cd src
  ```

* Clone this repo
  
  ```
  git clone git@bitbucket.org:bombsaway/backend.git
  ```

* Set GOPATH

  ```
  cd ../
  export GOPATH=`pwd`
  ```

* Get Dependencies (Only needed the first time unless you add a new dep)

  ```
  go get ./...
  ```

* Install

  ```
  go install ./...
  ```

* Run

  Prod:
  ```
  ./bin/backend --conf=../backend-deploy/conf/prod/prod.conf
  ```
  Dev:
  ```
  ./bin/backend --conf=../backend-deploy/conf/dev/dev.conf
  ```
