package snapshothashes

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/ledgerwatch/erigon-snapshot/webseed"
)

var branchReference = getBranchReference()

func getBranchReference() string {
	v, _ := os.LookupEnv("SNAPS_GIT_BRANCH")
	if v != "" {
		return v
	}
	return "main"
}

//go:embed mainnet.toml
var Mainnet []byte

//go:embed sepolia.toml
var Sepolia []byte

//go:embed amoy.toml
var Amoy []byte

//go:embed bor-mainnet.toml
var BorMainnet []byte

//go:embed gnosis.toml
var Gnosis []byte

//go:embed chiado.toml
var Chiado []byte

//go:embed holesky.toml
var Holesky []byte

//go:embed bsc.toml
var Bsc []byte

//go:embed chapel.toml
var Chapel []byte

func getURLByChain(chain string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/node-real/bsc-erigon-snapshot/%s/%s.toml", branchReference, chain)
}

func LoadSnapshots() (couldFetch bool) {
	var (
		mainnetUrl    = getURLByChain("mainnet")
		sepoliaUrl    = getURLByChain("sepolia")
		amoyUrl       = getURLByChain("amoy")
		borMainnetUrl = getURLByChain("bor-mainnet")
		gnosisUrl     = getURLByChain("gnosis")
		chiadoUrl     = getURLByChain("chiado")
		holeskyUrl    = getURLByChain("holesky")
		bscUrl        = getURLByChain("bsc")
		chapelUrl     = getURLByChain("chapel")
	)
	var hashes []byte
	var err error
	// Try to fetch the latest snapshot hashes from the web
	if hashes, err = fetchSnapshotHashes(mainnetUrl); err != nil {
		couldFetch = false
		return
	}
	Mainnet = hashes

	if hashes, err = fetchSnapshotHashes(sepoliaUrl); err != nil {
		couldFetch = false
		return
	}
	Sepolia = hashes

	if hashes, err = fetchSnapshotHashes(amoyUrl); err != nil {
		couldFetch = false
		return
	}
	Amoy = hashes

	if hashes, err = fetchSnapshotHashes(borMainnetUrl); err != nil {
		couldFetch = false
		return
	}
	BorMainnet = hashes

	if hashes, err = fetchSnapshotHashes(gnosisUrl); err != nil {
		couldFetch = false
		return
	}
	Gnosis = hashes

	if hashes, err = fetchSnapshotHashes(chiadoUrl); err != nil {
		couldFetch = false
		return
	}
	Chiado = hashes

	if hashes, err = fetchSnapshotHashes(holeskyUrl); err != nil {
		couldFetch = false
		return
	}
	Holesky = hashes

	if hashes, err = fetchSnapshotHashes(bscUrl); err != nil {
		couldFetch = false
		return
	}
	Bsc = hashes

	if hashes, err = fetchSnapshotHashes(chapelUrl); err != nil {
		couldFetch = false
		return
	}
	Chapel = hashes

	couldFetch = true
	return
}

func fetchSnapshotHashes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if len(res) == 0 {
		return nil, fmt.Errorf("empty response from %s", url)
	}
	return res, err
}
