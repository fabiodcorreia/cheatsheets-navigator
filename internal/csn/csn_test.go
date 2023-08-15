package csn

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestExpandHomeDir(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		// Test cases for expanding ~/
		{input: "~/file.txt", expected: filepath.Join(homeDir(), "file.txt")},
		{input: "~/", expected: homeDir()},
		{input: "~/dir/file.txt", expected: filepath.Join(homeDir(), "dir/file.txt")},

		// Test cases for expanding $HOME/
		{input: "$HOME/file.txt", expected: filepath.Join(homeDir(), "file.txt")},
		{input: "$HOME/", expected: homeDir()},
		{input: "$HOME/dir/file.txt", expected: filepath.Join(homeDir(), "dir/file.txt")},

		// Test cases for no expansion
		{input: "path/file.txt", expected: "path/file.txt"},
		{input: "/absolute/path/file.txt", expected: "/absolute/path/file.txt"},
		{input: "", expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := expandHomeDir(tc.input)
			if result != tc.expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, result)
			}
		})
	}
}

func TestGetRepository(t *testing.T) {
	tests := []struct {
		name        string
		repoPath    string
		expectedAbs string
		expectedErr error
	}{
		{
			name:        "ValidPath",
			repoPath:    "~/",
			expectedAbs: homeDir(),
			expectedErr: nil,
		},
		{
			name:        "EmptyPath",
			repoPath:    "",
			expectedAbs: "",
			expectedErr: fmt.Errorf("environment variable %q not found", "CSN_PAGES"),
		},
		{
			name:        "PathNotExist",
			repoPath:    "/non_existent_path",
			expectedAbs: "",
			expectedErr: fmt.Errorf("repository path %q doesn't exists", "/non_existent_path"),
		},
		{
			name:        "NotADirectory",
			repoPath:    filepath.Join(currentDir(), "./csn_test.go"),
			expectedAbs: "",
			expectedErr: fmt.Errorf("repository path %q is not a directory", filepath.Join(currentDir(), "./csn_test.go")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			absPath, err := getRepository(test.repoPath)

			if absPath != test.expectedAbs {
				t.Errorf("Expected absolute path: %q, but got: %q", test.expectedAbs, absPath)
			}

			if err != nil && test.expectedErr == nil {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && test.expectedErr != nil {
				t.Errorf("Expected error: %v, but got no error", test.expectedErr)
			} else if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("Expected error: %v, but got: %v", test.expectedErr, err)
			}
		})
	}
}

func homeDir() string {
	dirname, _ := os.UserHomeDir()
	return dirname
}

func currentDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
