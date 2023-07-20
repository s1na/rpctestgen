package testgen

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

func checkHeaderRLP(t *T, n uint64, got []byte) error {
	head := t.chain.GetHeaderByNumber(n)
	if head == nil {
		return fmt.Errorf("unable to load block %d from test chain", n)
	}
	want, err := rlp.EncodeToBytes(head)
	if err != nil {
		return err
	}
	if hexutil.Encode(got) != hexutil.Encode(want) {
		return fmt.Errorf("unexpected response (got: %s, want: %s)", got, hexutil.Bytes(want))
	}
	return nil
}

func checkBlockRLP(t *T, n uint64, got []byte) error {
	head := t.chain.GetBlockByNumber(n)
	if head == nil {
		return fmt.Errorf("unable to load block %d from test chain", n)
	}
	want, err := rlp.EncodeToBytes(head)
	if err != nil {
		return err
	}
	if hexutil.Encode(got) != hexutil.Encode(want) {
		return fmt.Errorf("unexpected response (got: %s, want: %s)", got, hexutil.Bytes(want))
	}
	return nil
}

// I guess Go doesn't have uint256 type?
func checkBlockNumber(value uint256, expected uint256) error {
	if value != expected {
		return fmt.Errorf("unexpected block number value(have: %d, want: %d)", value, expected)
	}
}

func checkBlockHash(value uint256, expected uint256) error {
	if value != expected {
		return fmt.Errorf("unexpected block hash value(have: %d, want: %d)", value, expected)
	}
}

func blockHashCallerByteCode() {
	//Solidity code:
	//contract blockHashCaller {
	//	function getBlockHash(uint256 blockNumber) public view returns (bytes32 blockHash) {
	//		blockHash = blockhash(blockNumber);
	//	}
	//}
	return hex2Bytes("0x6080604052348015600f57600080fd5b506004361060285760003560e01c8063ee82ac5e14602d575b600080fd5b60436004803603810190603f91906098565b6057565b604051604e919060d7565b60405180910390f35b600081409050919050565b600080fd5b6000819050919050565b6078816067565b8114608257600080fd5b50565b6000813590506092816071565b92915050565b60006020828403121560ab5760aa6062565b5b600060b7848285016085565b91505092915050565b6000819050919050565b60d18160c0565b82525050565b600060208201905060ea600083018460ca565b9291505056fea2646970667358221220a4d7face162688805e99e86526524ac3dadfb01cc29366d0d68b70dadcf01afe64736f6c63430008120033")
}
