#!/sbin/openrc-run

description="Colligo Web Service"

pidfile="/run/$RC_SVCNAME.pid"
command="/usr/sbin/colligo"
command_args=""