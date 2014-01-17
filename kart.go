// Kart is a speed-typing game. It is meant to demo
// a few interesting aspects of Go.
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
)

func main() {
	flag.Parse()
	rand.Seed(*seed)

	var src phrase.Source
	if *corpus != "" {
		src = loadCorpus(*corpus)
	} else {
		src = &phrase.Rand{Chars: "abcdef", Length: 9}
	}

	src = phrase.Clean(src)
	src = phrase.Truncate(src, 15) // keep phrases from getting too long

	ip, err := guessLocalIp4()
	if err != nil {
		log.Fatalf("Failed to guess local ipv4 addr: %v", err)
	}
	// log.Printf("Going to listen on %v:%v", ip, Port)

	listen, err := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.IPv4zero, Port: *port})
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("To play: telnet %v %v", ip, *port)

	var pp []*Player
	playerc := make(chan *Player)

	// Listen for players
	for i := 0; i < 2; i++ {
		go func() {
			p, err := AcceptPlayer(listen)
			if err != nil {
				log.Fatalf("Failed to accept player: %v", err)
			}
			playerc <- p
		}()
	}

	// Wait until we have two players
	for i := 0; i < 2; i++ {
		p := <-playerc
		pp = append(pp, p)
	}

	// Select a phrase to send
	w := src.Phrase()
	log.Printf("Time to play! Sending the players (%v) the phrase: %s", pp, w)

	// Send the phrase to all the players
	winnerc := make(chan *Player)
	for _, p := range pp {
		go func(p *Player) {
			if err := p.Race(w); err != nil {
				log.Fatalf("Error from player %v during race: %v", p, err)
			}
			winnerc <- p
		}(p)
	}

	// The first one to reply wins
	p := <-winnerc
	log.Printf("We have a winner! Congrats, %v!", p.Name)
}

// loadCorpus loads a corpus from the file at path;
// in case of error, loadCorpus calls log.Fatal.
func loadCorpus(path string) phrase.Corpus {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	src, err := phrase.NewCorpus(f)
	if err != nil {
		log.Fatalln(err)
	}
	return src
}

// guessLocalIp4 tries to deduce the local ipv4 address.
// We could just use ifconfig, but that's inelegant, and
// this isn't hard. Also, network config changes is one
// of the tricks used by jealous demo gods to wreak
// vengeance on unsuspecting presenters.
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
