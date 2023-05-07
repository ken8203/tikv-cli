# tikv-cli

CLI tool for TiKV with an interactive shell.


## Usage

 You can enter the interactive shell by root command.

```
Usage:
  tikv-cli [flags]
  tikv-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a key
  get         Get a key
  help        Help about any command
  put         Put a key
  ttl         Get the TTL of a key
  version     Print the version number of tikv-cli

Flags:
  -a, --api-version string   API version. v1/v1ttl/v2 (default "v2")
      --debug                Debug determines whether to enable logging in tikv/client-go
      --help                 Help for tikv-cli
  -h, --host string          PD host address (default "localhost")
  -m, --mode string          Client mode. raw/txn (default "txn")
  -p, --port string          PD port (default "2379")

Use "tikv-cli [command] --help" for more information about a command.
```

## Playground

Before using `tikv-cli`, you are gonna provision a TiKV cluster.

### With docker-compose

```
docker-compose up -d
```

### With [TiUP](https://github.com/pingcap/tiup)

```
tiup playground --mode tikv-slim --kv 2 --kv.config ./hack/config/tikv.toml
```
