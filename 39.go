package main

type contact struct {
	userID       string  // 16 bytes - largest first
	sendingLimit int32   // 4 bytes
	age          int32   // 4 bytes
	// Total: 24 bytes (16 + 4 + 4 = 24, no padding needed)
}

type perms struct {
	permissionLevel int  // 8 bytes - largest first
	canSend         bool // 1 byte
	canReceive      bool // 1 byte
	canManage       bool // 1 byte
	// Total: 11 bytes (8 + 1 + 1 + 1 = 11) 
	// But will be padded to 16 bytes for alignment
}
