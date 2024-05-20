package across

// https://github.com/czbag/zksync/blob/main/data/abi/dmail/abi.json
///go:generate abigen --abi abi.json --pkg across --type Storage --out across.go
