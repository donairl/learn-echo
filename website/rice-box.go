package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file3 := &embedded.EmbeddedFile{
		Filename:    "css/main.css",
		FileModTime: time.Unix(1555333842, 0),

		Content: string("body{\r\n    background: blue;\r\n    \r\n}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "css/no.html",
		FileModTime: time.Unix(1555334853, 0),

		Content: string("Denied"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1555333392, 0),

		Content: string("<html>\r\n    <bod>\r\n        Hello 1\r\n    </bod>\r\n</html>"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "second.html",
		FileModTime: time.Unix(1555335162, 0),

		Content: string("second\r\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1555335158, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file5, // "index.html"
			file6, // "second.html"

		},
	}
	dir2 := &embedded.EmbeddedDir{
		Filename:   "css",
		DirModTime: time.Unix(1555334847, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file3, // "css/main.css"
			file4, // "css/no.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir2, // "css"

	}
	dir2.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`web`, &embedded.EmbeddedBox{
		Name: `web`,
		Time: time.Unix(1555335158, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":    dir1,
			"css": dir2,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"css/main.css": file3,
			"css/no.html":  file4,
			"index.html":   file5,
			"second.html":  file6,
		},
	})
}
