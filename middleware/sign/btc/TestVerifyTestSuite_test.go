package btc

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This example will output `true, <nil>` for both signed messages, since the signature is valid and there are no errors.
func TestVerifyTestSuiteBitcoinMainne(t *testing.T) {
	// Bitcoin Mainnet
	fmt.Println(VerifyWithChain(SignedMessage{
		Address:   "18J72YSM9pKLvyXX1XAjFXA98zeEvxBYmw",
		Message:   "Test123",
		Signature: "Gzhfsw0ItSrrTCChykFhPujeTyAcvVxiXwywxpHmkwFiKuUR2ETbaoFcocmcSshrtdIjfm8oXlJoTOLosZp3Yc8=",
	}, &chaincfg.MainNetParams))
}
func TestVerifyTestSuiteBitcoinTestnet3Taproot(t *testing.T) {
	// Bitcoin Testnet3 - taproot
	chain, err := VerifyWithChain(SignedMessage{
		Address:   "tb1p8fekvp3f5s92rjl539nmn7wkl8xa4xh0h4n5lh9r8gdnqeyjn00q5v883t",
		Message:   "hello world~",
		Signature: "H76e4ZSmuR8XT8QTDXKThiGDymfsdxzBs7cWVRZ6vdZVE86ozTm4Om9LNKd4oVSniu+0K3lx24klEICD3WDlq00=",
	}, &chaincfg.TestNet3Params)
	assert.Nil(t, err)
	assert.True(t, chain)
}

// pk:
// wif: cTyQZkTnMzNRrJehkP97zp93KBZyw2557MWuiuP3a22YB4rfDD91
// hex: beb47646e710feb8a89d1957f5df0e7bc16431984841ebc49a0653b08f28ce1e
//
//puk:020d66cdc85d18dcb3992799e9475c4a03f7fa53259141d37088768bb188c5437e
func TestVerifyTestSuiteBitcoinTestnet3NativeSegwit(t *testing.T) {
	// Bitcoin Testnet3 - native segwit
	chain, err := VerifyWithChain(SignedMessage{
		Address:   "tb1q8td5rvfam3q90lumefqfxf5rzu6wyjjkxecrm3",
		Message:   "hello world~",
		Signature: "H5hxqB9HtwWfcFjP9lkHXR9ZV9di2rOIjB9MdtxMvjatNHw1olgrHh80Qe4G2KPjLwQvYvhVWr9k8MpSqUjtiFo=\n",
	}, &chaincfg.TestNet3Params)
	assert.Nil(t, err)
	assert.True(t, chain)
}

// pk cPFrgjL4EfRXy2GQ6VEpPLZfivWV1PFj4dPmibccZ8TNY6v4Tzu6
// hex 31fb0a9682eb2be81a2d0a35a2dc5244f116463d27e4c3894c94a05904d2a6e1
// puk 0356416d5cd931cfc1a116ba498ff6bf4e32c0dc84a26cbef69b99b047642d790b
func TestVerifyTestSuiteBitcoinTestnet3NestedSegwit(t *testing.T) {
	// Bitcoin Testnet3 - native segwit
	chain, err := VerifyWithChain(SignedMessage{
		Address:   "2N2UKiaxmN79ziqyBky8RRVVeHfPj8hXVxi",
		Message:   "hello world~",
		Signature: "IDrLhnsPxGUWvV2QgC7GMCStQiAOaPKKTVOFqOvTi9KNQs+V0Vtx70JJhFWh720bewk47viHMuD19utoHYHX1ss=",
	}, &chaincfg.TestNet3Params)
	assert.Nil(t, err)
	assert.True(t, chain)
}
