# Client - Server: Proof of Work Challenge

The application simulates an intereation between a client and a server.

The main.go file starts the process
The client (client.go) calls the server (server.go) to receive a random phrase from Words of Wisdom Book from the ones available on the server side.

The server responds with a challenge. This challenge is in the following struct:

```
type Challenge struct {
	Nonce    int32
	DataHash string
	Prefix   string
}
```

## Challenge
The nonce will at first be a zero, the prefix will be the number of zeroes set on the difficulty.
The hash is a SHA256 hash generated with a number interval.

The client must perform the proof of work using the challenge it received.
The prefix is a sequence of zeroes, determined by the difficulty set on the server. For now, it's fixed like follows:

```
return generateChallenge(5)
```

The client performs the proof of work by incresing the nonce and re-generating a hash until the prefix matches the prefix. For instance:

0000089c2a4aeae1c623657d7f09c94274d1786124edbe12ceac3f639b1041a0

After doing that, the client requests the data again, by calling 

```
data, err := server.GetData(chal)

```
using the challeng with the new nonce - the one that generated the valid hash.

The server will generate a hash with the challenge struct updated by the client.

If it macthes, a phrase will be returned. Otherwise, an error will.