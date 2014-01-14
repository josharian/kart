package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/josharian/kart/phrase"
)

var (
	corpus = flag.String("corpus", "", "corpus to select phrases from; if not provided, phrases will be randomly generated")
	port   = flag.Int("port", 11235, "port to listen on; defaults to 11235")
	seed   = flag.Int64("seed", time.Now().UTC().UnixNano(), "random seed; defaults to current ns timestamp")
	shout  = flag.Bool("shout", false, "SHOUT ALL THE PHRASES")
)

func main() {
	flag.Parse()
	rand.Seed(*seed)

	var src phrase.Source
	if *corpus != "" {
		f, err := os.Open(*corpus)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		src, err = phrase.NewCorpus(f)
	} else {
		src = &phrase.Rand{Chars: "abcdef", Length: 9}
	}

	if *shout {
		src = phrase.Shout(src)
	}

	src = phrase.Clean(src)
	// TODO: More phrase fun?

	ip, err := guessLocalIp4()
	if err != nil {
		log.Fatalf("Failed to guess local ipv4 addr: %v", err)
	}
	// log.Printf("Going to listen on %v:%v", ip, Port)

	// TODO: Play over socket, UDP, or other net.Listener
	listen, err := net.ListenTCP("tcp4", &net.TCPAddr{Port: *port})
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("To play: telnet %v %v", ip, *port)

	// TODO: Accept arbitrarily many players; use a timeout + select?
	// TODO: Be more robust here; when a player's connection causes an error,
	// don't stop the game, just drop that player from the running.
	// TODO: Instead of having a server, play peer-to-peer!

	var pp []*Player
	playerc := make(chan *Player)

	for i := 0; i < 2; i++ {
		go func() {
			p, err := AcceptPlayer(listen)
			if err != nil {
				log.Fatalf("Failed to accept player: %v", err)
			}
			playerc <- p
		}()
	}

	for i := 0; i < 2; i++ {
		p := <-playerc
		pp = append(pp, p)
	}

	w := src.Phrase()
	log.Printf("Time to play! Sending the players (%v) the phrase: %s", pp, w)
	winnerc := make(chan *Player)

	for _, p := range pp {
		go func(p *Player) {
			if err := p.Race(w); err != nil {
				log.Fatalf("Error from player %v during race: %v", p, err)
			}
			winnerc <- p
		}(p)
	}

	p := <-winnerc
	log.Printf("We have a winner! Congrats, %v!", p.Name)

	// TODO: Send consolation messages to everyone else, instead of
	// just hanging up on them.
	// TODO: Start a new game instead of exiting. We'll need to make
	// sure we close the old connections.
}

// guessLocalIp4 tries to deduce the local ipv4 address.
// The goal of using it instead of just ifconfig is to ward
// off the wrath of the demo gods by keeping everything in
// one place.
func guessLocalIp4() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			// not ip
			continue
		}
		ip := ipnet.IP
		if ip.To4() == nil {
			// not ipv4
			continue
		}
		if ip.IsLoopback() {
			// loopback
			continue
		}
		return ip, nil
	}
	return nil, fmt.Errorf("found no IP addresses; searched %v", addrs)
}
