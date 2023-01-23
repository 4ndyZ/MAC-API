#!/usr/bin/env sh

systemd_version=$(systemctl --version | head -1 | sed 's/systemd //g' | cut -d" " -f1)

cleanInstall() {
  # Create the user and group
  getent group mac-api >/dev/null || groupadd -r mac-api
  getent passwd mac-api >/dev/null || useradd -r -g mac-api -s /sbin/nologin \
     -c "User for the MAC-API" mac-api

  # RHEL/CentOS 7 cannot use ExecStartPre=+ to specify the pre start should be run as root
  # even if you want your service to run as non root.
  if [ "${systemd_version}" -lt 231 ]; then
    sed -i "s/=+/=/g" /usr/lib/systemd/system/mac-api.service
  fi

  systemctl daemon-reload
  systemctl unmask mac-api
}

upgrade() {
  systemctl daemon-reload
  # Restart the service if it is running
  systemctl is-active --quiet mac-api && systemctl restart mac-api
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
    cleanInstall
    ;;
  "2" | "upgrade")
    upgrade
    ;;
  *)
    cleanInstall
  ;;
esac