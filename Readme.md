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
	    "os"
	    "errors"
	    "fmt"
	    "termioslib"
	)
	
	
	func main () {
	    var (
			err error
			orig_termios termioslib.Termios
			work_termios termioslib.Termios
			serial_port *os.File
	    )

	    defer func () {
	        if err != nil { fmt.Println (err) }
	    } ();
	    
	    // open the serial port, other options of interest may be os.NDELAY and os.NOCTTY
		serial_port, err = os.Open("/dev/cu.usbserial-FTCW1RJN", os.O_RDWR , 777)
		if err != nil { return }
		
		defer func () {
			serial_port.Close()
		}();
		
		// flush all buffers
		if err = termioslib.Flush(serial_port.Fd(), termioslib.TCIOFLUSH); err != nil { return }
		
		// save a copy of the original terminal configuration
	    if err = termioslib.Getattr(serial_port.Fd(), &orig_termios); err != nil { return }
	    
	    // get a working copy
		if err = termioslib.Getattr(serial_port.Fd(), &work_termios); err != nil { return }
		
		work_termios.C_iflag = (termioslib.IGNPAR)
		work_termios.C_oflag = 0
	    work_termios.C_cflag = (termioslib.CRTSCTS | termioslib.CS8 | termioslib.CLOCAL | termioslib.CREAD)
	    work_termios.C_lflag = 0
		work_termios.C_cc[termioslib.VMIN] = 1
		work_termios.C_cc[termioslib.VTIME] = 0
		
		termioslib.Setspeed(&work_termios, termioslib.B57600)
		// set the working copy
		if err = termioslib.Setattr(serial_port.Fd(), termioslib.TCSANOW , &work_termios); err != nil { return }
		
		// set the settings back to the original when the program exits
		defer func () {
			err = termioslib.Setattr(serial_port.Fd(), termioslib.TCSANOW, &orig_termios)
	    } ();
	    
	    // basic file read block
		buffer := make([]byte, 80) 
		n, err := serial_port.Read(buffer)
	
		for err == nil {
			fmt.Printf("%s", string(buffer[0:n]))
			n, err = serial_port.Read(buffer)
		}
	}

## Future

Once the differences between Linux and Mac termios.h have been identified, the architecture dependent defines and functions will be pushed into separate files as directed in the last section of [How to write go code](http://golang.org/doc/code.html).

## About 

termioslib was written by Whitham D. Reeve II
