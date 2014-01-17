# More fun!

Kart is pretty crude and skeletal at the moment. There are all sorts of improvements just waiting to be made--that's the idea. If you want to hack on it, and you need some ideas, here are a few.

## Small

- Add a Notify function to Player, and use it to congratulate the winner and taunt the losers.
- Instead of closing the connection right away once someone has typed the phrase, wait for all players to finish, then send out a message to everyone with ordered names and times.
- Instead of two, set the number of players on the command line.
- Use color in the server console output.
- Add some test coverage for kart.go. Refactor as necessary.
- Have players run through multiple phrases, and display their progress on the server.

## Medium

- Pull phrases from external sources: tweets, wikipedia, fortune command, markov model...
- Let one of the players set the phrase, and the other players compete to write it first.
- Instead of just asking players to type what's presented, challenge them with a question (math, trivial pursuit facts, fill-in-the-blank, multiple choice, etc).
- Instead of using net.Conn, serve up html over http, and communicate with players over ajax or websocket.
- Instead of a fixed number, accept players for a fixed time at start-up, and then play.
- Be more robust. When one player's connection has a communication error, don't stop the game. Instead, just drop that player from the running.
- Start a new game instead of exiting at the end of a game.
- Create a persistent leaderboard; gather stats.
- Remember players by IP address.
- Support playing over a socket, UDP, or some other (non-TCP) net.Listener.
- Look at the unit test coverage; find a place to improve it.

## Big

- Instead of a client/server model, play peer-to-peer!
- Add a solo mode that collects your stats and reports your performance trends.
- Serve up an html leaderboard (or other stats) over http, or expose that data via a JSON API.
- Write a client instead of using telnet. This lets you do character-by-character interaction instead of line buffering.
- Put in some heuristics to detect that it is a human playing instead of a computer.
- Change the game mechanic to support team play.
- Write a unit test or two for Player that use a mock net.Listener and mock net.Conns.

## Twisted Big

- Instead of net.Conn, serve html over http, and wrap up the ajax or websocket interactions into a net.Listener. It'd be a bizarre thing to do, but shows how much fun you can really have with interfaces.
