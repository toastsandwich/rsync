package rsync

import (
	"fmt"
	"net"
	"os"

	"github.com/toastsandwich/rsync/utils"
	"github.com/toastsandwich/terror"
	terr "github.com/toastsandwich/terror"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func RemoteConn(alias string) (*ssh.Client, *terr.TracedError) {
	var tracedErr *terr.TracedError
	c, tracedErr := GetConfig(alias)
	if tracedErr != nil {
		return nil, tracedErr
	}

	var knownhostFile = utils.HomeDir() + "/.ssh/known_hosts"
	knownHosts, err := knownhosts.New(knownhostFile)

	customHostkeyCallback := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		err := knownHosts(hostname, remote, key)
		if err == nil {
			return nil // all OK !
		}

		if _, ok := err.(*knownhosts.KeyError); ok { // hostey was missing
			f, err := os.OpenFile(knownhostFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
			if err != nil {
				return terror.Wrap(err, "opening knownhost file")
			}
			defer f.Close()

			line := knownhosts.Line([]string{c.Host}, key)
			if _, err := fmt.Fprintln(f, line); err != nil {
				return terr.Wrap(err, "writing to knownhost file")
			}
		}
		return nil
	}

	conf := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: customHostkeyCallback,
	}

	conn, err := ssh.Dial("tcp", c.Host, conf)
	return conn, terror.Wrap(err, "dialing host")
}
