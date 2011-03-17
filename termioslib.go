// Copyright 2010 Whitham D. Reeve II. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package termioslib provides access to termios.h.
package termioslib

/*
#include <termios.h>
*/
import "C"

import (
	"os"
	"unsafe"
)

// termios types
type cc_t byte
type speed_t uint
type tcflag_t uint

type Termios struct {
    C_iflag, C_oflag, C_cflag, C_lflag tcflag_t;
    C_line cc_t;
    C_cc [NCCS]cc_t;
    C_ispeed, C_ospeed speed_t
}

// Special Control Characters
// Index into c_cc[] character array.
const (
	VEOF 		= uint (C.VEOF);
	VEOL 		= uint (C.VEOL);
	VEOL2 		= uint (C.VEOL2);
	VERASE		= uint (C.VERASE);
	VWERASE 	= uint (C.VWERASE);
	VKILL 		= uint (C.VKILL);
	VREPRINT	= uint (C.VREPRINT);
	/* 7 is spare */
	VINTR		= uint (C.VINTR);
	VQUIT		= uint (C.VQUIT);
	VSUSP		= uint (C.VSUSP);
	VDSUSP		= uint (C.VDSUSP);
	VSTART		= uint (C.VSTART);
	VSTOP		= uint (C.VSTOP);
	VLNEXT		= uint (C.VLNEXT);
	VDISCARD	= uint (C.VDISCARD);
	VMIN		= uint (C.VMIN);
	VTIME		= uint (C.VTIME);
	VSTATUS		= uint (C.VSTATUS);
	/* 19 is spare */
	NCCS		= uint (C.NCCS)
)

// Input flags - software input processing
const (
	IGNBRK		= tcflag_t (C.IGNBRK);	/* ignore BREAK condition */
	BRKINT		= tcflag_t (C.BRKINT); 	/* map BREAK to SIGINTR */
	IGNPAR		= tcflag_t (C.IGNPAR); 	/* ignore (discard) parity errors */
	PARMRK		= tcflag_t (C.PARMRK); 	/* mark parity and framing errors */
	INPCK		= tcflag_t (C.INPCK); 	/* enable checking of parity errors */
	ISTRIP		= tcflag_t (C.ISTRIP); 	/* strip 8th bit off chars */
	INLCR		= tcflag_t (C.INLCR); 	/* map NL into CR */
	IGNCR		= tcflag_t (C.IGNCR); 	/* ignore CR */
	ICRNL		= tcflag_t (C.ICRNL); 	/* map CR to NL (ala CRMOD) */
	IXON		= tcflag_t (C.IXON); 	/* enable output flow control */
	IXOFF		= tcflag_t (C.IXOFF); 	/* enable input flow control */
	IXANY		= tcflag_t (C.IXANY); 	/* any char will restart after stop */
	IMAXBEL		= tcflag_t (C.IMAXBEL); /* ring bell on input queue full */
	IUTF8		= tcflag_t (C.IUTF8) 	/* maintain state for UTF-8 VERASE */
)

// Output flags - software output processing
const (
	OPOST		= tcflag_t (C.OPOST); 	/* enable following output processing */
	ONLCR		= tcflag_t (C.ONLCR); 	/* map NL to CR-NL (ala CRMOD) */
	OXTABS		= tcflag_t (C.OXTABS); 	/* expand tabs to spaces */
	ONOEOT		= tcflag_t (C.ONOEOT) 	/* discard EOT's (^D) on output) */
)

