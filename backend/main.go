package main 

import (
	"home/savio/go/src/backend/src"// modifique este caminho com o diretorio da pasta
	"flag"
)

var(
	port = flag.String("p", ":3333", "set port")
)

func init() { flag.Parse() }

func main() { backend.Start(*port) }
