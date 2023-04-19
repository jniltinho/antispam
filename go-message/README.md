# go-message

[![godocs.io](https://godocs.io/github.com/jniltinho/antispam/go-message?status.svg)](https://godocs.io/github.com/jniltinho/antispam/go-message)
[![builds.sr.ht status](https://builds.sr.ht/~jniltinho/antispam/go-message/commits/master.svg)](https://builds.sr.ht/~jniltinho/antispam/go-message/commits/master?)

A Go library for the Internet Message Format. It implements:

* [RFC 5322]: Internet Message Format
* [RFC 2045], [RFC 2046] and [RFC 2047]: Multipurpose Internet Mail Extensions
* [RFC 2183]: Content-Disposition Header Field

## Features

* Streaming API
* Automatic encoding and charset handling (to decode all charsets, add
  `import _ "github.com/jniltinho/antispam/go-message/charset"` to your application)
* A [`mail`](https://godocs.io/github.com/jniltinho/antispam/go-message/mail) subpackage
  to read and write mail messages
* DKIM-friendly
* A [`textproto`](https://godocs.io/github.com/jniltinho/antispam/go-message/textproto)
  subpackage that just implements the wire format

## License

MIT

[RFC 5322]: https://tools.ietf.org/html/rfc5322
[RFC 2045]: https://tools.ietf.org/html/rfc2045
[RFC 2046]: https://tools.ietf.org/html/rfc2046
[RFC 2047]: https://tools.ietf.org/html/rfc2047
[RFC 2183]: https://tools.ietf.org/html/rfc2183
