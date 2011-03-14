pkg=syscall

CC=/usr/bin/i586-mingw32msvc-gcc
mksyscall=$GOROOT/src/pkg/syscall/mksyscall_windows.sh

case $GOARCH in
386)
	f=-m32
	;;
amd64)
	f=-m64
	;;
*)
	echo GOARCH $GOARCH not supported
	exit 1
	;;
esac

SFX=_$GOARCH.go

$mksyscall $pkg.go |
	sed 's/^package.*syscall$$/package syscall/' |
	sed '/^import/a \
		import "syscall"' |
	sed '/import *"DISABLEDunsafe"/d' |
	sed 's/Syscall/syscall.Syscall/' |
	sed 's/EINVAL/syscall.EINVAL/' |
	gofmt > z$pkg$SFX


godefs -g $pkg -f$f -c $CC types.c  |
	sed '/Pad_godefs_0/c\
	Flags	uint32' |
	awk -f fixtype.awk |
	gofmt >ztypes$SFX


(
	echo '#include <windows.h>'
	echo 'enum {'
	sed '/^[^/]/ s/.*/	$& = &,/' < const
	echo '};'
) > ,,const.c

godefs -g $pkg -c $CC ,,const.c |
	awk -f fixtype.awk |
	gofmt > zconst$SFX
rm -f ,,const.c
