#!/usr/bin/env sh

remove() {
  systemctl stop mac-api
}

purge() {
  remove
}

action="$1"
case "$action" in
  "0" | "remove")
    remove
    ;;
  "1" | "upgrade")
    ;;
  "purge")
    purge
    ;;
  *)
    remove
    ;;
esac