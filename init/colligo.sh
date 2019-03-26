#!/bin/sh
### BEGIN INIT INFO
# Provides:          colligo
# Required-Start:    $local_fs $network $named $time $syslog
# Required-Stop:     $local_fs $network $named $time $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: starts the colligo web service
# Description:       starts colligo using start-stop-daemon
### END INIT INFO

DESC="The Colligo Service"
NAME=colligo
DAEMON=/usr/local/sbin/colligo

DAEMONUSER=colligo
PIDFILE=/var/run/$NAME.pid
LOGFILE=/var/log/$NAME.log

USERBIND="setcap cap_net_bind_service=+ep"
STOP_SCHEDULE="${STOP_SCHEDULE:-QUIT/5/TERM/5/KILL/5}"
ULIMIT=8192

test -x $DAEMON || exit 0

# Set the ulimits
ulimit -n ${ULIMIT}

status() {
    if [ -f $PIDFILE ]; then
        if kill -0 $(cat "$PIDFILE"); then
            echo "$NAME is running"
        else
            echo "$NAME process is dead, but pidfile exists"
        fi
    else
        echo "$NAME is not running"
    fi
}

case "$1" in
    start)
        echo "Starting $NAME"
        start
    ;;
    stop)
        echo "Stopping $NAME"
        stop
    ;;
    restart)
        echo "Restarting $NAME"
        stop
        start
    ;;
    reload)
        echo "Reloading $NAME configuration"
        reload
    ;;
    status)
        status
    ;;
    *)
        echo "Usage: $0 {start|stop|restart|reload|status}"
        exit 2
    ;;
esac

exit 0