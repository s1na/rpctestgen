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

func checkStatus(value type?, expected type?) error {
	if value != expected {
		return fmt.Errorf("expected status (have: %d, want: %d)", value, expected)
	}
}

func checkError(value number, expected number) error {
	if value != expected {
		return fmt.Errorf("expected error (have: %d, want: %d)", value, expected)
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

func selfDestructor() {
	//Solidity code:
	//contract SelfDestructor {
	//	function destroy() public {
	//		selfdestruct(payable(0x0));
	//	}
	//}
	return hex2Bytes("6080604052348015600f57600080fd5b506004361060285760003560e01c806383197ef014602d575b600080fd5b60336035565b005b600073ffffffffffffffffffffffffffffffffffffffff16fffea26469706673582212208e566fde20a17fff9658b9b1db37e27876fd8934ccf9b2aa308cabd37698681f64736f6c63430008120033")
}

func getCode() {
	//library GetCode {
	//	function at(address addr) public view returns (bytes memory code) {
	//		assembly {
	//			// retrieve the size of the code, this needs assembly
	//			let size := extcodesize(addr)
	//			// allocate output byte array - this could also be done without assembly
	//			// by using code = new bytes(size)
	//			code := mload(0x40)
	//			// new "memory end" including padding
	//			mstore(0x40, add(code, and(add(add(size, 0x20), 0x1f), not(0x1f))))
	//			// store length in memory
	//			mstore(code, size)
	//			// actually retrieve the code, this needs assembly
	//			extcodecopy(addr, add(code, 0x20), 0, size)
	//		}
	//	}
	//}
	return hex2Bytes("73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063dce4a4471461003a575b600080fd5b610054600480360381019061004f91906100f8565b61006a565b60405161006191906101b5565b60405180910390f35b6060813b6040519150601f19601f602083010116820160405280825280600060208401853c50919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100c58261009a565b9050919050565b6100d5816100ba565b81146100e057600080fd5b50565b6000813590506100f2816100cc565b92915050565b60006020828403121561010e5761010d610095565b5b600061011c848285016100e3565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561015f578082015181840152602081019050610144565b60008484015250505050565b6000601f19601f8301169050919050565b600061018782610125565b6101918185610130565b93506101a1818560208601610141565b6101aa8161016b565b840191505092915050565b600060208201905081810360008301526101cf818461017c565b90509291505056fea26469706673582212206a5f0cd9f230619fa520fc4b9d4b518643258cad412f2fa33945ce528b4b895164736f6c63430008120033")
}
func gasSpender() {
	//contract GasSpender {
	//	function spendGas(gasToSpend: uint) view external {
	//		uint public gasLeftInitially = gasleft();
	//		while(true) {
	//			if (gasLeftInitially - gasleft() >= gasToSpend) return;
	//		}
	//	}
	//}
	return hex2Bytes("608060405234801561001057600080fd5b506004361061002b5760003560e01c8063815b8ab414610030575b600080fd5b61004a600480360381019061004591906100b6565b61004c565b005b60005a90505b60011561007657815a826100669190610112565b106100715750610078565b610052565b505b50565b600080fd5b6000819050919050565b61009381610080565b811461009e57600080fd5b50565b6000813590506100b08161008a565b92915050565b6000602082840312156100cc576100cb61007b565b5b60006100da848285016100a1565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061011d82610080565b915061012883610080565b92508282039050818111156101405761013f6100e3565b5b9291505056fea2646970667358221220a659ba4db729a6ee4db02fcc5c1118db53246b0e5e686534fc9add6f2e93faec64736f6c63430008120033")
}
