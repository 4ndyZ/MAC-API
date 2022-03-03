#!/usr/bin/env sh
getent group mac-api >/dev/null || \
	groupadd -r mac-api
getent passwd mac-api >/dev/null || \
	useradd -r -g mac-api -s /sbin/nologin \
    -c "User for the MAC-API" mac-api
exit 0
