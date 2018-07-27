#!/bin/sh

# 判断进程是否存在，如果不存在就启动它
#
# 使用：
# crontab -e 添加一条记录:
# */1 * * * *  /data/hlj-comet/autorun.sh

path=/data/hlj-comet
appName=hlj-comet
logPath=runtime/autorun.log

pid=`ps -ef |grep ${appName} |grep -v -e ${path} -e grep | awk '{print $2}'`
if [ "${pid}" = "" ]; then
    cd ${path}
    rm pid
    ./run.sh start

    now=`date "+%Y-%m-%d %H:%M:%S"`
    echo "[${now}] ${path}/${appName} is started" >> ${logPath}
fi
