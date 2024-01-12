package ticket_test

import (
	"fmt"
	"testing"

	"github.com/zxdev/ticket"
)

func TestTicket(t *testing.T) {

	tkt, s := ticket.NewTicket(0) // 6e799a80-d1d5-46d0-4c05-f576975c52af74d7
	fmt.Println(tkt, s)

	var tk ticket.Ticket
	fmt.Println(tk.Generate(0))           // f1309576-f967-6c4c-d478-9cba2250e8c537c4
	fmt.Println(tk.Generate(1))           // 00000001-5e28-7168-40b5-c511777825479e0b
	fmt.Println(tk.Generate(2))           // 00000002-5e28-7168-ff0b-fa8f99b1b991704f
	fmt.Println(tk.New())                 // e6c5a2df-760d-1f89-e818-772194e4d917d122
	fmt.Println(tk.String(), tk.Uint64()) // e6c5a2df-760d-1f89-e818-772194e4d917d122 2508812692032332151
	fmt.Println(tk.Generate(0).Uint64())  // 16162510338270547577
}
