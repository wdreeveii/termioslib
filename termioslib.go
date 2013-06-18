// Copyright 2010 Whitham D. Reeve II. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package termioslib provides access to termios.h.
package termioslib

/*
#include <termios.h>
#include <unistd.h>
*/
import "C"

type Termios struct {
	i	C.struct_termios
    C_iflag, C_oflag, C_cflag, C_lflag C.tcflag_t
    C_cc [NCCS]C.cc_t
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
	//VDSUSP		= uint (C.VDSUSP);
	VSTART		= uint (C.VSTART);
	VSTOP		= uint (C.VSTOP);
	VLNEXT		= uint (C.VLNEXT);
	VDISCARD	= uint (C.VDISCARD);
	VMIN		= uint (C.VMIN);
	VTIME		= uint (C.VTIME);
	//VSTATUS		= uint (C.VSTATUS);
	/* 19 is spare */
	NCCS		= uint (C.NCCS)
)

// Input flags - software input processing
const (
	IGNBRK		= C.tcflag_t(C.IGNBRK)	/* ignore BREAK condition */
	BRKINT		= C.tcflag_t(C.BRKINT) 	/* map BREAK to SIGINTR */
	IGNPAR		= C.tcflag_t(C.IGNPAR) 	/* ignore (discard) parity errors */
	PARMRK		= C.tcflag_t(C.PARMRK) 	/* mark parity and framing errors */
	INPCK		= C.tcflag_t(C.INPCK)	/* enable checking of parity errors */
	ISTRIP		= C.tcflag_t(C.ISTRIP) 	/* strip 8th bit off chars */
	INLCR		= C.tcflag_t(C.INLCR) 	/* map NL into CR */
	IGNCR		= C.tcflag_t(C.IGNCR)	/* ignore CR */
	ICRNL		= C.tcflag_t(C.ICRNL)	/* map CR to NL (ala CRMOD) */
	IXON		= C.tcflag_t(C.IXON)	/* enable output flow control */
	IXOFF		= C.tcflag_t(C.IXOFF) 	/* enable input flow control */
	IXANY		= C.tcflag_t(C.IXANY) 	/* any char will restart after stop */
	IMAXBEL		= C.tcflag_t(C.IMAXBEL) /* ring bell on input queue full */
	IUTF8		= C.tcflag_t(C.IUTF8) 	/* maintain state for UTF-8 VERASE */
)

// Output flags - software output processing
const (
	OPOST		= C.tcflag_t(C.OPOST) 	/* enable following output processing */
	ONLCR		= C.tcflag_t(C.ONLCR) 	/* map NL to CR-NL (ala CRMOD) */
	//OXTABS		= tcflag_t (C.OXTABS); 	/* expand tabs to spaces */
	//ONOEOT		= tcflag_t (C.ONOEOT) 	/* discard EOT's (^D) on output) */
)

// Standard speeds
const (
	B0			= C.speed_t(C.B0)
	B50			= C.speed_t(C.B50)
	B75			= C.speed_t(C.B75)
	B110		= C.speed_t(C.B110)
	B134		= C.speed_t(C.B134)
	B150		= C.speed_t(C.B150)
	B200		= C.speed_t(C.B200)
	B300		= C.speed_t(C.B300)
	B600		= C.speed_t(C.B600)
	B1200		= C.speed_t(C.B1200)
	B1800		= C.speed_t(C.B1800)
	B2400		= C.speed_t(C.B2400)
	B4800		= C.speed_t(C.B4800)
	B9600		= C.speed_t(C.B9600)
	B19200		= C.speed_t(C.B19200)
	B38400		= C.speed_t(C.B38400)
	B57600		= C.speed_t(C.B57600)
	B115200		= C.speed_t(C.B115200)
	B230400		= C.speed_t(C.B230400)
	EXTA		= C.speed_t(C.EXTA)
	EXTB		= C.speed_t(C.EXTB)
)