// Standard speeds
const (
	B0			= speed_t (C.B0);
	B50			= speed_t (C.B50);
	B75			= speed_t (C.B75);
	B110		= speed_t (C.B110);
	B134		= speed_t (C.B134);
	B150		= speed_t (C.B150);
	B200		= speed_t (C.B200);
	B300		= speed_t (C.B300);
	B600		= speed_t (C.B600);
	B1200		= speed_t (C.B1200);
	B1800		= speed_t (C.B1800);
	B2400		= speed_t (C.B2400);
	B4800		= speed_t (C.B4800);
	B9600		= speed_t (C.B9600);
	B19200		= speed_t (C.B19200);
	B38400		= speed_t (C.B38400);
	// non posix but generally supported none the less
	B7200		= speed_t (C.B7200);
	B14400		= speed_t (C.B14400);
	B28800		= speed_t (C.B28800);
	B57600		= speed_t (C.B57600);
	B76800		= speed_t (C.B76800);
	B115200		= speed_t (C.B115200);
	B230400		= speed_t (C.B230400);
	EXTA		= speed_t (C.EXTA);
	EXTB		= speed_t (C.EXTB);
)

// Control flags - hardware control of terminal
const (
	CIGNORE		= tcflag_t (C.CIGNORE);	/* ignore control flags */
	CSIZE		= tcflag_t (C.CSIZE);	/* character size mask */
		CS5		= tcflag_t (C.CS5);		/* 5 bits (pseudo) */
		CS6		= tcflag_t (C.CS6);		/* 6 bits */
		CS7		= tcflag_t (C.CS7);		/* 7 bits */
		CS8		= tcflag_t (C.CS8);		/* 8 bits */
	CSTOPB		= tcflag_t (C.CSTOPB);	/* send 2 stop bits */
	CREAD		= tcflag_t (C.CREAD);	/* enable receiver */
	PARENB		= tcflag_t (C.PARENB);	/* parity enable */
	PARODD		= tcflag_t (C.PARODD);	/* odd parity, else even */
	HUPCL		= tcflag_t (C.HUPCL);	/* hang up on last close */
	CLOCAL		= tcflag_t (C.CLOCAL);	/* ignore modem status lines */
	CCTS_OFLOW	= tcflag_t (C.CCTS_OFLOW);	/* CTS flow control of output */
	CRTS_IFLOW	= tcflag_t (C.CRTS_IFLOW);	/* RTS flow control of input */
	CRTSCTS		= tcflag_t (C.CRTSCTS);
	CDTR_IFLOW	= tcflag_t (C.CDTR_IFLOW);	/* DTR flow control of input */
	CDSR_OFLOW	= tcflag_t (C.CDSR_OFLOW);	/* DSR flow control of output */
	CCAR_OFLOW	= tcflag_t (C.CCAR_OFLOW);	/* DCD flow control of output */
	MDMBUF		= tcflag_t (C.MDMBUF)		/* old name for CCAR_OFLOW */
)

// "Local" flags - dumping ground for other state
const (
	ECHOKE		= tcflag_t (C.ECHOKE);	/* visual erase for line kill */
	ECHOE		= tcflag_t (C.ECHOE);	/* visually erase chars */
	ECHOK		= tcflag_t (C.ECHOK);	/* echo NL after line kill */
	ECHO		= tcflag_t (C.ECHO);	/* enable echoing */
	ECHONL		= tcflag_t (C.ECHONL);	/* echo NL even if ECHO is off */
	ECHOPRT		= tcflag_t (C.ECHOPRT);	/* visual erase mode for hardcopy */
	ECHOCTL		= tcflag_t (C.ECHOCTL);	/* echo control chars as ^(Char) */
	ISIG		= tcflag_t (C.ISIG);	/* enable signals INTR, QUIT, [D]SUSP */
	ICANON		= tcflag_t (C.ICANON);	/* canonicalize input lines */
	ALTWERASE	= tcflag_t (C.ALTWERASE);	/* use alternate WERASE algorithm */
	IEXTEN		= tcflag_t (C.IEXTEN);	/* enable DISCARD and LNEXT */
	EXTPROC		= tcflag_t (C.EXTPROC);	/* external processing */
	TOSTOP		= tcflag_t (C.TOSTOP);	/* stop background jobs from output */
	FLUSHO		= tcflag_t (C.FLUSHO);	/* output being flushed (state) */
	NOKERNINFO	= tcflag_t (C.NOKERNINFO);	/* no kernel output from VSTATUS */
	PENDIN		= tcflag_t (C.PENDIN);	/* XXX retype pending input (state) */
	NOFLSH		= tcflag_t (C.NOFLSH)	/* don't flush after interrupt */
)

