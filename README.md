# MAC-API
This projects provides an easy to deploy API to lookup up MAC-Vendors using the MAC address or OUI.

The API provides a fast webserver written in Golang as API endpoint. The service is getting the data from the offical IEEE MAC address assignment.

## Function
The API provides two endpoints for requests. The first one `/v1/oui` provides information about Vendor OUIs. The seconds one `/v1/mac` provides information about a MAC address

## Usage
### OUI Lookup
```Typ: GET-Request
API-Endpoint /v1/oui/<OUI>
Return: JSON
{
    "Vendor": <Vendorname>,
    "OUI": <Allocated OUI>,
    "Typ": <Prefix Typ>,
    "Address": <Vendor location>
}
```
Example:
```curl --include http://localhost:8080/v1/oui/FCFC48
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 19 Apr 2021 08:50:13 GMT
Content-Length: 104

{"Vendor":"Apple, Inc.","OUI":"FC-FC-48","Typ":"MA-L","Address":"1 Infinite Loop Cupertino CA US 95014"}
```

### MAC Lookup
```
Typ: GET-Request
API-Endpoint /v1/mac/<MAC-Address>
Return: JSON
{
    "MAC": <MAC address>
    "Vendor": <Vendorname>,
    "OUI": <Allocated OUI>,
    "Typ": <Prefix Typ>,
    "Address": <Vendor location>
}
```

Example:
```curl --include http://localhost:8080/v1/mac/FCFC48CC3E5D
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 19 Apr 2021 08:48:33 GMT
Content-Length: 130

{"MAC":"FC-FC-48-CC-3E-5D","Vendor":"Apple, Inc.","OUI":"FC-FC-48","Typ":"MA-L","Address":"1 Infinite Loop Cupertino CA US 95014"}
```

## Prerequisites
It is recommend to have a server where you can deploy the API, but it is also possible to start the microservice manually on a local machine.

## Installation and configuration
Download the prebuilt binary packages from the [release page](https://github.com/4ndyZ/MAC-API/releases) and install them on your server.

### Installation
#### Linux
###### DEB Package
If you are running a Debian-based Linux Distribution choose the `.deb` Package for your operating system architecture and download it. You are able to use curl to download the package.

Now you are able to install the package using APT.
`sudo apt install ./mac-api-vX.X-.linux.XXXX.deb`

After installing the package configure the API. The configuration file is located under `/etc/mac-api/config.yml`.

At this point you are able to enable the Systemd service using `systemctl`.
`sudo systemctl enable mac-api`

And start the service also using `systemctl`.
`sudo systemctl start mac-api`

###### RPM Package
When running a RHEL-based Linux Distribution choose the `.rpm` package for your operating system architecture and download it.

Now you are able to install the package.
`sudo dnf install ./mac-api-vX.X-.linux.XXXX.rpm`

After installing the package configure the API. The configuration file is located under `/etc/mac-api/config.yml`.

No you you are able to enable the Systemd service using `systemctl`.
`sudo systemctl enable mac-api`

And start the service also using `systemctl`.
`sudo systemctl start mac-api`

#### Windows/Other
If you plan to run the API on Windows or another OS the whole process is a bit more complicated because there is no installation package avaible only prebuilt binaries.

Download the prebuilt binary for your operating system.

Exctract the prebuilt binary and change the configuration file located under `config/config.conf`.

After successful changing the configuration file you are able to run the prebuilt binary.

### Configuration
The API tries to access the configuration file located under `/etc/mac-api/config.conf`. It the configuration file is not accessable or found the API will fallback to the local file located unter `config/config.conf`.

### Logging
The API while try to put the log file in the `/var/log/mac-api` folder. If the service is not able to access or find that folder, the logging file gets created in the local folder `logs`.

If you want to enable debug messages please change the configuration file  or run the API with the commandline parameter `-debug`.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[GPL-3.0](https://github.com/4ndyZ/MAC-API/blob/master/LICENSE)
