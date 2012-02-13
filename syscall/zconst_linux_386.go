// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs ,,const.go

//line ,,const.go:1
package syscall

//line ,,const.go:10

//line ,,const.go:9
const (
	//line ,,const.go:11
	BRKINT = 0x2
	IGNPAR = 0x4
	ICRNL  = 0x100
	INPCK  = 0x10
	ISTRIP = 0x20
	IXON   = 0x400
	OPOST  = 0x1
	ECHO   = 0x8
	ICANON = 0x2
	IEXTEN = 0x8000
	ISIG   = 0x1
	VTIME  = 0x5
	VMIN   = 0x6
	CLOCAL = 0x800
	//line ,,const.go:28

	//line ,,const.go:27
	CBAUD    = 0x100f
	B0       = 0x0
	B50      = 0x1
	B75      = 0x2
	B110     = 0x3
	B134     = 0x4
	B150     = 0x5
	B200     = 0x6
	B300     = 0x7
	B600     = 0x8
	B1200    = 0x9
	B1800    = 0xa
	B2400    = 0xb
	B4800    = 0xc
	B9600    = 0xd
	B19200   = 0xe
	B38400   = 0xf
	B57600   = 0x1001
	B115200  = 0x1002
	B230400  = 0x1003
	B460800  = 0x1004
	B500000  = 0x1005
	B576000  = 0x1006
	B921600  = 0x1007
	B1000000 = 0x1008
	B1152000 = 0x1009
	B1500000 = 0x100a
	B2000000 = 0x100b
	B2500000 = 0x100c
	B3000000 = 0x100d
	B3500000 = 0x100e
	B4000000 = 0x100f
	//line ,,const.go:62

	//line ,,const.go:61
	CSIZE = 0x30
	CS5   = 0x0
	CS6   = 0x10
	CS7   = 0x20
	CS8   = 0x30
	//line ,,const.go:69

	//line ,,const.go:68
	PARENB = 0x100
	PARODD = 0x200
	//line ,,const.go:72

	//line ,,const.go:71
	CRTSCTS = 0x80000000
	//line ,,const.go:75

	//line ,,const.go:74
	CSTOPB = 0x40
	//line ,,const.go:78

	//line ,,const.go:77
	TCGETS  = 0x5401
	TCSETS  = 0x5402
	TCSETSW = 0x5403
	TCSETSF = 0x5404
	//line ,,const.go:83

	//line ,,const.go:82
	TIOCMBIC = 0x5417
	TIOCMBIS = 0x5416
	TIOCMSET = 0x5418
	TIOCMGET = 0x5415
	//line ,,const.go:89

	//line ,,const.go:88
	TIOCM_LE  = 0x1
	TIOCM_DTR = 0x2
	TIOCM_RTS = 0x4
	TIOCM_ST  = 0x8
	TIOCM_SR  = 0x10
	TIOCM_CTS = 0x20
	TIOCM_CAR = 0x40
	TIOCM_RNG = 0x80
	TIOCM_DSR = 0x100
)
