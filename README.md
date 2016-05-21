```
             _       _    _           _      _      _
  _ _ ___ __| |_ _ _(_)__| |_   _ __ (_)_ _ (_)  __| |_ _  ___  ___ ___ _ ___ _____ _ _
 | '_/ -_|_-<  _| '_| / _|  _| | '  \| | ' \| | / _` | ' \(_-< (_-</ -_) '_\ V / -_) '_|
 |_| \___/__/\__|_| |_\__|\__| |_|_|_|_|_||_|_| \__,_|_||_/__/ /__/\___|_|  \_/\___|_|
```

A mini DNS Server to offer subdomains
restricted on a ip subnet (ipv4 and ipv6)

current it will store the information in a sqlite file
(later maybe in a postgresql file)

## Compile

`go get -v -u github.com/genofire/restrict-mini-dns-server`

## Configuration

config.yml:
```
webserver:
  enable: true
  api: true
  address: 0.0.0.0
  port: 1234
  webroot: ./webroot
database:
  path: ./test.db
dnsserver:
  enable: true
  domain: example.net
  ipv6prefix: fd2f:24e0:5e2a:93::/64
  ipv4prefix: 192.168.0.0/16
```
## Run

`restrict-mini-dns-server -config config.yml`
