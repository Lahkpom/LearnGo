package main

//*File types
const (
	fileRegular   int = iota //*Caulquier archivo que no entre dentro de ninguna clasificación concreta
	filedirectory            //*Carpetas
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

//*File extension
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)
