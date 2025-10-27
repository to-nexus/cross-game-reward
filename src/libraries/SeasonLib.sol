// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title SeasonLib
 * @notice 시즌 관리 헬퍼 라이브러리
 * @dev 시즌 검증 및 상태 확인 로직
 */
library SeasonLib {
    /**
     * @notice 시즌 활성 여부 확인
     * @param startBlock 시작 블록
     * @param endBlock 종료 블록
     * @param isFinalized 종료 여부
     * @return active 활성 여부
     */
    function isSeasonActive(uint startBlock, uint endBlock, bool isFinalized) internal view returns (bool active) {
        if (isFinalized) return false;
        return block.number >= startBlock && block.number <= endBlock;
    }

    /**
     * @notice 시즌 종료 여부 확인
     * @param endBlock 종료 블록
     * @return ended 종료 여부
     */
    function isSeasonEnded(uint endBlock) internal view returns (bool ended) {
        return block.number > endBlock;
    }

    /**
     * @notice 시즌 범위 검증
     * @param startBlock 시작 블록
     * @param endBlock 종료 블록
     */
    function validateSeasonBlocks(uint startBlock, uint endBlock) internal pure {
        require(endBlock > startBlock, "Invalid season blocks");
    }

    /**
     * @notice 시즌 내 블록인지 확인
     * @param blockNumber 확인할 블록
     * @param startBlock 시즌 시작 블록
     * @param endBlock 시즌 종료 블록
     * @return inSeason 시즌 내 여부
     */
    function isBlockInSeason(uint blockNumber, uint startBlock, uint endBlock) internal pure returns (bool inSeason) {
        return blockNumber >= startBlock && blockNumber <= endBlock;
    }

    /**
     * @notice 유효 시작 블록 계산
     * @param userJoinBlock 사용자 참여 블록
     * @param seasonStartBlock 시즌 시작 블록
     * @return effectiveStart 유효 시작 블록
     */
    function calculateEffectiveStart(uint userJoinBlock, uint seasonStartBlock)
        internal
        pure
        returns (uint effectiveStart)
    {
        return userJoinBlock > seasonStartBlock ? userJoinBlock : seasonStartBlock;
    }
}
