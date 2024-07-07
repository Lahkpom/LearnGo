#Algunos comandos:

// Para listar todos los módulos que hay en el proyecto
go list -m all

// Para saber qué dependencias dependen de qué paquetes
go mod graph

// Para las versiones que están en desarrollo y no se deben considerar estables
v0.0.0

// Cuando ya se considera acpto para sacar al público
v1.0.0

// Cuando se corrigen errores o bugs
v1.0.1

// Cuando se agrega una nueva función
v1.1.0

// Cuando se modifican las funciones o se hacen cambios
v2.0.0

Este último ya indica que pierde compatibilidad con versiones anteriores