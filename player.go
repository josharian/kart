package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Player handles all interactions with a player.
type Player struct {
	Name string            // player's name, self-reported
	conn net.Conn          // connection we're talking to the player over
	rw   *bufio.ReadWriter // buffered io atop conn
}

// AcceptPlayer waits for a player to connect to
// listen, and asks the player for their name.
func AcceptPlayer(listen net.Listener) (p *Player, err error) {
	p = new(Player)
	p.conn, err = listen.Accept()
	if err != nil {
		return nil, err
	}
	p.rw = bufio.NewReadWriter(bufio.NewReader(p.conn), bufio.NewWriter(p.conn))

	for p.Name == "" {
		log.Printf("Asking %v for their name", p)
		prompt := "Hi! What is your name?\n"
		reply, err := p.exchange(prompt)
		if err != nil {
			return nil, err
		}
		p.Name = reply
	}

	log.Printf("%v's name is %v", p.conn.RemoteAddr().String(), p)
	return p, nil
}

// Race gives phrase to the player and blocks until the
// player successfully repeats that phrase back to us.
//
// MORE FUN: Instead of doing line-based buffering, do
// character-by-character interaction. This would probably need
// a client to play, rather than telnet. Once this is in place,
// maybe show everyone everyone else's progress.
//
// MORE FUN: Some other game mechanism entirely?
//
// MORE FUN: Cheating prevention. Put in some heuristics
// to detect that it is a human playing instead of a computer.
func (p *Player) Race(phrase string) error {
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

// exchange sends prompt to the player and returns their response,
// with any leading/trailing whitespace trimmed.
func (p *Player) exchange(prompt string) (string, error) {
	// send prompt
	if _, err := p.rw.WriteString(prompt); err != nil {
		return "", err
	}
	if err := p.rw.Flush(); err != nil {
		return "", err
	}

	// read and clean up reply
	raw, err := p.rw.ReadString('\n')
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
	case p.conn != nil:
		return p.conn.RemoteAddr().String()
	default:
		return "unknown player"
	}
}