// Control flags - hardware control of terminal
const (
	//CIGNORE		= tcflag_t (C.CIGNORE);	/* ignore control flags */
	CSIZE		= C.tcflag_t(C.CSIZE)	/* character size mask */
		CS5		= C.tcflag_t(C.CS5)		/* 5 bits (pseudo) */
		CS6		= C.tcflag_t(C.CS6)		/* 6 bits */
		CS7		= C.tcflag_t(C.CS7)		/* 7 bits */
		CS8		= C.tcflag_t(C.CS8)		/* 8 bits */
	CSTOPB		= C.tcflag_t(C.CSTOPB)	/* send 2 stop bits */
	CREAD		= C.tcflag_t(C.CREAD)	/* enable receiver */
	PARENB		= C.tcflag_t(C.PARENB)	/* parity enable */
	PARODD		= C.tcflag_t(C.PARODD)	/* odd parity, else even */
	HUPCL		= C.tcflag_t(C.HUPCL)	/* hang up on last close */
	CLOCAL		= C.tcflag_t(C.CLOCAL)	/* ignore modem status lines */
	//CCTS_OFLOW	= tcflag_t (C.CCTS_OFLOW);	/* CTS flow control of output */
	//CRTS_IFLOW	= tcflag_t (C.CRTS_IFLOW);	/* RTS flow control of input */
	CRTSCTS		= C.tcflag_t(C.CRTSCTS)
	//CDTR_IFLOW	= tcflag_t (C.CDTR_IFLOW);	/* DTR flow control of input */
	//CDSR_OFLOW	= tcflag_t (C.CDSR_OFLOW);	/* DSR flow control of output */
	//CCAR_OFLOW	= tcflag_t (C.CCAR_OFLOW);	/* DCD flow control of output */
	//MDMBUF		= tcflag_t (C.MDMBUF)		/* old name for CCAR_OFLOW */
)

// "Local" flags - dumping ground for other state
const (
	ECHOKE		= C.tcflag_t(C.ECHOKE)	/* visual erase for line kill */
	ECHOE		= C.tcflag_t(C.ECHOE)	/* visually erase chars */
	ECHOK		= C.tcflag_t(C.ECHOK)	/* echo NL after line kill */
	ECHO		= C.tcflag_t(C.ECHO)	/* enable echoing */
	ECHONL		= C.tcflag_t(C.ECHONL)	/* echo NL even if ECHO is off */
	ECHOPRT		= C.tcflag_t(C.ECHOPRT)	/* visual erase mode for hardcopy */
	ECHOCTL		= C.tcflag_t(C.ECHOCTL)	/* echo control chars as ^(Char) */
	ISIG		= C.tcflag_t(C.ISIG)	/* enable signals INTR, QUIT, [D]SUSP */
	ICANON		= C.tcflag_t(C.ICANON)	/* canonicalize input lines */
	//ALTWERASE	= tcflag_t (C.ALTWERASE);	/* use alternate WERASE algorithm */
	IEXTEN		= C.tcflag_t(C.IEXTEN)	/* enable DISCARD and LNEXT */
	EXTPROC		= C.tcflag_t(C.EXTPROC)	/* external processing */
	TOSTOP		= C.tcflag_t(C.TOSTOP)	/* stop background jobs from output */
	FLUSHO		= C.tcflag_t(C.FLUSHO)	/* output being flushed (state) */
	//NOKERNINFO	= tcflag_t (C.NOKERNINFO);	/* no kernel output from VSTATUS */
	PENDIN		= C.tcflag_t(C.PENDIN)	/* XXX retype pending input (state) */
	NOFLSH		= C.tcflag_t(C.NOFLSH)	/* don't flush after interrupt */
)

// Commands passed to tcsetattr() for setting the termios structure.
const (
	TCSANOW		= int (C.TCSANOW)		/* make change immediate */
	TCSADRAIN	= int (C.TCSADRAIN)	/* drain output, then change */
	TCSAFLUSH	= int (C.TCSAFLUSH)	/* drain output, flush input */
)

const (
	TCIFLUSH	= int (C.TCIFLUSH)
	TCOFLUSH	= int (C.TCOFLUSH)
	TCIOFLUSH	= int (C.TCIOFLUSH)
	TCOOFF		= int (C.TCOOFF)
	TCOON		= int (C.TCOON)
	TCIOFF		= int (C.TCIOFF)
	TCION		= int (C.TCION)
)

//speed_t	cfgetispeed(const struct Termios *);
func Getispeed (src * Termios) (result C.speed_t) {
	result = C.cfgetispeed(&(src.i))
	return
}

