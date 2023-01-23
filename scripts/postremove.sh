# Package uninstall
uninstall() {
  userdel -f mac-api >/dev/null || :
}

# Package uninstall and purge
purge() {
  rm -drf /etc/mac-api || :
  rm -drf /var/log/mac-api || :
}

# Package upgrade
upgrade() {
  :
}

action="$1"
case "$action" in
  "0" | "remove")
    uninstall
    ;;
  "1" | "upgrade")
    upgrade
    ;;
  "purge")
    uninstall
    purge
    ;;
esac
