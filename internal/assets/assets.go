// +build dev

package assets

import (
	"net/http"
	"path/filepath"
)

// Files contains web assets
var Files http.FileSystem = http.Dir(filepath.Join("web", "build"))
