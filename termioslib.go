// Copyright 2010 Whitham D. Reeve II. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package termioslib provides POSIX compatible access to termios.h.
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

// Subscript names for the array c_cc
const (
	VEOF 		= uint (C.VEOF)		/* EOF Character */
	VEOL 		= uint (C.VEOL)		/* EOL character */
	VERASE		= uint (C.VERASE)	/* ERASE character */
	VINTR		= uint (C.VINTR)	/* INTR character */
	VKILL 		= uint (C.VKILL)	/* KILL character */
	VQUIT		= uint (C.VQUIT)	/* QUIT character */
	VSTART		= uint (C.VSTART)	/* START character */
	VSTOP		= uint (C.VSTOP)	/* STOP character */
	VSUSP		= uint (C.VSUSP)	/* SUSP character */
	VMIN		= uint (C.VMIN)		/* MIN value */
	VTIME		= uint (C.VTIME)	/* TIME value */
	NCCS		= uint (C.NCCS)		/* Size of the array c_cc for control characters */
	//VEOL2 	= uint (C.VEOL2) not-posix
	//VWERASE 	= uint (C.VWERASE) not-posix
	//VREPRINT	= uint (C.VREPRINT); not-posix
	//VDSUSP	= uint (C.VDSUSP); not-posix
	//VLNEXT	= uint (C.VLNEXT); not-posix
	//VDISCARD	= uint (C.VDISCARD); not-posix
	//VSTATUS	= uint (C.VSTATUS); not-posix
)

// Input modes - software input processing
const (
	BRKINT		= C.tcflag_t(C.BRKINT) 	/* map BREAK to SIGINTR */
	ICRNL		= C.tcflag_t(C.ICRNL)	/* map CR to NL (ala CRMOD) */
	IGNBRK		= C.tcflag_t(C.IGNBRK)	/* ignore BREAK condition */
	IGNCR		= C.tcflag_t(C.IGNCR)	/* ignore CR */
	IGNPAR		= C.tcflag_t(C.IGNPAR) 	/* ignore (discard) parity errors */
	INLCR		= C.tcflag_t(C.INLCR) 	/* map NL into CR */
	INPCK		= C.tcflag_t(C.INPCK)	/* enable checking of parity errors */
	ISTRIP		= C.tcflag_t(C.ISTRIP) 	/* strip 8th bit off chars */
	IXANY		= C.tcflag_t(C.IXANY) 	/* any char will restart after stop */
	IXOFF		= C.tcflag_t(C.IXOFF) 	/* enable input flow control */
	IXON		= C.tcflag_t(C.IXON)	/* enable output flow control */
	PARMRK		= C.tcflag_t(C.PARMRK) 	/* mark parity and framing errors */
	//IMAXBEL		= C.tcflag_t(C.IMAXBEL) /* ring bell on input queue full */
	//IUTF8		= C.tcflag_t(C.IUTF8) 	/* maintain state for UTF-8 VERASE */
)

// Output modes - software output processing
const (
	OPOST		= C.tcflag_t(C.OPOST) 	/* enable following output processing */
	ONLCR		= C.tcflag_t(C.ONLCR) 	/* map NL to CR-NL (ala CRMOD) */
	OCRNL		= C.tcflag_t(C.OCRNL)	/* Map CR to NL on output */
	ONOCR		= C.tcflag_t(C.ONOCR)	/* No CR output at column 0 */
	ONLRET		= C.tcflag_t(C.ONLRET)	/* NL performs CR function */
	OFDEL		= C.tcflag_t(C.OFDEL)	/* Fill is DEL */
	OFILL		= C.tcflag_t(C.OFILL)	/* Use fill characters for delay */
	NLDLY		= C.tcflag_t(C.NLDLY)	/* Select newline delays: */
	NL0			= C.tcflag_t(C.NL0)		/* Newline type 0 */
	NL1			= C.tcflag_t(C.NL1)		/* Newline type 1 */
	
	CRDLY		= C.tcflag_t(C.CRDLY)	/* Select carriage-return delays: */
	CR0			= C.tcflag_t(C.CR0)		/* Carriage-return delay type 0 */
	CR1			= C.tcflag_t(C.CR1)		/* Carriage-return delay type 1 */
	CR2			= C.tcflag_t(C.CR2)		/* Carriage-return delay type 2 */
	CR3			= C.tcflag_t(C.CR3)		/* Carriage-return delay type 3 */
	
	TABDLY		= C.tcflag_t(C.TABDLY)	/* Select horizontal-tab delays: */
	TAB0		= C.tcflag_t(C.TAB0)	/* Horizontal-tab delay type 0 */
	TAB1		= C.tcflag_t(C.TAB1)	/* Horizontal-tab delay type 1 */
	TAB2		= C.tcflag_t(C.TAB2)	/* Horizontal-tab delay type 2 */
	TAB3		= C.tcflag_t(C.TAB3)	/* Horizontal-tab delay type 3 */
	
	BSDLY		= C.tcflag_t(C.BSDLY)	/* Select backspace delays: */
	BS0			= C.tcflag_t(C.BS0)		/* Backspace-delay type 0 */
	BS1			= C.tcflag_t(C.BS1)		/* Backspace-delay type 0 */
	
	VTDLY		= C.tcflag_t(C.VTDLY)	/* Select vertical-tab delays: */
	VT0			= C.tcflag_t(C.VT0)		/* Vertical-tab delay type 0 */
	VT1			= C.tcflag_t(C.VT1)		/* Vertical-tab delay type 1 */
	
	FFDLY		= C.tcflag_t(C.FFDLY)	/* Select form-feed delays: */
	FF0			= C.tcflag_t(C.FF0)		/* Form-feed delay type 0 */
	FF1			= C.tcflag_t(C.FF1)		/* Form-feed delay type 1 */
)

