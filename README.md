# MAC-API


## Function


## Example


## Prerequisites
It is recommend to have a server where you can deploy the API, but it is also possible to start the microservice manually on a local machine.

## Installation and configuration
Download the prebuilt binary packages from the [release page](https://github.com/4ndyZ/MAC-API/releases) and install them on your server.

### Installation
#### Linux
###### DEB Package
If you are running a Debian-based Linux Distribution choose the `.deb` Package for your operating system architecture and download it. You are able to use curl to download the package.

Now you are able to install the package using APT.
`sudo apt install ./MAC-API-vX.X-.linux.XXXX.deb`

After installing the package configure the microservice. The configuration file is located under `/etc/corona-dashboard/config.yml`.

At this point you are able to enable the Systemd service using `systemctl`.
`sudo systemctl enable mac-api`

And start the service also using `systemctl`.
`sudo systemctl start mac-api`

###### RPM Package
When running a RHEL-based Linux Distribution choose the `.rpm` package for your operating system architecture and download it.

Now you are able to install the package.
`sudo rpm -i MAC-API-vX.X-.linux.XXXX.rpm`

After installing the package configure the microservice. The configuration file is located under `/etc/corona-dashboard/config.yml`.

No you you are able to enable the Systemd service using `systemctl`.
`sudo systemctl enable mac-api`

And start the service also using `systemctl`.
`sudo systemctl start mac-api`

#### Windows/Other
If you plan to run the microservice on Windows or another OS the whole process is a bit more complicated because there is no installation package avaible only prebuilt binaries.

Download the prebuilt binary for your operating system.

Exctract the prebuilt binary and change the configuration file located under `config/config.conf`.

After successful changing the configuration file you are able to run the prebuilt binary.

### Configuration
The microservice tries to access the configuration file located under `/etc/mac-api/config.conf`. It the configuration file is not accessable or found the service will fallback to the local file located unter `config/config.conf`.

### Logging
The microservice while try to put the log file in the `/var/log/mac-api` folder. If the service is not able to access or find that folder, the logging file gets created in the local folder `logs`.

If you want to enable debug messages please change the configuration file  or run the microservice with the commandline parameter `-debug`.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[GPL-3.0](https://github.com/4ndyZ/Corona-Dashboard/blob/master/LICENSE)
