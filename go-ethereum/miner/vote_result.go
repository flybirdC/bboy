package miner

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"strconv"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)


type VoteContent struct {

	//beatTime common.Hash
	Address common.Address
	Num int64
}

//每n块获取
func GetWitNessResult(n uint64, currentHeader *types.Header, bc *core.BlockChain) (int, []common.Address) {
	//当前区块高度与n的同余

	hm := currentHeader.Number.Uint64()
	totalVote := []VoteContent{}
	//记票
	totalNum := make(map[int]int)
	if hm % n == 0  && hm != 0{
		for i:=0;i<10 ;i++  {
			//得到块
			block := bc.GetBlockByNumber(hm)
			//得到块交易数据
			transactions := block.Body().Transactions
			//遍历交易切片取值，根据交易属性witness
			for _,v := range transactions{
				if v.GetTranType() == ethapi.TransactionType0 {
					//取交易随机数结果
					index,err := strconv.Atoi(string(v.Data()))
					if err != nil {
						break
					}
					//得到所有投票
					totalVote = append(totalVote,VoteContent{v.GetAddress(),int64(index)})
					totalNum[index] = totalNum[index]+1
				}
			}
			//遍历递减
			hm--
		}
	}
	//唱票数, 得到中奖数
	var keyluck  = 0
	 var valueluck = 0
	for k,v := range totalNum{
		if v > valueluck {
			valueluck = v
			keyluck = k
		}
	}
	//得到中奖人地址
	luckAddress := []common.Address{}
	for _,vote := range totalVote{
		if keyluck == int(vote.Num) {
			luckAddress = append(luckAddress,vote.Address)
		}
	}
	return keyluck, nil
}

