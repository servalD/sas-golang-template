// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {ERC721} from "../lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol";

contract MyNFT is ERC721 {
    uint256 public tokenCounter;

    constructor() ERC721("MyNFT", "MNFT") {
        tokenCounter = 0;
    }
    
    // Fonction de mint public
    function mint(address to) public {
        _mint(to, tokenCounter);
        tokenCounter++;
    }
}
