package main

import "time"

//*file types.
const (
	FILE_REGULAR int = iota
	FILE_DIRECTORY
	FILE_EXECUTABLE
	FILE_COMPRESS
	FILE_IMAGE
	FILE_LINK
)

//*file extension.
const (
	EXE = ".exe"
	DEB = ".deb"
	GZ  = ".gz"
	TAR = ".tar"
	RAR = ".rar"
	PNG = ".png"
	JPG = ".jpg"
	GIF = ".gif"
)

//* File contra las estructuras para cada archivo.
type File struct {
	//* size no entendí para que sirve, dijo algo de renderizar.
	//* mode contendrá los permisos de lectura escritura que tenga el archivo.
	name             string
	fileType         int
	isDirectory      bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	lastModification time.Time
	mode             string
}

//* styleFileType contendrá el estilo para cada tipo de archivo.
type styleFileType struct {
	icon   string
	color  string
	symbol string
}

//* mapStyleByFileType es donde vamos a poner el estilo que va a tener cada tipo de dato
var mapStyleByFileType = map[int]styleFileType{
	FILE_REGULAR:    {icon: "📖"},
	FILE_DIRECTORY:  {"📁", "BLUE", "/"},
	FILE_EXECUTABLE: {"🚀", "GREEN", "*"},
	FILE_COMPRESS:   {icon: "📦", color: "RED"},
	FILE_IMAGE:      {icon: "🖼️", color: "MAGENTA"},
	FILE_LINK:       {icon: "⛓️", color: "CYAN"},
}

func edls() {

}
