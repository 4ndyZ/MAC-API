#!/usr/bin/env sh
systemctl stop mac-api
userdel -f mac-api >/dev/null
exit 0