//speed_t	cfgetospeed(const struct Termios *);
func Getospeed (src * Termios) (result C.speed_t) {
	result = C.cfgetospeed(&src.i)
	return
}

//int	cfsetispeed(struct Termios *, speed_t);
func Setispeed(dst * Termios, baud C.speed_t) (error) {
	dst.i.c_iflag = dst.C_iflag
	dst.i.c_oflag = dst.C_oflag
	dst.i.c_cflag = dst.C_cflag
	dst.i.c_lflag = dst.C_lflag
	dst.i.c_cc = dst.C_cc
	_,rv := C.cfsetispeed(&dst.i, baud)
	if rv == nil {
		dst.C_iflag = dst.i.c_iflag
		dst.C_oflag = dst.i.c_oflag
		dst.C_cflag = dst.i.c_cflag
		dst.C_lflag = dst.i.c_lflag
		dst.C_cc	= dst.i.c_cc
	}
	return rv
}

//int	cfsetospeed(struct Termios *, speed_t);
func Setospeed(dst * Termios, baud C.speed_t) (error) {
	dst.i.c_iflag = dst.C_iflag
	dst.i.c_oflag = dst.C_oflag
	dst.i.c_cflag = dst.C_cflag
	dst.i.c_lflag = dst.C_lflag
	dst.i.c_cc = dst.C_cc
	_,rv := C.cfsetospeed(&dst.i, baud)
	if rv == nil {
		dst.C_iflag = dst.i.c_iflag
		dst.C_oflag = dst.i.c_oflag
		dst.C_cflag = dst.i.c_cflag
		dst.C_lflag = dst.i.c_lflag
		dst.C_cc	= dst.i.c_cc
	}
	return rv
}

func Getattr(fd uintptr, dst * Termios) (error) {
	_,rv := C.tcgetattr(C.int(fd),&dst.i)
	if rv == nil {
		dst.C_iflag = dst.i.c_iflag
		dst.C_oflag = dst.i.c_oflag
		dst.C_cflag = dst.i.c_cflag
		dst.C_lflag = dst.i.c_lflag
		dst.C_cc	= dst.i.c_cc
	}
	return rv
}

//int	tcsetattr(int, int, const struct termios *);
func Setattr(fd uintptr, optional_action int, src * Termios) (error) {
	src.i.c_iflag = src.C_iflag
	src.i.c_oflag = src.C_oflag
	src.i.c_cflag = src.C_cflag
	src.i.c_lflag = src.C_lflag
	src.i.c_cc = src.C_cc
	_,rv := C.tcsetattr(C.int(fd), C.TCSANOW, &src.i)
	return rv
}

//int	tcdrain(int) __DARWIN_ALIAS_C(tcdrain);
func Drain(fd uintptr) (error) {
	_, rv := C.tcdrain(C.int(fd))
	return rv
}

//int	tcflow(int, int);
func Flow(fd uintptr, action int) (error) {
	_,rv := C.tcflow(C.int(fd), C.int(action))
	return rv
}

//int	tcflush(int, int);
func Flush(fd uintptr, queue_selector int) (error) {
	_, rv := C.tcflush(C.int(fd), C.int(queue_selector))
	return rv
}
//int	tcsendbreak(int, int);
func Sendbreak(fd uintptr, duration int) (error) {
	_, rv := C.tcflush(C.int(fd), C.int(duration))
	return rv

}
//void	cfmakeraw(struct termios *);
func Makeraw(src * Termios) {
	C.cfmakeraw(&src.i)
}

// int	cfsetspeed(struct termios *, speed_t);
func Setspeed (dst * Termios, baud C.speed_t) (error) {
	dst.i.c_iflag = dst.C_iflag
	dst.i.c_oflag = dst.C_oflag
	dst.i.c_cflag = dst.C_cflag
	dst.i.c_lflag = dst.C_lflag
	dst.i.c_cc = dst.C_cc
	_, rv := C.cfsetspeed(&dst.i, baud)
	if rv == nil {
		dst.C_iflag = dst.i.c_iflag
		dst.C_oflag = dst.i.c_oflag
		dst.C_cflag = dst.i.c_cflag
		dst.C_lflag = dst.i.c_lflag
		dst.C_cc	= dst.i.c_cc
	}
	return rv
}
