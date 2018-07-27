#!/bin/sh

appDir="runtime"
appOldDir="runtime/old"
appName="hlj-comet"
appPath=${appDir}/${appName}

startNat() {
    ./gnatsd -m 6059 &
}

startNatL() {
    ./gnatsd -m 6059 -DV -l runtime/nats.log &
}

start() {
	if [ -f pid ]; then
		echo "有pid文件存在，请先停止进程"
		return
	fi

	if [ ! -f ${appPath} ]; then
		echo "没有找到可执行文件"
		return
	fi

	echo "正在启动..."
    run ${appPath}
	echo "进程已启动..."
}

stop() {
	if [ ! -f pid ]; then
		echo "没有pid文件，没有找到进行中的进程"
		return
	fi

    echo "正在关闭进程..."
    for i in `seq 100`
    do
        kill -s QUIT `cat -- pid`
        if [ $? = 0 ]; then
            echo "进程已经关闭"
            return
        fi

        sleep 0.1
    done

    if [ -f pid ]; then
        echo "关闭进程失败"
    fi
}

update() {
	if [ ! -f ${appPath}".new" ]; then
        echo "没有找到新文件"
        return
	fi

    echo "正在备份文件..."
    oldVer=`./${appPath} -version`
    mv ${appPath} ${appOldDir}/${appName}${oldVer}

    echo "正在替换文件..."
    mv ${appPath}.new ${appPath}
    chmod +x ${appPath}
}

use() {
	if [ -f pid ]; then
		echo "有pid文件存在，请先停止进程"
	    return
	fi

    oldPath=${appOldDir}/${appName}$1
    if [ ! -f ${oldPath} ]; then
        echo "没有找到可执行文件:" ${oldPath}
        return
    fi

    run ${oldPath}
}

run() {
	./$1 2>&1 | tee -a runtime/std.log &
}

case "$1" in
    startNat)
        startNat
        ;;
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    update)
        update
        ;;
	use)
	    use $2
		;;
    *)

	echo "Usage: $0 {startNat|start|stop|restart|update|use}"
esac
