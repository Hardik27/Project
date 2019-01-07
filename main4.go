// package main

// import (

// 	"bufio"
// 	"context"
// 	"flag"
// 	"os"
// 	"strings"

// 	"github.com/perlin-network/noise/crypto/ed25519"
// 	"./messages"
// 	"github.com/perlin-network/noise/log"
// 	"github.com/perlin-network/noise/network"
// 	"github.com/perlin-network/noise/network/discovery"
// 	"github.com/perlin-network/noise/types/opcode"
// )

// type ChatPlugin struct{ *network.Plugin }

// func (state *ChatPlugin) Receive(ctx *network.PluginContext) error {
// 	switch msg := ctx.Message().(type) {
// 	case *messages.ChatMessage:
// 		log.Info().Msgf("<%s> %s", ctx.Client().ID.Address, msg.Message)
// 	}

// 	return nil
// }

// func main() {
// 	// process other flags
// 	portFlag := flag.Int("port", 8003, "port to listen to")
// 	hostFlag := flag.String("host", "localhost", "host to listen to")
// 	protocolFlag := flag.String("protocol", "tcp", "protocol to use (kcp/tcp)")
// 	peersFlag := flag.String("peers", "", "peers to connect to")
// 	flag.Parse()

// 	port := uint16(*portFlag)
// 	host := *hostFlag
// 	protocol := *protocolFlag
// 	peers := strings.Split(*peersFlag, ",")

// 	keys := ed25519.RandomKeyPair()

// 	log.Info().Msgf("Private Key: %s", keys.PrivateKeyHex())
// 	log.Info().Msgf("Public Key: %s", keys.PublicKeyHex())

// 	opcode.RegisterMessageType(opcode.Opcode(1000), &messages.ChatMessage{})
// 	builder := network.NewBuilder()
// 	builder.SetKeys(keys)
// 	builder.SetAddress(network.FormatAddress(protocol, host, port))

// 	// Register peer discovery plugin.
// 	builder.AddPlugin(new(discovery.Plugin))

// 	// Add custom chat plugin.
// 	builder.AddPlugin(new(ChatPlugin))

// 	net, err := builder.Build()
// 	if err != nil {
// 		log.Fatal().Err(err)
// 		return
// 	}

// 	go net.Listen()

// 	if len(peers) > 0 {
// 		net.Bootstrap("tcp://localhost:8000","tcp://localhost:8001","tcp://localhost:8002")
// 	}

// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		input, _ := reader.ReadString('\n')

// 		// skip blank lines
// 		if len(strings.TrimSpace(input)) == 0 {
// 			continue
// 		}

// 		log.Info().Msgf("<%s> %s", net.Address, input)

// 		ctx := network.WithSignMessage(context.Background(), true)
// 		net.Broadcast(ctx, &messages.ChatMessage{Message: input})
// 	}
// }

package main

import (
    //"bufio"
	"context"
	"flag"
	//"os"
	"strings"
	"strconv"
	"./pow"
	"github.com/perlin-network/noise/crypto/ed25519"
	"./messages"
	"github.com/perlin-network/noise/log"
	"github.com/perlin-network/noise/network"
	"github.com/perlin-network/noise/network/discovery"
	"github.com/perlin-network/noise/types/opcode"
)

type POWPlugin struct{ *network.Plugin }

func (state *POWPlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *messages.POWMessage:
		h := pow.Encode(msg.Addr+msg.Pubkey+strconv.Itoa(int(msg.Blocknum))+strconv.Itoa(int(msg.Nonce)))
		if h==msg.Result && pow.IsValidResult(h){
			log.Info().Msgf("<%s> true",ctx.Client().ID.Address)
		}else{
			log.Info().Msgf("<%s> false",ctx.Client().ID.Address)
		}
		log.Info().Msgf("<%s> %s", ctx.Client().ID.Address, strconv.Itoa(int(msg.Nonce)) +"   "+ msg.Pubkey +"   " + msg.Addr+"   " + msg.Result)
	}

	return nil
}

func main() {
	// process other flags
	portFlag := flag.Int("port", 3003, "port to listen to")
	hostFlag := flag.String("host", "localhost", "host to listen to")
	protocolFlag := flag.String("protocol", "tcp", "protocol to use (kcp/tcp)")
	peersFlag := flag.String("peers", "", "peers to connect to")
	flag.Parse()

	port := uint16(*portFlag)
	host := *hostFlag
	protocol := *protocolFlag
	peers := strings.Split(*peersFlag, ",")

	keys := ed25519.RandomKeyPair()

	log.Info().Msgf("Private Key: %s", keys.PrivateKeyHex())
	log.Info().Msgf("Public Key: %s", keys.PublicKeyHex())

	opcode.RegisterMessageType(opcode.Opcode(1000), &messages.POWMessage{})
	builder := network.NewBuilder()
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress(protocol, host, port))

	// Register peer discovery plugin.
	builder.AddPlugin(new(discovery.Plugin))

	// Add custom chat plugin.
	builder.AddPlugin(new(POWPlugin))

	net, err := builder.Build()
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	go net.Listen()

	if len(peers) > 0 {
		net.Bootstrap("tcp://localhost:3001","tcp://localhost:3002","tcp://localhost:3000")
	}

	//reader := bufio.NewReader(os.Stdin)
	nonce,result,blocknum,difficulty := pow.Pow(keys.PublicKeyHex())
	ctx := network.WithSignMessage(context.Background(), true)
		net.Broadcast(ctx, &messages.POWMessage{Nonce: int32(nonce),Pubkey:keys.PublicKeyHex(),Addr:pow.GetOutboundIP().String(),Blocknum:int32(blocknum),Difficulty:int32(difficulty),Result:result})
	for {
		//input, _ := reader.ReadString('\n')

		// skip blank lines
		/*if len(strings.TrimSpace(input)) == 0 {
			continue
		}*/

		//log.Info().Msgf("<%s> %s", net.Address, input)

		
	}

}