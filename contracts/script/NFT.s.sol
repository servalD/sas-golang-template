// script/DeployMyNFT.s.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/NFT.sol";

contract DeployMyNFT is Script {
    function run() external {
        vm.startBroadcast();
        MyNFT nft = new MyNFT();
        // Mint 5 NFT au déployeur
        for (uint256 i = 0; i < 5; i++) {
            nft.mint(msg.sender);
        }
        vm.stopBroadcast();
    }
}
