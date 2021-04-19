#!/bin/sh
systemctl stop mac-api
getent passwd mac-api >/dev/null || \
	userdel -f mac-api
getent group mac-api >/dev/null || \
	groupdel mac-api
exit 0
