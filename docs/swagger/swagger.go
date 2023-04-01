package swagger

import "embed"

// UIPage хранит веб-морду для swagger-ui.
//
//go:embed index.html
var UIPage embed.FS
