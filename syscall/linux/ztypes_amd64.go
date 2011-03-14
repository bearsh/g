// godefs -g syscall -f-m64 types.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants

// Types

type Termios struct {
	Iflag        uint32
	Oflag        uint32
	Cflag        uint32
	Lflag        uint32
	Line         uint8
	Cc           [32]uint8
	Pad_godefs_0 [3]byte
	Ispeed       uint32
	Ospeed       uint32
}

type Int int32
