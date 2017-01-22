IPinfo - A golang IP address echo server
========================================

A small IP address echo server for use with a dynamic DNS change script.

Installation
------------

```
git clone https://github.com/kblin/ipinfo.git
cd ipinfo
go build
sudo cp ipinfo /usr/local/bin
sudo cp ipinfo.service /etc/systemd/system
systemctl daemon-reload
systemctl enable ipinfo
systemctl start ipinfo
```

License
-------

Licensed under the Apache License version 2, see [`LICENSE`](LICENSE) file for details.
