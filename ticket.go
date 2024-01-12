package ticket

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

// Ticket type
type Ticket [18]byte

// NewTicket constructor generates a new Ticket
// using the numeric identifer when non-zero
func NewTicket(n uint32) (Ticket, string) {
	var tk Ticket
	return tk, tk.Generate(n).String()
}

// New convienence random ticket generator; wraps ticket.Generate(0)
func (tk *Ticket) New() Ticket { return tk.Generate(0) }

/*
Generate a ticket with random or sequential+timestamp head identifer.

	tk.Generate(0) for all random bytes
	h:0 e4e45937-79c9-c3b4-07e4-7c13d989f9235e15 e4e45937-79c9
	h:n 00000001-5d9b-95d2-de8d-9c7cb21451fac9c1 00000001-5d9b

	[4]byte random or numeric header
	[2]byte random or high 32bit timestamp when n>0
	[2]byte random or low 32bit timestamp when n>0
	[2]byte random uint16
	[8]byte random uint64
*/
func (tk *Ticket) Generate(n uint32) Ticket {
	rand.Read(tk[:]) // random bytes
	if n != 0 {
		// use hex timestamp and overwrite with header h value
		binary.BigEndian.PutUint64(tk[:], uint64(time.Now().Unix()))
		binary.BigEndian.PutUint32(tk[:4], n)
	}
	return *tk
}

// String returns the canonical string representation of the Ticket as string
// ticket: c8e459a9-7979-a384-00e4-9ede014d9c4e0142
func (tk Ticket) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", tk[0:4], tk[4:6], tk[6:8], tk[8:10], tk[10:])
}

/*
Uint64 returns the 64-bit random string tail segment as a uint64 integer.

	ticket: c8e459a9-7979-a384-00e4-9ede014d9c4e0142 -> uint64(9ede014d9c4e0142)
	return: 4316742087152263098
*/
func (tk Ticket) Uint64() uint64 { return binary.LittleEndian.Uint64(tk[10:18]) }

/*
Short returns a shortened version consisting of the left two head blocks.

	h:0 e4e45937-79c9-c3b4-07e4-7c13d989f9235e15 -> e4e45937-79c9
	h:1 00000001-5d9b-95d2-de8d-9c7cb21451fac9c1 -> 00000001-5d9b
*/
func (tk Ticket) Short() string { return fmt.Sprintf("%x-%x", tk[0:4], tk[4:6]) }

/*
Tail returns a shortned version consisting of right two tail blocks.

	h:0 e4e45937-79c9-c3b4-07e4-7c13d989f9235e15 -> 07e4-7c13d989f9235e15
	h:1 00000001-5d9b-95d2-de8d-9c7cb21451fac9c1 -> de8d-9c7cb21451fac9c1
*/
func (tk Ticket) Tail() string { return fmt.Sprintf("%x-%x", tk[7:9], tk[9:]) }
