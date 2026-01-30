package rsync

import (
	"fmt"
	"log"
	"os"

	"github.com/toastsandwich/terror"
)

func fatal(err *terror.TracedError) {
	log.Fatalf("%s\nstack:%s", err.Error(), err.Trace())
	log.Fatal(err)
}

func fatalIf(err *terror.TracedError) {
	if err != nil {
		log.Fatalf("%s\n\tstack:%s", err.Error(), err.Trace())
	}
}

func exitIf(cond bool, msg string) {
	if cond {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
