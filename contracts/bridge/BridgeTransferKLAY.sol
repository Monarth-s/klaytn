// Copyright 2019 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

pragma solidity 0.5.6;

import "./BridgeTransfer.sol";


contract BridgeTransferVINI is BridgeTransfer {
    bool public isLockedVINI;

    event VINILocked();
    event VINIUnlocked();

    modifier lockedVINI {
        require(isLockedVINI == true, "unlocked");
        _;
    }

    modifier unlockedVINI {
        require(isLockedVINI == false, "locked");
        _;
    }

    // lockVINI can to prevent request VINI transferring.
    function lockVINI()
        external
        onlyOwner
        unlockedVINI
    {
        isLockedVINI = true;

        emit VINILocked();
    }

    // unlockToken can allow request VINI transferring.
    function unlockVINI()
        external
        onlyOwner
        lockedVINI
    {
        isLockedVINI = false;

        emit VINIUnlocked();
    }

    // handleVINITransfer sends the VINI by the request.
    function handleVINITransfer(
        bytes32 _requestTxHash,
        address _from,
        address payable _to,
        uint256 _value,
        uint64 _requestedNonce,
        uint64 _requestedBlockNumber,
        bytes memory _extraData
    )
        public
        onlyOperators
    {
        _lowerHandleNonceCheck(_requestedNonce);

        if (!_voteValueTransfer(_requestedNonce)) {
            return;
        }

        _setHandledRequestTxHash(_requestTxHash);

        handleNoncesToBlockNums[_requestedNonce] = _requestedBlockNumber;
        _updateHandleNonce(_requestedNonce);

        emit HandleValueTransfer(
            _requestTxHash,
            TokenType.VINI,
            _from,
            _to,
            address(0),
            _value,
            _requestedNonce,
            lowerHandleNonce,
            _extraData
        );

        _to.transfer(_value);
    }

    // _requestVINITransfer requests transfer VINI to _to on relative chain.
    function _requestVINITransfer(address _to, uint256 _feeLimit,  bytes memory _extraData)
        internal
        unlockedVINI
    {
        require(isRunning, "stopped bridge");
        require(msg.value > _feeLimit, "insufficient amount");

        uint256 fee = _payVINIFeeAndRefundChange(_feeLimit);

        emit RequestValueTransfer(
            TokenType.VINI,
            msg.sender,
            _to,
            address(0),
            msg.value.sub(_feeLimit),
            requestNonce,
            fee,
            _extraData
        );
        requestNonce++;
    }

    // () requests transfer VINI to msg.sender address on relative chain.
    function () external payable {
        _requestVINITransfer(msg.sender, feeOfVINI, new bytes(0));
    }

    // requestVINITransfer requests transfer VINI to _to on relative chain.
    function requestVINITransfer(address _to, uint256 _value, bytes calldata _extraData) external payable {
        uint256 feeLimit = msg.value.sub(_value);
        _requestVINITransfer(_to, feeLimit, _extraData);
    }

    // chargeWithoutEvent sends VINI to this contract without event for increasing
    // the withdrawal limit.
    function chargeWithoutEvent() external payable {}

    // setVINIFee set the fee of VINI transfer.
    function setVINIFee(uint256 _fee, uint64 _requestNonce)
        external
        onlyOperators
    {
        if (!_voteConfiguration(_requestNonce)) {
            return;
        }
        _setVINIFee(_fee);
    }
}
