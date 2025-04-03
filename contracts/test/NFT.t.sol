// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/NFT.sol";

contract MyNFTTest is Test {
    MyNFT nft;

    function setUp() public {
        nft = new MyNFT();
    }

    function testMintFiveNFTs() public {
        for (uint256 i = 0; i < 5; i++) {
            nft.mint(address(this));
        }
        assertEq(nft.balanceOf(address(this)), 5, "Balance should be 5");
        assertEq(nft.tokenCounter(), 5, "Token counter should be 5");
    }
}
