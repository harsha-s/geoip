# Geo IP Lookup

This is a simple Go application that looks up the geographic location of an IP address using the [IPStack](https://ipstack.com/documentation) API and obtains the lattitude and longitude of an IP address.

## Prerequisites

Before you begin, ensure you have met the following requirements:

* You have installed the latest version of [Go](https://golang.org/dl/), [Docker](https://www.docker.com/products/docker-desktop), and [Make](https://www.gnu.org/software/make/).

## Testing the App
```bash
make test
```

## Running the App as a Docker Container

To run this application as a container, follow these steps:

1. Build the Docker image

```bash
make docker-build
```
This will build the Docker image and tag it as `geoip:latest`.

2. Run the Docker container:
Example usages:
```bash
  docker run -it --env IPSTACK_API_KEY=<API_KEY> geoip:latest -i 8.8.8.8
  docker run -it --env IPSTACK_API_KEY=<API_KEY> geoip:latest -i www.google.com
```
Update `<API_KEY>` with your IPStack API key.

Use the `-i` flag to specify the IP address to look up. You can also use a domain name, and the application will resolve the IP address for you. 


## Running the App as a Command-Line Binary

To run this application as a command-line binary, follow these steps:

1. Build the binary:

For macOS:
```bash
make mac
```
This will make the binary `geoip` available in the `cmd/mac` directory.

For Linux:
```bash
make linux
```
This will make the binary `geoip` available in the `cmd/linux` directory.

2. Run the binary:
Go to the directory where the binary is located and run the following commands as shown below in the example usages:
```bash
export IPSTACK_API_KEY=<API_KEY>
./geoip -i 8.8.8.8 
./geoip -i www.google.com
```
Update `<API_KEY>` with your IPStack API key.

Use the `-i` flag to specify the IP address to look up. You can also use a domain name, and the application will resolve the IP address for you. 


