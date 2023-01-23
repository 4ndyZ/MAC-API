systemd_version=$(systemctl --version | head -1 | sed 's/systemd //g' | cut -d" " -f1)

# Initial package installation
install() {
  if [ -x "/usr/lib/systemd/systemd-update-helper" ]; then
    /usr/lib/systemd/systemd-update-helper install-system-units mac-api.service || :
  fi
}

# Package upgrade
upgrade() {
  if [ -x "/usr/lib/systemd/systemd-update-helper" ]; then
    /usr/lib/systemd/systemd-update-helper mark-restart-system-units mac-api.service || :
    /usr/lib/systemd/systemd-update-helper system-reload-restart || :
  fi
}

# Fix for old distributions that cannot use ExecStartPre=+ to specify the pre start should be run as root
# even if you want your service to run as non root.
fix_old() {
  if [ "${systemd_version}" -lt 231 ]; then
    sed -i "s/=+/=/g" /etc/systemd/system/mac-api.service || :
  fi
}

action="$1"
if  [ "$1" = "configure" ] && [ -z "$2" ]; then
  # deb passes $1=configure
  action="install"
elif [ "$1" = "configure" ] && [ -n "$2" ]; then
  # deb passes $1=configure $2=<current version>
  action="upgrade"
fi

case "$action" in
  "1" | "install")
    fix_old
    install
    ;;
  "2" | "upgrade")
    fix_old
    upgrade
    ;;
esac