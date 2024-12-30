package credential_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"git.sr.ht/~jamesponddotco/credential-go"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		name          string
		giveDirectory string
		givePrefix    string
		want          credential.Store
		wantErr       bool
	}{
		{
			name:          "unset",
			giveDirectory: "",
			givePrefix:    "clivvy",
			want:          credential.Store{},
			wantErr:       true,
		},
		{
			name:          "valid repository",
			giveDirectory: "/tmp",
			givePrefix:    "clivvy",
			want: credential.Store{
				Path:   "/tmp",
				Prefix: "clivvy",
			},
			wantErr: false,
		},
		{
			name:          "non-existent directory",
			giveDirectory: "/run/credentials/non-existent",
			givePrefix:    "clivvy",
			want:          credential.Store{},
			wantErr:       true,
		},
		{
			name:          "missing Prefix",
			giveDirectory: "/tmp",
			givePrefix:    "",
			want:          credential.Store{},
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CREDENTIALS_DIRECTORY", tt.giveDirectory)

			got, err := credential.Open(tt.givePrefix)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Open() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Get(t *testing.T) {
	tests := []struct {
		name    string
		give    string
		want    string
		wantErr bool
	}{
		{
			name:    "valid credential",
			give:    "test",
			want:    "test",
			wantErr: false,
		},
		{
			name:    "non-existent credential",
			give:    "non-existent",
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid credential",
			give:    "test/credential",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			t.Setenv("CREDENTIALS_DIRECTORY", tmpDir)

			if tt.want != "" {
				if err := os.WriteFile(filepath.Join(tmpDir, "clivvy-"+tt.give), []byte(tt.want), 0o600); err != nil {
					t.Fatal(err)
				}
			}

			secret, err := credential.Open("clivvy")
			if err != nil {
				t.Fatal(err)
			}

			got, err := secret.Get(tt.give)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
