package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Player struct {
	Name string
	Conn net.Conn
	RW   *bufio.ReadWriter
}

// TODO: Write unit tests for Player that use a mock net.Listener
func AcceptPlayer(listen net.Listener) (p *Player, err error) {
	p = new(Player)
	p.Conn, err = listen.Accept()
	if err != nil {
		return nil, err
	}
	p.RW = bufio.NewReadWriter(bufio.NewReader(p.Conn), bufio.NewWriter(p.Conn))

	for p.Name == "" {
		log.Printf("Asking %v for their name", p)
		prompt := "Hi! What is your name?\n"
		reply, err := p.exchange(prompt)
		if err != nil {
			return nil, err
		}
		p.Name = reply
	}

	log.Printf("%v is %v", p.Conn.RemoteAddr().String(), p)
	return p, nil
}

// Returns once the player has successfully repeated the phrase back to us.
func (p *Player) Race(phrase string) error {
	// TODO: Instead of doing line-based buffering, be interactive.
	// Maybe show others progress at the same time? (Might need a client
	// to play instead of just telnet.)
	// TODO: Some other game mechanism entirely?
	// TODO: Cheating prevention: Have them enter the string reversed? In all lower case?
	// Detect the rate at which each character gets entered? (Requires interactivity.)
	// Do something else clever?

	var reply string
	var err error
	for reply != phrase {
		log.Printf("Sending phrase %q to %v", phrase, p)
		prompt := fmt.Sprintf("Type this back *perfectly* as fast you can:\n%s\n", phrase)
		reply, err = p.exchange(prompt)
		if err != nil {
			return err
		}
		log.Printf("Reply from %s: %q", p, reply)
	}

	log.Printf("Success! Well done, %v", p)
	return nil
}

func (p *Player) exchange(prompt string) (string, error) {
	// send prompt
	if _, err := p.RW.WriteString(prompt); err != nil {
		return "", err
	}
	if err := p.RW.Flush(); err != nil {
		return "", err
	}

	// read and clean up reply
	raw, err := p.RW.ReadString('\n')
	if err != nil {
		return "", err
	}

	reply := strings.TrimSpace(raw)
	return reply, nil
}

func (p *Player) String() string {
	switch {
	case p.Name != "":
		return p.Name
	case p.Conn != nil:
		return p.Conn.RemoteAddr().String()
	default:
		return "unknown player"
	}
}