// Standard speeds
const (
	B0			= C.speed_t(C.B0)		/* Hang up */
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
	//EXTA		= C.speed_t(C.EXTA)
	//EXTB		= C.speed_t(C.EXTB)
)

// Control modes - hardware control of terminal
const (
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
	//CIGNORE		= tcflag_t (C.CIGNORE);	/* ignore control flags */
	//CCTS_OFLOW	= tcflag_t (C.CCTS_OFLOW);	/* CTS flow control of output */
	//CRTS_IFLOW	= tcflag_t (C.CRTS_IFLOW);	/* RTS flow control of input */
	//CRTSCTS		= C.tcflag_t(C.CRTSCTS)
	//CDTR_IFLOW	= tcflag_t (C.CDTR_IFLOW);	/* DTR flow control of input */
	//CDSR_OFLOW	= tcflag_t (C.CDSR_OFLOW);	/* DSR flow control of output */
	//CCAR_OFLOW	= tcflag_t (C.CCAR_OFLOW);	/* DCD flow control of output */
	//MDMBUF		= tcflag_t (C.MDMBUF)		/* old name for CCAR_OFLOW */
)

// Local modes - dumping ground for other state
const (
	ECHO		= C.tcflag_t(C.ECHO)	/* enable echoing */
	ECHOE		= C.tcflag_t(C.ECHOE)	/* visually erase chars */
	ECHOK		= C.tcflag_t(C.ECHOK)	/* echo NL after line kill */
	ECHONL		= C.tcflag_t(C.ECHONL)	/* echo NL even if ECHO is off */
	ICANON		= C.tcflag_t(C.ICANON)	/* canonicalize input lines */
	IEXTEN		= C.tcflag_t(C.IEXTEN)	/* enable DISCARD and LNEXT */
	ISIG		= C.tcflag_t(C.ISIG)	/* enable signals INTR, QUIT, [D]SUSP */
	NOFLSH		= C.tcflag_t(C.NOFLSH)	/* don't flush after interrupt */
	TOSTOP		= C.tcflag_t(C.TOSTOP)	/* stop background jobs from output */
	//EXTPROC		= C.tcflag_t(C.EXTPROC)	/* external processing */
	//FLUSHO		= C.tcflag_t(C.FLUSHO)	/* output being flushed (state) */
	//PENDIN		= C.tcflag_t(C.PENDIN)	/* XXX retype pending input (state) */
	//ECHOPRT		= C.tcflag_t(C.ECHOPRT)	/* visual erase mode for hardcopy */
	//ECHOCTL		= C.tcflag_t(C.ECHOCTL)	/* echo control chars as ^(Char) */
	//ECHOKE		= C.tcflag_t(C.ECHOKE)	/* visual erase for line kill */
	//ALTWERASE	= tcflag_t (C.ALTWERASE);	/* use alternate WERASE algorithm */
	//NOKERNINFO	= tcflag_t (C.NOKERNINFO);	/* no kernel output from VSTATUS */
)

// Commands passed to tcsetattr() for setting the termios structure.
const (
	TCSANOW		= int (C.TCSANOW)		/* make change immediate */
	TCSADRAIN	= int (C.TCSADRAIN)	/* drain output, then change */
	TCSAFLUSH	= int (C.TCSAFLUSH)	/* drain output, flush input */
)

