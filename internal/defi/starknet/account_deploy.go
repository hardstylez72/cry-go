package starknet

//
//// site https://starkscan.co/tx/0x6ea856e3209e04658d99729a8f485131b0944b64d5145b7924ae5a1f59f3b19
//func (c *Client) DeployAccount(ctx context.Context, req *Account) error {
//
//	account, err := caigo.NewGatewayAccount(req.PrivateKey, req.PublicKey, c.GWP)
//	if err != nil {
//		return err
//	}
//
//	data := []types.FunctionCall{{
//		ContractAddress:    types.Hash{},
//		EntryPointSelector: "",
//		Calldata:           nil,
//	}}
//
//	//account.Call()
//
//	deploy, err := c.GW.DeployAccount(ctx, types.DeployAccountRequest{
//		MaxFee:              nil,
//		Version:             1,
//		Signature:           nil,
//		Nonce:               nil,
//		ContractAddressSalt: req.PublicKey,
//		ConstructorCalldata: []string{req.PublicKey},
//	})
//	if err != nil {
//		return err
//	}
//
//	if err := c.WaitForTransaction(deploy.TransactionHash); err != nil {
//		return err
//	}
//
//	tx, err := c.GW.Transaction(ctx, gateway.TransactionOptions{TransactionHash: deploy.TransactionHash})
//	if err != nil {
//		fmt.Println("can't fetch transaction data:", err)
//		os.Exit(1)
//	}
//	types.StrToFelt(tx.Transaction.ContractAddress)

//	return nil
//}
//
//func (c *Client) WaitForTransaction(txHash string) error {
//	acceptedOnL2 := false
//	var receipt *gateway.TransactionReceipt
//	var err error
//	fmt.Println("Polling until transaction is accepted on L2...")
//	for !acceptedOnL2 {
//		_, receipt, err = c.GW.WaitForTransaction(context.Background(), txHash, 5, 10)
//		if err != nil {
//			fmt.Println(receipt.Status)
//			return fmt.Errorf("Transaction Failure (%s): can't poll to desired status: %s", txHash, err.Error())
//		}
//		fmt.Println("Current status : ", receipt.Status)
//		if receipt.Status == types.TransactionAcceptedOnL2 {
//			acceptedOnL2 = true
//		}
//	}
//	return nil
//}
