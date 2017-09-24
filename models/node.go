package models

// PackageJSON describes a package.json file so far as needed within this application
type PackageJSON struct {
	Name    string `json:"name"` // Application Name
	Author  string `json:"author"`
	Version string `json:"version"`

	Description string `json:"description"` // Application description
	Main        string `json:"main"`        // Entrypoint of the application

	License string `json:"license"`
}
