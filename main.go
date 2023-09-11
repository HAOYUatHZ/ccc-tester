package main

import (
	"context"
	"flag"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"io"
	"math/big"
	"os"
	"tool/accounts"
	"tool/api"
)

var (
	wrapJson   = flag.Bool("wrap", false, "add json rpc wrapJson")
	endpoint   = flag.String("endpoint", "ws://127.0.0.1:8546", "The endpoint to connect to blockchain node")
	keystore   = flag.String("keystore", "./docker/l2geth/genesis-keystore", "Keystore file path")
	password   = flag.String("password", "scrolltest", "The keystore password")
	dump       = flag.String("dump", "erc20", "e.g: erc20, native, nft, greeter, sushi, dao, uniswapv2, ecAdd, ecPairing, ecRecover, sha256, keccak256, modexp")
	timesInner = flag.Int64("n", 50, "iteration times in a contract call")
	timesOuter = flag.Int64("m", 1, "times of contract call")
)

func init() {
	output := io.Writer(os.Stderr)
	usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
	if usecolor {
		output = colorable.NewColorableStderr()
	}
	ostream := log.StreamHandler(output, log.TerminalFormat(usecolor))
	glogger := log.NewGlogHandler(ostream)
	// Set log level
	glogger.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogger)
}

func main() {
	// Parse the flags and set up the logger to print everything requested
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Enable wrap json rpc result or not.
	api.WrapJson = *wrapJson

	// create client
	client, err := ethclient.DialContext(ctx, *endpoint)
	if err != nil {
		log.Crit("Failed to connect chain node", "err", err)
	}

	// load keystore file
	auths, err := accounts.NewAccounts(ctx, 2, client, *keystore, *password)
	if err != nil {
		log.Crit("failed to create accounts", "err", err)
	}
	root := auths.Root
	auth := auths.GetAccount()

	solName := api.SolType(*dump)
	switch solName {
	case api.NativeName:
		err = api.Native(ctx, client, root, auth.From, big.NewInt(100))
	case api.ERC20Name:
		err = api.NewERC20(ctx, client, root, auth)
	case api.NftName:
		err = api.NewNft(ctx, client, root, auth)
	case api.GreeterName:
		err = api.NewGreeter(ctx, client, root)
	case api.SushiName:
		err = api.NewSushi(ctx, client, root)
	case api.DaoName:
		err = api.NewDao(ctx, client, root, auth)
	case api.Uniswapv2Name:
		err = api.NewUniswapv2(ctx, client, root, auth, *timesOuter)
	case api.EcAddName:
		err = api.NewEcc(ctx, client, root, auth, "add", *timesInner, *timesOuter)
	case api.EcMulName:
		err = api.NewEcc(ctx, client, root, auth, "mul", *timesInner, *timesOuter)
	case api.EcPairingName:
		err = api.NewEcc(ctx, client, root, auth, "pairing", *timesInner, *timesOuter)
	case api.EcRecoverName:
		err = api.NewEcc(ctx, client, root, auth, "recover", *timesInner, *timesOuter)
	case api.Sha256Name:
		err = api.NewHash(ctx, client, root, auth, "sha256", *timesInner, *timesOuter)
	case api.Keccak256Name:
		err = api.NewHash(ctx, client, root, auth, "keccak256", *timesInner, *timesOuter)
	case api.ModExpName:
		err = api.NewModExp(ctx, client, root, auth, *timesInner, *timesOuter)
	default:
		log.Error("unexpected dump option")
		return
	}

	if err != nil {
		log.Error("dump traces for contract fail", "contract name", solName, "err", err)
	} else {
		log.Info("dump traces for contract successfully", "contract name", solName)
	}
}
