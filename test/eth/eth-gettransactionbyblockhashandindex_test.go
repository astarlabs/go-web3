/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file eth-sendtransaction_test.go
 * @authors:
 *      Sigma Prime <sigmaprime.io>
 * @date 2018
 */
package test

import (
"testing"

"github.com/regcostajr/go-web3"
"github.com/regcostajr/go-web3/complex/types"
"github.com/regcostajr/go-web3/dto"
"github.com/regcostajr/go-web3/providers"
"math/big"
)

func TestGetTransactionByBlockHashAndIndex(t *testing.T) {

    var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

    coinbase, err := connection.Eth.GetCoinbase()

    if err != nil {
        t.Error(err)
        t.FailNow()
    }

    txVal := big.NewInt(2000000)

    transaction := new(dto.TransactionParameters)
    transaction.From = coinbase
    transaction.To = coinbase
    //transaction.Value = big.NewInt(0).Mul(big.NewInt(500), big.NewInt(1E18))
    transaction.Value = txVal
    transaction.Gas = big.NewInt(40000)

    txID, err := connection.Eth.SendTransaction(transaction)

    t.Log("Tx Submitted: ", txID)

    if err != nil {
        t.Error(err)
        t.FailNow()
    }

    blockNumber, err := connection.Eth.GetBlockNumber()
    num := types.ComplexIntParameter(blockNumber.ToInt64())
    if err != nil {
        t.Error(err)
        t.Fail()
    }

    block, err := connection.Eth.GetBlockByNumber(num, false)

    if err != nil {
        t.Error(err)
        t.Fail()
    }

    index := types.ComplexIntParameter(0)
    tx, err := connection.Eth.GetTransactionByBlockHashAndIndex(block.Hash, index)

    if err != nil {
        t.Error(err)
        t.FailNow()
    }


    if tx.From != coinbase || tx.To != coinbase || tx.Value.ToInt64() != txVal.Int64() || tx.Hash != txID {
        t.Errorf("Incorrect transaction from hash and index")
        t.FailNow()
    }

    t.Log("BLOCK", block.Hash)
    t.Log("BLOCK", index)
    t.Log("BLOCK", tx.Hash)
    t.Log("BLOCK", tx.BlockHash)
    t.Log("BLOCK", tx.BlockNumber)
    t.Log("BLOCK", tx.TransactionIndex)

}
