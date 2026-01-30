// usage rsync src dst
package main

import (
	"os"

	"github.com/toastsandwich/rsync/rsync"
)

func main() {
	rsync.Rsync(os.Args[1:])
}
