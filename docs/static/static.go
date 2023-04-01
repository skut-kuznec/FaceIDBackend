package static

import "embed"

// SwaggerStatics хранит файлы для статики для swagger-ui.
//
//go:embed *
var SwaggerStatics embed.FS
