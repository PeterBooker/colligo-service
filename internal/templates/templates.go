// +build dev

package templates

import (
	"net/http"
	"path/filepath"
)

// Files contains web templates
var Files http.FileSystem = http.Dir(filepath.Join("web", "templates"))
