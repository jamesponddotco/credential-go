# credential

[![Go Documentation](https://godocs.io/git.sr.ht/~jamesponddotco/credential-go?status.svg)](https://godocs.io/git.sr.ht/~jamesponddotco/credential-go)
[![Go Report Card](https://goreportcard.com/badge/git.sr.ht/~jamesponddotco/credential-go)](https://goreportcard.com/report/git.sr.ht/~jamesponddotco/credential-go)
[![Coverage Report](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://git.sr.ht/~jamesponddotco/credential-go/tree/trunk/item/cover.out)
[![builds.sr.ht status](https://builds.sr.ht/~jamesponddotco/credential-go.svg)](https://builds.sr.ht/~jamesponddotco/credential-go?)

Package `credential` provides a simple and secure interface for
retrieving secrets from [`systemd`'s credential management
system](https://systemd.io/CREDENTIALS/). It enables Go applications to
safely access sensitive information such as cryptographic keys,
certificates, passwords, and identity data in `systemd`-managed
services.

## Installation

To install `credential` and use it in your project, run:

```console
go get git.sr.ht/~jamesponddotco/credential-go@latest
```

You'll want to ensure your system meets these requirements:

- Go 1.23 or later.
- `systemd`-based Linux distribution.
- Proper `systemd` service configuration with credentials.

## Documentation

- Please [see the Go reference
  documentation](https://godocs.io/git.sr.ht/~jamesponddotco/credential-go)
  for more information about the package.
- Refer to [`systemd` credentials
  documentation](https://systemd.io/CREDENTIALS/) for more details on
  its credential management system.

## Usage

To use `credential`, your `systemd` service unit must be configured with
credentials, as the `CREDENTIALS_DIRECTORY` environment variable
required by the package is set by `systemd` when running as a service
with credentials configured.

Example `systemd` service configuration:

```ini
[Unit]
Description=My Application Service

[Service]
ExecStart=/usr/local/bin/myapp
LoadCredential=myapp-database-password:/path/to/secret/file
PrivateMounts=yes

[Install]
WantedBy=multi-user.target
```

Here's a basic example of how to use the package:

```go
package main

import (
	"fmt"
	"log"

	"git.sr.ht/~jamesponddotco/credential-go"
)

func main() {
	// Open the credential store with your application's name as the prefix.
	store, err := credential.Open("myapp")
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve a secret from the store.
	secret, err := store.Get("database-password")
	if err != nil {
		log.Fatal(err)
	}

	// Print the secret or do something else with it.
	fmt.Println("Database password:", secret)
}
```

## Contributing

Anyone can help make `credential` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/credential-devel) and report
bugs on the [issue
tracker](https://todo.sr.ht/~jamesponddotco/credential).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the MIT License](LICENSE.md).

## Resources

The following resources are available:

- [Package documentation](https://godocs.io/git.sr.ht/~jamesponddotco/credential-go).
- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/credential-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/credential-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/credential).

---

Released under the [MIT License](LICENSE.md).
