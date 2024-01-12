# ticket

Generate a random ticket or a ticket with numeric identifer.

	tk.Generate(0) for all random bytes
	h:0 e4e45937-79c9-c3b4-07e4-7c13d989f9235e15 e4e45937-79c9
	h:n 00000001-5d9b-95d2-de8d-9c7cb21451fac9c1 00000001-5d9b

	[4]byte random or numeric header
	[2]byte random or high 32bit timestamp when n>0
	[2]byte random or low 32bit timestamp when n>0
	[2]byte random uint16
	[8]byte random uint64

```golang

// ticket configurator package level
tk, uniq := ticket.NewTicket(0) 
fmt.Println(uniq)
// 6e799a80-d1d5-46d0-4c05-f576975c52af74d7 6e799a80-d1d5-46d0-4c05-f576975c52af74d7

// alternate usage format
var tk ticket.Ticket
fmt.Println(tk.New())                 // e6c5a2df-760d-1f89-e818-772194e4d917d122
fmt.Println(tk.String(), tk.Uint64()) // e6c5a2df-760d-1f89-e818-772194e4d917d122 2508812692032332151

fmt.Println(tk.Generate(0))           // f1309576-f967-6c4c-d478-9cba2250e8c537c4
fmt.Println(tk.Generate(1))           // 00000001-5e28-7168-40b5-c511777825479e0b
fmt.Println(tk.Generate(2))           // 00000002-5e28-7168-ff0b-fa8f99b1b991704f

// generate a .Short() head segment from a ticket or just shorter ticket
tk.Generate(1).Short() // e4e45937-79c9-c3b4-07e4-7c13d989f9235e15 
tk.Short()             // e4e45937-79c9
tk.New().Short()       // e4e45937-79c9

```

