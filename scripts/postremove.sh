#!/usr/bin/env sh

remove() {
  userdel -f mac-api >/dev/null
  systemctl daemon-reload
}

purge() {
  remove
  rm -drf /etc/mac-api
  rm -drf /var/log/mac-api
}

upgrade() {
  systemctl daemon-reload
  # Restart the service if it is running
  systemctl is-active --quiet mac-api && systemctl restart mac-api
}

action="$1"
case "$action" in
  "0" | "remove")
    remove
    ;;
  "1" | "upgrade")
    upgrade
    ;;
  "purge")
    purge
    ;;
  *)
    remove
    ;;
esac
