# termioslib

termioslib is a simple cgo based wrapper around termios.h

## Overview

termioslib is a simple no-nonsense wrapper for UNIX termios.h. It exposes all the defines and functions in termios.h including some extensions not defined by POSIX. You might need to remove or rename the additional defines/functions for your specific system. This was developed on Mac OS X. I welcome feedback on whether or not modifications need to be made to run on other Unix/Linux.

## Installation

1. Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html). termios targets the 'release' tag. To get the release tag, simply run `hg update -r release`.
2. Make sure that $GOROOT is set. termios installs itself as a Go package, so it requires $GOROOT. See Go's install instructions for more information about $GOROOT. 
2. git clone git://github.com/wdreeveii/termioslib.git
3. cd termioslib && make install

## Example

	package main
	
	import (
	    "fmt";
	    "os";
	    "termioslib";
	)
	
	
	func main () {
	    var (
			err error
			orig_termios termioslib.Termios
			work_termios termioslib.Termios
			ser *os.File
	    )

	    defer func () {
	        if err != nil { fmt.Println (err) }
	    } ()
	    
	    // open the serial port, other options of interest may be os.NDELAY and os.NOCTTY
	    ser, err = os.Open("/dev/cu.usbserial-FTCW1RJN", os.O_RDWR , 777)
	    if err != nil { return }
		
	    defer func () {
		ser.Close()
	    }()
		
	    // flush all buffers
	    if err = termioslib.Flush(ser.Fd(), termioslib.TCIOFLUSH); err != nil { return }
		
	    // save a copy of the original terminal configuration
	    if err = termioslib.Getattr(ser.Fd(), &orig_termios); err != nil { return }
	    
	    // get a working copy
	    if err = termioslib.Getattr(ser.Fd(), &work_termios); err != nil { return }
	    
	    work_termios.C_iflag &= ^(termioslib.IGNBRK | termioslib.BRKINT | termioslib.IGNPAR | termioslib.PARMRK | termioslib.INPCK | termioslib.ISTRIP | termioslib.INLCR | termioslib.IGNCR | termioslib.ICRNL | termioslib.IXON | termioslib.IXOFF | termioslib.IXANY | termioslib.IMAXBEL | termioslib.IUTF8)
	    work_termios.C_oflag &= ^(termioslib.OPOST | termioslib.ONLCR )
	    work_termios.C_lflag &= ^(termioslib.ISIG | termioslib.ICANON | termioslib.IEXTEN | termioslib.ECHO | termioslib.ECHOE | termioslib.ECHOK | termioslib.ECHONL | termioslib.NOFLSH | termioslib.TOSTOP | termioslib.ECHOPRT | termioslib.ECHOCTL | termioslib.ECHOKE)
	    work_termios.C_cflag &= ^(termioslib.CSIZE | termioslib.PARENB | termioslib.PARODD | termioslib.HUPCL | termioslib.CSTOPB | termioslib.CRTSCTS)
	    work_termios.C_cflag |= (termioslib.CS8 | termioslib.CREAD | termioslib.CLOCAL)
	    work_termios.C_cc[termioslib.VMIN] = 1
	    work_termios.C_cc[termioslib.VTIME] = 0
	    
	    termioslib.Setspeed(&work_termios, termioslib.B57600)
	    // set the working copy
	    if err = termioslib.Setattr(ser.Fd(), termioslib.TCSANOW , &work_termios); err != nil { return }
	    
	    // set the settings back to the original when the program exits
	    defer func () {
		err = termioslib.Setattr(ser.Fd(), termioslib.TCSANOW, &orig_termios)
	    } ()
	    
	    // basic file read block
	    buffer := make([]byte, 80) 
	    n, err := ser.Read(buffer)
	
	    for err == nil {
		    fmt.Printf("%s", string(buffer[0:n]))
		    n, err = ser.Read(buffer)
	    }
	}

## Future

At one time or another, this code has worked on both Linux and Mac OS X. It was updated in April 2013 to go 1.0

## About 

termioslib was written by Whitham D. Reeve II
