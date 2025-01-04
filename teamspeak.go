package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thannaske/go-ts3"
)

const (
	VERSION = "1.0.0"
)

func main() {
	serverPtr := flag.String("server", "127.0.0.1:10011", "IP address or hostname of Teamspeak 3 server query (default: 127.0.0.1)")
	usernamePtr := flag.String("username", "serveradmin", "Teamspeak 3 server query user (default: serveradmin)")
	passwordPtr := flag.String("password", "", "Teamspeak 3 server query password (default: none)")
	serverIDPtr := flag.Int("id", 1, "Server ID (default: 1)")

	flag.Parse()

	c, err := ts3.NewClient(*serverPtr)

	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Could not establish server query connection to", *serverPtr)
		os.Exit(1)
	}

	defer c.Close()

	if err := c.Login(*usernamePtr, *passwordPtr); err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Authentication failure")
		os.Exit(1)
	}

	if err := c.Use(*serverIDPtr); err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Could not select Teamspeak 3 server instance by ID")
		os.Exit(1)
	}

	server, err := c.Server.Info()
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Could not get server info")
		os.Exit(1)
	}

	var serverAutoStart int
	if server.AutoStart {
		serverAutoStart = 1
	} else {
		serverAutoStart = 0
	}

	var voice_clients = server.ClientsOnline - server.QueryClientsOnline
	var serverStatus int

	if server.Status == "online" {
		serverStatus = 1
	} else {
		serverStatus = 2
	}

	var format string = "teamspeak_server,port=%d,id=%d online=%d,v_clients=%d,q_clients=%d,m_clients=%d,autostart=%d,bytes_out=%d,bytes_in=%d,channels=%d,reserved_slots=%d,uptime=%d,packets_in=%d,packets_out=%d,ft_bytes_in_total=%d,ft_bytes_out_total=%d,pl_control=%f,pl_speech=%f,pl_keepalive=%f,pl_total=%f,bytes_out_speech=%d,bytes_in_speech=%d,bytes_out_control=%d,bytes_in_control=%d,bytes_out_keepalive=%d,bytes_in_keepalive=%d,packets_out_speech=%d,packets_in_speech=%d,packets_out_control=%d,packets_in_control=%d,packets_keepalive_out=%d,packets_keepalive_in=%d,avg_ping=%f\n"
	fmt.Printf(format, server.Port, server.ID, serverStatus, voice_clients, server.QueryClientsOnline, server.MaxClients, serverAutoStart, server.BytesSentTotal, server.BytesReceivedTotal, server.ChannelsOnline, server.ReservedSlots, server.Uptime, server.PacketsReceivedTotal, server.PacketsSentTotal, server.TotalBytesUploaded, server.TotalBytesDownloaded, server.TotalPacketLossControl, server.TotalPacketLossSpeech, server.TotalPacketLossKeepalive, server.TotalPacketLossTotal, server.SpeechBytesSent, server.SpeechBytesReceived, server.ControlBytesSent, server.ControlBytesReceived, server.KeepaliveBytesSent, server.KeepaliveBytesReceived, server.SpeechPacketsSent, server.SpeechPacketsReceived, server.ControlPacketsSent, server.ControlPacketsReceived, server.KeepalivePacketsSent, server.KeepalivePacketsReceived, server.TotalPing)
}
