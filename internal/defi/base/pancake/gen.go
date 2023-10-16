package pancake

//go:generate abigen --abi router.json --pkg pancake --type router --out router.go
//go:generate abigen --abi pool.json --pkg pancake --type pool --out pool.go
//go:generate abigen --abi factory.json --pkg pancake --type factory --out factory.go
//go:generate abigen --abi quoter.json --pkg pancake --type quoter --out quoter.go
