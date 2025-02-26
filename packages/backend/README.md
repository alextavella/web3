# Backend

### Getting started

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

### Generate contract.go

```bash
abigen --abi=../solidity/build/MyContract.abi --bin=../solidity/build/MyContract.bin --pkg=contracts --out=pkg/contracts/mycontract.go
```

### Running

```bash
go run cmd/main.go
```
