# Solidity

### Getting start

```bash
pnpm add -D hardhat
pnpm hardhat init
```

### Compile

```bash
pnpm hardhat compile
```

### Generate BIN

```bash
cat artifacts/contracts/MyContract.sol/MyContract.json | jq '.abi' > build/MyContract.abi
cat artifacts/contracts/MyContract.sol/MyContract.json | jq -r '.bytecode' > build/MyContract.bin
```
