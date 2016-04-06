package main

import (
	"log"
	"net"
	"os/exec"
	"syscall"

	"github.com/tucnak/climax"
)

var sshCommand = climax.Command{
	Name:  "ssh",
	Brief: "runs ssh to a given machine's internal IP",
	Usage: `[flags] <selector>`,

	Flags: []climax.Flag{
		{
			Name:     "identity",
			Short:    "i",
			Usage:    `-i <key>`,
			Help:     `specifies the ssh key that should be used to connect`,
			Variable: true,
		},
		{
			Name:     "user",
			Short:    "u",
			Usage:    `-u <user>`,
			Help:     `specifies the username to login with`,
			Variable: true,
		},
	},

	Examples: []climax.Example{
		{
			Usecase:     `"minion.*"`,
			Description: `Will SSH to a minion if one matches the selector.`,
		},
	},

	Handle: truckSSH,
}

func isLocalAddress(ip string) bool {
	addr := net.ParseIP(ip)
	networks := []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}

	for _, network := range networks {
		_, cidrnet, err := net.ParseCIDR(network)
		if err != nil {
			panic(err)
		}
		if cidrnet.Contains(addr) {
			return true
		}
	}

	return false
}

// Returns either the first local address, or the first address.
func getBestAddress(ips []string) string {
	for _, ip := range ips {
		if isLocalAddress(ip) {
			return ip
		}
	}
	return ips[0]
}

func selectMinion(minions map[string][]string) []string {
	if len(minions) == 0 {
		return []string{}
	} else if len(minions) == 1 {
		for _, minion := range minions {
			return minion
		}
	}

	options := []string{}
	for id := range minions {
		options = append(options, id)
	}

	menu := ListMenu{Title: "Select Minion", Options: options}
	option := menu.Show()

	return minions[option]
}

func truckSSH(ctx climax.Context) int {
	sshopts := []string{"ssh"}
	minions := map[string][]string{}

	if len(ctx.Args) == 0 {
		log.Println("no selector specified")
		return 1
	}

	err := salt(&minions, ctx.Args[0], "network.ip_addrs")
	if err != nil {
		log.Println("couldn't get IPs:", err.Error())
		return 1
	}

	if len(minions) == 0 {
		log.Println("no minions matched")
		return 0
	}

	// Figure out what IP address to use.
	minion := selectMinion(minions)
	ip := getBestAddress(minion)

	// Identity option.
	if identity, ok := ctx.Get("identity"); ok {
		sshopts = append(sshopts, "-i", identity)
	}

	// Form hostname.
	hostname := ip
	if user, ok := ctx.Get("user"); ok {
		hostname = user + "@" + ip
	}

	sshopts = append(sshopts, hostname)

	// Add arguments from our invocation.
	sshopts = append(sshopts, ctx.Args[1:]...)

	// Find the SSH command.
	sshPath, err := exec.LookPath("ssh")
	if err != nil {
		log.Println("couldn't find ssh path:", err)
	}

	// Go!
	err = syscall.Exec(sshPath, sshopts, []string{})
	if err != nil {
		log.Println("couldn't fork to ssh:", err)
	}

	return 0
}
