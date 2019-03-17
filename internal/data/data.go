// +build dev

package data

import (
	"net/http"
	"path/filepath"
)

// Assets contains web assets.
var Assets http.FileSystem = http.Dir(filepath.Join("../../", "web", "build"))

// Templates contains web templates.
var Templates http.FileSystem = http.Dir(filepath.Join("../../", "web", "templates"))
