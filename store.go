package credential

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

const (
	// ErrDirectoryUnset indicates that the CREDENTIALS_DIRECTORY environment
	// variable is not set. This typically means the program is not running as a
	// systemd service unit with credentials configured.
	ErrDirectoryUnset xerrors.Error = "CREDENTIALS_DIRECTORY environment variable not set; is this a systemd service?"

	// ErrDirectoryAccess indicates that the credentials directory could not be
	// accessed by whatever reason, likely a permission issue.
	ErrDirectoryAccess xerrors.Error = "failed to access credentials directory"

	// ErrMissingPrefix indicates that the credential prefix was not provided. A
	// prefix is required to namespace credentials and prevent naming conflicts.
	ErrMissingPrefix xerrors.Error = "credentials prefix cannot be empty"

	// ErrInvalidName indicates that the requested credential name is invalid.
	// Names cannot be empty or contain path separators.
	ErrInvalidName xerrors.Error = "credential name cannot be empty or contain path separators"

	// ErrCredentialValue indicates that the value of the credential could not
	// be read for whatever reason.
	ErrCredentialValue xerrors.Error = "failed to read credential's value"
)

// EnvironmentVariableName is the name of the environment variable that contains
// the path to the directory where credentials are stored.
const EnvironmentVariableName = "CREDENTIALS_DIRECTORY"

// Store represents the directory where secrets are stored by systemd.
type Store struct {
	// Path is the absolute path to the credentials' directory.
	Path string

	// Prefix is the namespace prefix for credentials.
	Prefix string
}

// Open returns a new Store instance using the specified Prefix. It returns an
// error if the CREDENTIALS_DIRECTORY environment variable is not set or if the
// directory is not accessible.
func Open(prefix string) (Store, error) {
	var store Store

	if prefix == "" {
		return store, ErrMissingPrefix
	}

	path, found := os.LookupEnv(EnvironmentVariableName)
	if !found || path == "" {
		return store, ErrDirectoryUnset
	}

	if _, err := os.Stat(path); err != nil {
		return store, fmt.Errorf("%w %q: %w", ErrDirectoryAccess, path, err)
	}

	store.Path = path
	store.Prefix = strings.TrimSpace(strings.ToLower(prefix))

	return store, nil
}

// Get retrieves the credential with the given name as a string. It returns an
// error if the credential doesn't exist or is not accessible.
func (s Store) Get(name string) (string, error) {
	value, err := s.GetBytes(name)
	if err != nil {
		return "", err
	}

	return xunsafe.BytesToString(value), nil
}

// GetBytes retrieves the credential with the given name as a byte slice. It
// returns an error if the credential doesn't exist or is not accessible.
func (s Store) GetBytes(name string) ([]byte, error) {
	if name == "" || strings.ContainsAny(name, `/\`) || strings.Contains(name, "..") {
		return nil, ErrInvalidName
	}

	value, err := os.ReadFile(filepath.Join(s.Path, s.Prefix+"-"+name))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCredentialValue, err)
	}

	return value, nil
}
