// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract MyContract {
    uint256 public count;

    event CountIncremented(uint256 newCount);

    function increment() public {
        count += 1;
        emit CountIncremented(count);
    }

    function getCount() public view returns (uint256) {
        return count;
    }
}
