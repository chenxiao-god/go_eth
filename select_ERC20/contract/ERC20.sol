// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
contract ERC20MinerReward is ERC20 {
     address private owner;
    event LogNewAlert(string description, address indexed _from, uint256 _n);
    constructor() ERC20("MinerReward123", "MRW123") {
        owner = msg.sender;
    }
    function _reward() public {
        _mint(owner, 1000*10**18);
        emit LogNewAlert("_rewarded", block.coinbase, block.number);
    }
}