// Commands passed to tcflush() and tcflow()
const (
	TCIFLUSH	= int (C.TCIFLUSH)
	TCOFLUSH	= int (C.TCOFLUSH)
	TCIOFLUSH	= int (C.TCIOFLUSH)
	TCOOFF		= int (C.TCOOFF)
	TCOON		= int (C.TCOON)
	TCIOFF		= int (C.TCIOFF)
	TCION		= int (C.TCION)
)

// Getispeed returns the input baud rate in the Termios
// structure referenced by src.
func Getispeed (src * Termios) (result C.speed_t) {
	result = C.cfgetispeed(&src.i)
	return
}

// Getospeed returns the output baud rate in the Termios
// structure referenced by src.
func Getospeed (src * Termios) (result C.speed_t) {
	result = C.cfgetospeed(&src.i)
	return
}

// Setispeed sets the input baud rate in the Termios
// structure referenced by dst to speed.
func Setispeed(dst * Termios, speed C.speed_t) (error) {
	dst.i.c_iflag = dst.C_iflag
	dst.i.c_oflag = dst.C_oflag
	dst.i.c_cflag = dst.C_cflag
	dst.i.c_lflag = dst.C_lflag
	dst.i.c_cc = dst.C_cc
	_,rv := C.cfsetispeed(&dst.i, speed)
	if rv == nil {
		dst.C_iflag = dst.i.c_iflag
		dst.C_oflag = dst.i.c_oflag
		dst.C_cflag = dst.i.c_cflag
		dst.C_lflag = dst.i.c_lflag
		dst.C_cc	= dst.i.c_cc
	}
	return rv
}

// Setospeed sets the input baud rate in the Termios
// structure referenced by dst to speed.
func Setospeed(dst * Termios, speed C.speed_t) (error) {
	dst.i.c_iflag = dst.C_iflag
	dst.i.c_oflag = dst.C_oflag
	dst.i.c_cflag = dst.C_cflag
	dst.i.c_lflag = dst.C_lflag
	dst.i.c_cc = dst.C_cc
	_,rv := C.cfsetospeed(&dst.i, speed)
	if rv == nil {
		dst.C_iflag = dst.i.c_iflag
		dst.C_oflag = dst.i.c_oflag
		dst.C_cflag = dst.i.c_cflag
		dst.C_lflag = dst.i.c_lflag
		dst.C_cc	= dst.i.c_cc
	}
	return rv
}

// Getattr copies the parameters associated with the terminal
// referenced by fd in the Termios structure referenced by dst.
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

// Setattr sets the parameters associated with the terminal
// from the Termios structure referenced by src.  The optional_actions
// field is created by or'ing the desired tcsetattr commands.
func Setattr(fd uintptr, optional_action int, src * Termios) (error) {
	src.i.c_iflag = src.C_iflag
	src.i.c_oflag = src.C_oflag
	src.i.c_cflag = src.C_cflag
	src.i.c_lflag = src.C_lflag
	src.i.c_cc = src.C_cc
	_,rv := C.tcsetattr(C.int(fd), C.int(optional_action), &src.i)
	return rv
}

// Drain waits until all output written to the terminal
// referenced by fd has been transmitted to the terminal.
func Drain(fd uintptr) (error) {
	//int	tcdrain(int);
	_, rv := C.tcdrain(C.int(fd))
	return rv
}

// Flow suspends transmission of data to, or the reception of data from,
// the terminal referenced by fd, depending on the value of action.
func Flow(fd uintptr, action int) (error) {
	//int	tcflow(int, int);
	_,rv := C.tcflow(C.int(fd), C.int(action))
	return rv
}

// Flush discards any data written to the terminal referenced by fd
// which has not been transmitted to the terminal, or any data received
// from the terminal but not yet read, depending on the value of action.
func Flush(fd uintptr, queue_selector int) (error) {
	//int	tcflush(int, int);
	_, rv := C.tcflush(C.int(fd), C.int(queue_selector))
	return rv
}

// Sendbreak transmits a continuous stream of zero-valued
// bits for four-tenths of a second to the terminal referenced by fildes.
// The duration parameter is ignored in this implementation.
func Sendbreak(fd uintptr, duration int) (error) {
	//int	tcsendbreak(int, int);
	_, rv := C.tcflush(C.int(fd), C.int(duration))
	return rv

}

/* More non-posix functions. 
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
*/
