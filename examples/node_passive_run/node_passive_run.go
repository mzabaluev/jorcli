//$(which go) run $0 $@; exit $?

package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/rinor/jorcli/jcli"
	"github.com/rinor/jorcli/jnode"
)

// fatalOn be careful with it in production,
// since it uses os.Exit(1) which affects the control flow.
// use pattern:
// if err != nil {
// 	....
// }
func fatalOn(err error, str ...string) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Fatalf("%s:%d %s -> %s", fn, line, str, err.Error())
	}
}

// seed generated from an int. For the same int the same seed is returned.
// Useful for reproducible batch key generation,
// for example the index of a slice/array can be a param.
func seed(i int) string {
	in := []byte(strconv.Itoa(i))
	out := make([]byte, 32-len(in), 32)
	out = append(out, in...)

	return hex.EncodeToString(out)
}

// b2s converts []byte to string with all leading
// and trailing white space removed, as defined by Unicode.
func b2s(b []byte) string {
	return strings.TrimSpace(string(b))
}

/* seeds used [30] */
const (
	seedPrivateID = 30 // seed for p2p private_id
)

func main() {

	var (
		err error

		// Rest
		restAddr    = "127.0.0.44" // rest ip
		restPort    = 8001         // rest port
		restAddress = restAddr + ":" + strconv.Itoa(restPort)

		// P2P
		p2pIPver = "ip4" // ipv4 or ipv6
		p2pProto = "tcp" // tcp

		// P2P Public
		p2pPubAddr       = "127.0.0.44" // PublicAddres
		p2pPubPort       = 9001         // node P2P Public Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPubPort)

		// P2P Listen
		p2pListenAddr    = "127.0.0.44" // ListenAddress
		p2pListenPort    = 9001         // node P2P Public Port
		p2pListenAddress = "/" + p2pIPver + "/" + p2pListenAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pListenPort)

		// Trusted peers
		leaderAddr = "/ip4/127.0.0.11/tcp/9001"                                              // Leader (genesis) node (example 1)
		leaderID   = "ed25519_pk1thawa4wxfhn9hh9xll04npw9pv0djgnvcun90nw9szupfw95lvns94qgpu" // Leader public_id

		gepAddr = "/ip4/127.0.0.22/tcp/9001"                                              // Genesis stake pool node (example 2)
		gepID   = "ed25519_pk1z5u62jwftwrepu53nj655cdzjrhv4dlry9d7c602j6dagfpwp34q5gjcmr" // Genesis stake pool public_id

		delegatorAddr = "/ip4/127.0.0.33/tcp/9001"                                              // stake pool node (example 3)
		delegatorID   = "ed25519_pk19qzyd6xxed7rc3nxj0qgnsuyxkpqvlcue44l7l3f5kkr9dj378ss2wnm22" // delegator pool public_id

		// Genesis Block0 Hash retrieved from example (1)
		block0Hash = "999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d"
	)

	// Set RUST_BACKTRACE=full env
	err = os.Setenv("RUST_BACKTRACE", "full")
	fatalOn(err, "Failed to set env (RUST_BACKTRACE=full)")

	// set binary name/path if not default,
	// provided as example since the ones set here,
	// are also the default values.
	jcli.BinName("jcli")         // default is "jcli"
	jnode.BinName("jormungandr") // default is "jormungandr"

	// get jcli version
	jcliVersion, err := jcli.VersionFull()
	fatalOn(err, b2s(jcliVersion))
	log.Printf("Using: %s", jcliVersion)

	// get jormungandr version
	jormungandrVersion, err := jnode.VersionFull()
	fatalOn(err, b2s(jormungandrVersion))
	log.Printf("Using: %s", jormungandrVersion)

	// create a new temporary directory inside your systems temp dir
	workingDir, err := ioutil.TempDir("", "jnode_")
	fatalOn(err, "workingDir")
	log.Println()
	log.Printf("Working Directory: %s", workingDir)

	///////////////////
	//  node config  //
	///////////////////

	// p2p node private_id
	nodePrivateID, err := jcli.KeyGenerate(seed(seedPrivateID), "Ed25519", "")
	fatalOn(err, b2s(nodePrivateID))
	// node's unique identifier on the network
	nodePublicID, err := jcli.KeyToPublic(nodePrivateID, "", "")
	fatalOn(err, b2s(nodePublicID))
	// node's unique identifier on the network as displayed in logs
	nodePublicIDBytes, err := jcli.KeyToBytes(nodePublicID, "", "")
	fatalOn(err, b2s(nodePublicIDBytes))

	nodeCfg := jnode.NewNodeConfig()

	nodeCfg.Storage = "jnode_storage"

	nodeCfg.Rest.Enabled = true       // default is "false" (rest disabled)
	nodeCfg.Rest.Listen = restAddress // 127.0.0.1:8443 is default value

	nodeCfg.Explorer.Enabled = true // default is "false" (explorer disabled)

	nodeCfg.P2P.PublicAddress = p2pPublicAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.ListenAddress = p2pListenAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.PrivateID = b2s(nodePrivateID)   // jörmungandr will generate a random key, if not set
	nodeCfg.P2P.AllowPrivateAddresses = true     // for private addresses

	// add trusted peer to config file
	nodeCfg.AddTrustedPeer(leaderAddr, leaderID)
	nodeCfg.AddTrustedPeer(gepAddr, gepID)
	nodeCfg.AddTrustedPeer(delegatorAddr, delegatorID)

	nodeCfg.Log.Level = "info" // default is "trace"

	nodeCfgYaml, err := nodeCfg.ToYaml()
	fatalOn(err)
	// need this file for starting the node (--config)
	nodeCfgFile := workingDir + string(os.PathSeparator) + "node-config.yaml"
	err = ioutil.WriteFile(nodeCfgFile, nodeCfgYaml, 0644)
	fatalOn(err)

	// fmt.Printf("%s", nodeCfgYaml)

	//////////////////////
	// running the node //
	//////////////////////

	node := jnode.NewJnode()

	node.WorkingDir = workingDir
	node.ConfigFile = nodeCfgFile
	node.GenesisBlockHash = block0Hash // add block0 hash

	// add trusted peer cmd args (not needed if using config)
	node.AddTrustedPeer(leaderAddr, leaderID)       // add leader from example (1) as trusted
	node.AddTrustedPeer(gepAddr, gepID)             // add genesis stake pool from example (2) as trusted
	node.AddTrustedPeer(delegatorAddr, delegatorID) // add delegator stake pool from example (3) as trusted

	node.Stdout, err = os.Create(filepath.Join(workingDir, "stdout.log"))
	fatalOn(err)
	node.Stderr, err = os.Create(filepath.Join(workingDir, "stderr.log"))
	fatalOn(err)

	// Run the node (Start + Wait)
	err = node.Run()
	if err != nil {
		log.Fatalf("node.Run FAILED: %v", err)
	}

	log.Println()
	log.Printf("Genesis Hash: %s", block0Hash)
	log.Println()
	log.Printf("NodePublicID for trusted: %s", nodePublicID)
	log.Printf("NodePublicID in logs    : %s", b2s(nodePublicIDBytes))
	log.Println()

	log.Println("Passive/Explorer Node - Running...")
	node.Wait()                                   // Wait for the node to stop.
	log.Println("...Passive/Explore Node - Done") // All done. Node has stopped.
}
