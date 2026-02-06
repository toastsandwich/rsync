// usage rsync src dst
package main

import (
	"github.com/toastsandwich/rsync/cmd"
)

func main() {
	// rsync.Rsync(os.Args[1:])
	cmd.Root()
}
