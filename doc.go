// Package credential provides a simple interface for retrieving secrets from
// systemd's credential management system, which allows for secure storage and
// retrieval of sensitive information.
//
// Credential names must follow these rules:
// - Cannot be empty.
// - Cannot contain path separators (/ or \).
// - Cannot contain path traversal sequences (..).
//
// Each credential is limited to 1MB in size, as enforced by systemd.
//
// Example usage:
//
//	// Open the credential store with a prefix.
//	store, err := credential.Open("myapp")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Retrieve a secret.
//	secret, err := store.Get("database-password")
//	if err != nil {
//	    if errors.Is(err, credential.ErrInvalidName) {
//	        log.Fatal("Invalid credential name")
//	    }
//	    log.Fatal(err)
//	}
//
//	fmt.Println("Database password:", secret)
//
// The package is designed to work with systemd's credential system and expects
// the CREDENTIALS_DIRECTORY environment variable to be set. When running as a
// systemd service, credentials are typically stored in /run/credentials.
//
// Credential names are prefixed with the application name to prevent naming
// conflicts. For example, if the prefix is "myapp" and the credential name is
// "database-password", the actual file will be named "myapp-database-password".
package credential
