package rsync

import (
	"fmt"

	"github.com/toastsandwich/terror"
)

func Rsync(srcAlias, dstAlias string) *terror.TracedError {
	fmt.Println("Rsync")
	return RemoteRsync(srcAlias, dstAlias)
}

// func Rsync(opts Option) {
// 	exitIf(len(opts.Args) != 2, "usage: rsync <src> <dst>")
// 	src, dst := opts.Args[0], opts.Args[1]

// 	srcF, err := os.Open(src)
// 	fatalIf(terror.Wrap(err, "open source file"))
// 	defer srcF.Close()

// 	dstF, err := os.OpenFile(dst, os.O_RDWR, os.ModeAppend)
// 	fatalIf(terror.Wrap(err, "open destination file"))
// 	defer dstF.Close()

// 	const blockSize = 32

// 	read := func(f *os.File) [][]byte {
// 		blocks := make([][]byte, 0)
// 		off := int64(0)
// 		for {
// 			p := make([]byte, blockSize)
// 			n, err := f.ReadAt(p, off)
// 			if err != nil {
// 				if err == io.EOF {
// 					blocks = append(blocks, p[:n]) // flush the buffer
// 					break
// 				} else {
// 					fatal(terror.New(err))
// 				}
// 			}
// 			blocks = append(blocks, p[:n])
// 			off += int64(n)
// 		}
// 		return blocks
// 	}

// 	srcBlocks := read(srcF)
// 	dstBlocks := read(dstF)

// 	// there are cases
// 	// 1. there is difference
// 	// 2. src is missing the block
// 	// 3. dst is missing the block
// 	for i, block := range srcBlocks {
// 		if i < len(dstBlocks) {
// 			if !bytes.Equal(block, dstBlocks[i]) {
// 				dstBlocks[i] = block
// 			}
// 		} else {
// 			dstBlocks = append(dstBlocks, block)
// 		}
// 	}

// 	off := int64(0)
// 	for _, b := range dstBlocks {
// 		n, err := dstF.WriteAt(b, off)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			} else {
// 				fatal(terror.New(err))
// 			}
// 		}
// 		off += int64(n)
// 	}
// }
