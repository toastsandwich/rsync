# rsync

My implementation of rsync command

## So far
rsync, syncs two files on local machine.
- Creates blocks of both source file and destination file.
- Compares Source Blocks and Destination Blocks.
- If
  1. Blocks differ, swap Destination Block with Source Block
  2. If Destination Block is missing append Source Block
