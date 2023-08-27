package assets

import "embed"

//go:embed templates
var Assets embed.FS

//go:embed optimized
var AssetsOptimized embed.FS
