package tree

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestPrintTree_Basic(t *testing.T) {
	dir := setupTestDir(t, map[string]struct{}{
		"README.md":   {},
		"src/main.go": {},
		"src/lib.go":  {},
	})

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: -1})
	if err != nil {
		t.Fatal(err)
	}

	want := dir + `
├── src
│   ├── lib.go
│   └── main.go
└── README.md

1 directory, 3 files
`
	if buf.String() != want {
		t.Fatalf("output mismatch:\ngot:\n%q\nwant:\n%q", buf.String(), want)
	}
}

func TestPrintTree_SingleFile(t *testing.T) {
	dir := setupTestDir(t, map[string]struct{}{
		"README.md": {},
	})

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: -1})
	if err != nil {
		t.Fatal(err)
	}

	want := dir + `
└── README.md

0 directories, 1 file
`
	if buf.String() != want {
		t.Fatalf("output mismatch:\ngot:\n%q\nwant:\n%q", buf.String(), want)
	}
}

func TestPrintTree_MaxDepth(t *testing.T) {
	dir := setupTestDir(t, map[string]struct{}{
		"src/nested/deep.go": {},
		"README.md":          {},
	})

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: 1})
	if err != nil {
		t.Fatal(err)
	}

	want := dir + `
├── src
└── README.md

1 directory, 1 file
`
	if buf.String() != want {
		t.Fatalf("output mismatch:\ngot:\n%q\nwant:\n%q", buf.String(), want)
	}
}

func TestPrintTree_MaxDepthZero(t *testing.T) {
	dir := setupTestDir(t, map[string]struct{}{
		"README.md": {},
	})

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: 0})
	if err != nil {
		t.Fatal(err)
	}

	want := dir + `

0 directories, 0 files
`
	if buf.String() != want {
		t.Fatalf("output mismatch:\ngot:\n%q\nwant:\n%q", buf.String(), want)
	}
}

func TestPrintTree_ShowAll(t *testing.T) {
	dir := setupTestDir(t, map[string]struct{}{
		".hidden":     {},
		"visible":     {},
		".git/config": {},
	})

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: -1, All: true})
	if err != nil {
		t.Fatal(err)
	}

	want := dir + `
├── .git
│   └── config
├── .hidden
└── visible

1 directory, 3 files
`
	if buf.String() != want {
		t.Fatalf("output mismatch:\ngot:\n%q\nwant:\n%q", buf.String(), want)
	}
}

func TestPrintTree_DirsOnly(t *testing.T) {
	dir := setupTestDir(t, map[string]struct{}{
		"src/nested/deep.go": {},
		"README.md":          {},
	})

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: -1, DirsOnly: true})
	if err != nil {
		t.Fatal(err)
	}

	want := dir + `
└── src
    └── nested

2 directories
`
	if buf.String() != want {
		t.Fatalf("output mismatch:\ngot:\n%q\nwant:\n%q", buf.String(), want)
	}
}

func TestPrintTree_NoReport(t *testing.T) {
	dir := setupTestDir(t, map[string]struct{}{
		"README.md": {},
	})

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: -1, NoReport: true})
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Contains(buf.Bytes(), []byte("directories")) {
		t.Error("expected no summary line with NoReport enabled")
	}
}

func TestPrintTree_EmptyDir(t *testing.T) {
	dir := t.TempDir()

	var buf bytes.Buffer
	err := PrintTree(dir, Options{Writer: &buf, MaxDepth: -1})
	if err != nil {
		t.Fatal(err)
	}

	want := dir + "\n\n0 directories, 0 files\n"
	if buf.String() != want {
		t.Fatalf("output mismatch:\ngot:\n%q\nwant:\n%q", buf.String(), want)
	}
}

func setupTestDir(t *testing.T, files map[string]struct{}) string {
	t.Helper()
	dir := t.TempDir()
	for path := range files {
		fullPath := filepath.Join(dir, path)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}