// Commands passed to tcsetattr() for setting the termios structure.
const (
	TCSANOW		= int (C.TCSANOW);		/* make change immediate */
	TCSADRAIN	= int (C.TCSADRAIN);	/* drain output, then change */
	TCSAFLUSH	= int (C.TCSAFLUSH)	/* drain output, flush input */
)

const (
	TCIFLUSH	= int (C.TCIFLUSH);
	TCOFLUSH	= int (C.TCOFLUSH);
	TCIOFLUSH	= int (C.TCIOFLUSH);
	TCOOFF		= int (C.TCOOFF);
	TCOON		= int (C.TCOON);
	TCIOFF		= int (C.TCIOFF);
	TCION		= int (C.TCION)
)

//speed_t	cfgetispeed(const struct Termios *);
func Getispeed (src * Termios) (result speed_t) {
	result = speed_t(C.cfgetispeed((*C.struct_termios)(unsafe.Pointer(src))))
	return
}

//speed_t	cfgetospeed(const struct Termios *);
func Getospeed (src * Termios) (result speed_t) {
	result = speed_t(C.cfgetospeed((*C.struct_termios)(unsafe.Pointer(src))))
	return
}

//int	cfsetispeed(struct Termios *, speed_t);
func Setispeed(dst * Termios, baud speed_t) (err os.Error) {
	rv := C.cfsetispeed((*C.struct_termios)(unsafe.Pointer(dst)), C.speed_t(baud))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}

//int	cfsetospeed(struct Termios *, speed_t);
func Setospeed(dst * Termios, baud speed_t) (err os.Error) {
	rv := C.cfsetospeed((*C.struct_termios)(unsafe.Pointer(dst)), C.speed_t(baud))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}

func Getattr(fd int, dst * Termios) (err os.Error) {
	rv := C.tcgetattr(C.int(fd),(*C.struct_termios)(unsafe.Pointer(dst)))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}

//int	tcsetattr(int, int, const struct termios *);
func Setattr(fd int, optional_action int, src * Termios) (err os.Error) {
	rv := C.tcsetattr(C.int(fd), C.int(optional_action), (*C.struct_termios)(unsafe.Pointer(src)))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}

//int	tcdrain(int) __DARWIN_ALIAS_C(tcdrain);
func Drain(fd int) (err os.Error) {
	rv := C.tcdrain(C.int(fd))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}

//int	tcflow(int, int);
func Flow(fd int, action int) (err os.Error) {
	rv := C.tcflow(C.int(fd), C.int(action))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}

//int	tcflush(int, int);
func Flush(fd int, queue_selector int) (err os.Error) {
	rv := C.tcflush(C.int(fd), C.int(queue_selector))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}
//int	tcsendbreak(int, int);
func Sendbreak(fd int, duration int) (err os.Error) {
	rv := C.tcflush(C.int(fd), C.int(duration))
	if rv != 0 {
		return os.Errno(rv)
	}
	return nil
}
//void	cfmakeraw(struct termios *);
func Makeraw(src * Termios) {
	C.cfmakeraw((*C.struct_termios)(unsafe.Pointer(src)))
}

// int	cfsetspeed(struct termios *, speed_t);
func Setspeed (dst * Termios, baud speed_t) (err os.Error) {
	rv := C.cfsetspeed(((*C.struct_termios)(unsafe.Pointer(dst))), C.speed_t(baud))
	if rv != 0 {
			return os.Errno(rv)
	}
	return nil
}
