// SPDX-License-Identifier: MIT

// Written by Abderraouf Zine

pragma solidity >=0.7.0 <0.9.0;

contract Artwork {
    string artworkName;
    uint256 goalPrice;
    address originalOwner;
    address currentOwner;
    uint256 lastSellingPrice = 0;

    constructor(string memory _artworkName, uint256 _goalPrice) {
        artworkName = _artworkName;
        goalPrice = _goalPrice;
        originalOwner = msg.sender;
        currentOwner = msg.sender;
    }

    function buy(uint256 offer) public {
        require(msg.sender != currentOwner, "You arleady own this artwork.");
        require(
            offer >= goalPrice - goalPrice / 20,
            "You need to offer more than this."
        );
        currentOwner = msg.sender;
        lastSellingPrice = offer;
    }
}
