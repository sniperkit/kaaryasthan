#!/bin/bash
set -e
go fmt $(glide nv -x);if [ $? != 0 ];then exit 1;fi
go vet $(glide nv -x);if [ $? != 0 ];then exit 1;fi
golint db/db.go
golint web/web.go
pkgs=$(glide nv -x|grep -v "./db/"|grep -v "./web/")
for i in ${pkgs};do golint -set_exit_status $i; if [ $? != 0 ];then ERR=1; else ERR=0;fi;done;if [ $ERR == 1 ]; then exit 1;fi
go test $(glide nv) -v -race
#for i in $(glide nv);do go test $i -v -race;done
#gometalinter $(glide nv) --errors --deadline=1m --disable=gotype
