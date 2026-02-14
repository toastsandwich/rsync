package rsync

import (
	"fmt"

	terr "github.com/toastsandwich/terror"
	"golang.org/x/crypto/ssh"
)

func RemoteRsync(srcAlias, dstAlias string) *terr.TracedError {
	fmt.Println("RemoteSync")
	// srcConn, err := RemoteConn(srcAlias)
	// if err != nil {
	// 	return terr.New(err)
	// }
	// defer srcConn.Close()

	dstConn, err := RemoteConn(dstAlias)
	if err != nil {
		return terr.New(err)
	}
	defer dstConn.Close()

	return rsync(nil, dstConn)
}

func rsync(src, dst *ssh.Client) *terr.TracedError {
	request := "ls"
	// ok, reply, err := dst.SendRequest("example", true, []byte(request))
	// if err != nil {
	// 	return terr.Wrapf(err, "sending request to %s", dst.RemoteAddr())
	// }
	// else if !ok {
	// 	return terr.Newf("something went wrong")
	// }
	sess, err := dst.NewSession()
	if err != nil {
		return terr.New(err)
	}
	defer sess.Close()

	out, err := sess.CombinedOutput(request)
	if err != nil {
		return terr.New(err)
	}
	fmt.Println(string(out))
	return nil
}
