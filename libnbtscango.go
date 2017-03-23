package libnbtscango

import (
	"os"
	"fmt"
	"byteexec"
)
/* example result
# nbtscan -v 192.168.30.10
Doing NBT name scan for addresses from 192.168.30.10


NetBIOS Name Table for Host 192.168.30.10:

Incomplete packet, 155 bytes long.
Name             Service          Type
----------------------------------------
WIN2K8-1022       <00>             UNIQUE
WORKGROUP         <00>              GROUP
WIN2K8-1022       <20>             UNIQUE

Adapter address: 00-50-56-bd-57-4e
----------------------------------------
*/

type struct Nbt {
	host string
	group string
}

var nbtscanBin []byte

func init() {
	var err error
	nbtscanBin, err = Asset("data/nbtscan")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func scanByIP(ip string) (nbt Nbt, err error) {
	be, err := byteexec.New(nbtscanBin)
	if err != nil {
		log.Fatalf("Uh oh: %s", err)
	}
	cmd := be.Command("-v", ip)
	// cmd is an os/exec.Cmd
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
	nbt = Nbt{host: "abc", group: "group"}
	return
}
