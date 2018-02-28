#########################################################################
# File Name: startup_test.sh
# Author: vritxii
# mail: nkdzt@foxmail.com
# Created Time: Wed Feb 28 18:07:28 2018
#########################################################################
#!/bin/bash

export ETCD_ADDR="172.18.10.136:2379"
export ETCD_PREFIX="database"
export ETCD_SVC="parauser"

go test -timeout 30m 
