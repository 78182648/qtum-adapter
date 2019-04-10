/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package qtum

import (
	"testing"
	"encoding/hex"
	"github.com/shopspring/decimal"
	"github.com/blocktree/openwallet/openwallet"
)

var (
	isTestNet = true
)

func Test_addressTo32bytesArg(t *testing.T) {
	address := "qdphfFinfJutJFvtnr2UaCwNAMxC3HbVxa"

	to32bytesArg, err := AddressTo32bytesArg(address, isTestNet)
	if err != nil {
		t.Errorf("To32bytesArg failed unexpected error: %v\n", err)
	}else {
		t.Logf("To32bytesArg success.")
	}

	t.Logf("This is to32bytesArg string for you to use: %s\n", hex.EncodeToString(to32bytesArg))
}


func Test_getUnspentByAddress(t *testing.T) {

	contractAddress := "0x91a6081095ef860d28874c9db613e7a4107b0281"
	address := "qMrFCXxSuiTEDqd311VxivECpJqH5KsJg6"
	var tokenDecimal uint64 = 8

	//contractAddress := "482be94ca327f1dd1d9857a5a212df091f44980f"
	//address := "qJ2HTPYoMF1DPBhgURjRqemun5WimD57Hy"
	//var tokenDecimal uint64 = 4

	unspent, err := tw.GetQRC20UnspentByAddress(contractAddress, address, tokenDecimal, isTestNet)
	if err != nil {
		t.Errorf("GetUnspentByAddress failed unexpected error: %v\n", err)
	}

	if err != nil {
		t.Errorf("strconv.ParseInt failed unexpected error: %v\n", err)
	}else {
		t.Logf("QRC20Unspent %s: %s = %v\n", contractAddress, address, unspent)
	}
}

func Test_AmountTo32bytesArg(t *testing.T){
	var amount int64= 100000000
	bytesArg, err := AmountTo32bytesArg(amount)
	if err != nil {
		t.Errorf("strconv.ParseInt failed unexpected error: %v\n", err)
	}else {
		t.Logf("hexAmount = %s\n", bytesArg)
	}
}

func Test_QRC20Transfer(t *testing.T) {

	//contractAddress := "482be94ca327f1dd1d9857a5a212df091f44980f"
	//from := "qUaHAjfRLknMBuSsA5kBfkn9xLMDFc2FdV"
	//to := "qJ2HTPYoMF1DPBhgURjRqemun5WimD57Hy"
	//var tokenDecimal uint64 = 4

	contractAddress := "0x91a6081095ef860d28874c9db613e7a4107b0281"
	from := "qVT4jAoQDJ6E4FbjW1HPcwgXuF2ZdM2CAP"
	to := "qJRyTVtn1bUjeYDztupJzinnN7sn7nZms7"
	var tokenDecimal uint64 = 8

	//contractAddress := "2d6cd7b2ef6de7e3866e17650c7d9c53529c0239"
	//from := "qHdSjkNTqSF3sVmiMpzU7ujSgZ9EobiTki"
	//to := "qLuqv9A7vigsaD8vZDqk7aS8fNH3EtNZjf"
	//var tokenDecimal uint64 = 15

	gasPrice := "0.00000040"
	var gasLimit int64 = 250000
	var amount decimal.Decimal = decimal.NewFromFloat(100)

	result, err := tw.QRC20Transfer(contractAddress, from, to, gasPrice, amount, gasLimit, tokenDecimal, isTestNet)
	if err != nil {
		t.Errorf("QRC20Transfer failed unexpected error: %v\n", err)
	}else {
		t.Logf("QRC20Transfer = %s\n", result)
	}
}

func Test_GetTokenBalanceByAddress(t *testing.T) {
	contract := openwallet.SmartContract{
		Address:  "f2033ede578e17fa6231047265010445bca8cf1c",
		Symbol:   "QTUM",
		Name:     "QCASH",
		Token:    "QC",
		Decimals: 8,
	}
	addrs := []string{
		//"qVT4jAoQDJ6E4FbjW1HPcwgXuF2ZdM2CAP",
		//"qQLYQn7vCAU8irPEeqjZ3rhFGLnS5vxVy8",
		//"qMXS1YFtA5qr2UfhcDMthTCK6hWhJnzC47",
		//"qJq5GbHeaaNbi6Bs5QCbuCZsZRXVWPoG1k",
		"QVTQ8QaKXcEzsGX74JePUkge5K4r41Ey3v",
		"Qf2tU1GExXe8smNyoMgis9u1CLnQ2q1mam",
	}

	//contract := openwallet.SmartContract{
	//	Address: "482be94ca327f1dd1d9857a5a212df091f44980f",
	//}
	//addrs := []string{
	//	"qUaHAjfRLknMBuSsA5kBfkn9xLMDFc2FdV",
	//	"qJ2HTPYoMF1DPBhgURjRqemun5WimD57Hy",
	//}

	balanceList, err := tw.ContractDecoder.GetTokenBalanceByAddress(contract, addrs...)
	if err != nil {
		t.Errorf("get token balance by address failed, err=%v", err)
		return
	}

	//输出json格式
	//objStr, _ := json.MarshalIndent(balanceList, "", " ")
	//t.Logf("balance list:%v", string(objStr))

	for i:=0; i<len(balanceList); i++ {
		t.Logf("%s: %s\n",addrs[i], balanceList[i].Balance.ConfirmBalance)
	}
}