#########################################################################
# File Name: startup_test.sh
# Author: tianlei
# mail:  taylor840326@gmail.com
# github: https://github.com/taylor840326
# Created Time: Wed Feb 28 18:07:28 2018
#########################################################################
#!/bin/bash

export ETCD_ADDR="172.18.10.136:2379"
export ETCD_PREFIX="database"
export ETCD_SVC="parauser"

# users
go test -timeout 30m -v -test.run TestCreateOrUpdateOneUser
go test -timeout 30m -v -test.run TestQueryAllUsers
go test -timeout 30m -v -test.run TestDeleteOneUser
go test -timeout 30m -v -test.run TestQueryAllUsers

# servers

