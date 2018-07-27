#!/bin/sh

appName="hlj-comet"
path="runtime"
appPath=${path}/${appName}

tunnelPort=10800
castleDomain="castle.hljnbw.cn"
castlePort="60022"
castleUser="18268068838"
castleKeyPath="/e/keys/castle/id_rsa"

testDomain="test.hunliji.com"
testUser="deploy"
testPath="/data/www/hlj-comet"

prodDomain="10.81.86.22"
prodUser="scuser"
prodPath="/data/hlj-comet"
prodKeyPath="/e/keys/scuser/id_rsa.pub"


tunnel() {
    echo 正在打开堡垒机到正式服务器的隧道...
	ssh -i ${castleKeyPath} -N -L ${tunnelPort}:${prodDomain}:22 ${castleUser}@${castleDomain} -p ${castlePort}
}

prod() {
	echo 正在复制${appPath}到正式服务器...
	scp -i ${prodKeyPath} -P ${tunnelPort} ${appPath} ${prodUser}@127.0.0.1:${prodPath}/${appName}.new
}

pre() {
	echo 正在复制${appPath}到预发布服务器...
	scp -i ${prodKeyPath} ${appPath} ${prodUser}@${prodDomain}:${prodPath}7/${appName}7.new
}

test() {
	echo 正在复制${appPath}到测试服务器...
	scp ${appPath} ${testUser}@${testDomain}:${testPath}/${appName}.new
}

case "$1" in
    tunnel)
        tunnel
        ;;
    prod)
        prod
        ;;
    pre)
        pre
        ;;
    test)
        test
        ;;
    *)

	echo "Usage: $0 {tunnel|prod|pre|test}"
esac
