# GODNSBL

Package godnsbl lets you perform RBL (Real-time Blackhole List - https://en.wikipedia.org/wiki/DNSBL)
lookups using Go.

The command-line tool in `cmd` demonstrates the use of [goroutines](https://tour.golang.org/concurrency/1) to perform concurrent lookups.

To test:

```
git clone https://github.com/jniltinho/antispam/godnsbl
cd godnsbl/cmd/godnsbl
go run main.go 127.0.0.2
```

The output will be a JSON-formatted list of results with the following fields:

```
[
...
{
  "rbl": "b.barracudacentral.org",
  "address": "127.0.0.2",
  "listed": true,
  "text": "http://www.barracudanetworks.com/reputation/?pr=1\u0026ip=127.0.0.2",
  "error": false,
  "error_type": null
}
...
]
```
