package main

import (
	"time"

	"github.com/fatih/color"
)

// Windos os system
const Windows = "windows"

// *file types.
const (
	FILE_REGULAR int = iota
	FILE_DIRECTORY
	FILE_EXECUTABLE
	FILE_COMPRESS
	FILE_IMAGE
	FILE_LINK
)

// *file extensions.
const (
	EXE = ".exe"
	DEB = ".deb"
	ZIP = ".zip"
	GZ  = ".gz"
	TAR = ".tar"
	RAR = ".rar"
	PNG = ".png"
	JPG = ".jpg"
	GIF = ".gif"
)

// * File contra las estructuras para cada archivo.
type file struct {
	//* size no entendí para que sirve, dijo algo de renderizar.
	//* mode contendrá los permisos de lectura escritura que tenga el archivo.
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

// * styleFileType contendrá el estilo para cada tipo de archivo.
type styleFileType struct {
	icon   string
	color  color.Attribute
	symbol string
}

// * mapStyleByFileType es donde vamos a poner el estilo que va a tener cada tipo de dato
var mapStyleByFileType = map[int]styleFileType{
	FILE_REGULAR:    {icon: "📖"},
	FILE_DIRECTORY:  {"📁", color.FgBlue, "/"},
	FILE_EXECUTABLE: {"🚀", color.FgGreen, "*"},
	FILE_COMPRESS:   {icon: "📦", color: color.FgRed},
	FILE_IMAGE:      {icon: "🖼️", color: color.FgMagenta},
	FILE_LINK:       {icon: "⛓️", color: color.FgCyan},
}

var (
	blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	green   = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	red     = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	magenta = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	cyan    = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	yellow  = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
)

func edls() {

}
