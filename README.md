[![ci](https://github.com/atotto/fileserver/actions/workflows/ci.yml/badge.svg)](https://github.com/atotto/fileserver/actions/workflows/ci.yml)
[![release](https://github.com/atotto/fileserver/actions/workflows/release.yml/badge.svg)](https://github.com/atotto/fileserver/actions/workflows/release.yml)
[![docker](https://github.com/atotto/fileserver/actions/workflows/docker.yml/badge.svg)](https://github.com/atotto/fileserver/actions/workflows/docker.yml)

fileserver
==========

simple fileserver

Install
-------

```
go install github.com/atotto/fileserver/cmd/fserv@latest
```

Or download a prebuilt binary from the [releases page](https://github.com/atotto/fileserver/releases).

Usage
-----

```
Usage of fserv:
  -addr="127.0.0.1": TCP network address
  -port=8080: port number
  -root="./": server root dir
  -tls: use tls (https)
  -cert: cert path (default "${HOME}/.config/fserv/fserv.local+1.pem")
  -key: key path (default "${HOME}/.config/fserv/fserv.local+1-key.pem")
```

Docker
------

```sh
docker pull ghcr.io/atotto/fileserver:latest

# Serve the current directory on port 8080
docker run --rm -p 8080:8080 -v "$(pwd):/data" ghcr.io/atotto/fileserver:latest
```

Supported platforms: `linux/amd64`, `linux/arm64`, `linux/arm/v7`, `linux/arm/v6`

> **TLS:** When using `-tls`, pass `-cert` and `-key` explicitly — the `$HOME` default path is not available inside the container.
