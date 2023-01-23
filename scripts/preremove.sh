# Package uninstall
uninstall() {
  systemctl stop mac-api.service || :
  if [ -x "/usr/lib/systemd/systemd-update-helper" ]; then
    /usr/lib/systemd/systemd-update-helper remove-system-units mac-api.service || :
  fi
}

# Package uninstall and purge
purge() {
  uninstall
}

case "$1" in
  "0" | "remove")
    uninstall
    ;;
  "1" | "upgrade")
    ;;
  "purge")
    purge
    ;;
esac