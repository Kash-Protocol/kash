Kashd
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/Kash-Protocol/kashd)

## Kash: Pioneering the Purest Form of Decentralized Stablecoin Payments

⚠️ **IMPORTANT: This project is currently in heavy development and may be highly unstable. It is recommended to wait for the stable testing version release before engaging in mining.**

Kashd is the reference full node implementation of Kash, a groundbreaking cryptocurrency that aims to establish the most pure form of decentralized stablecoin payments. Building upon the rapid and efficient Kaspa protocol, Kash integrates the Djed stablecoin protocol to create a robust and versatile digital currency ecosystem.

### Core Features of Kash (KSH)

- **A Vision for Decentralized Stability**: Kash's primary goal is to revolutionize the stablecoin market by introducing KUSD – a stable digital currency envisioned to be the epitome of decentralized stability.

- **Integration of Djed Stablecoin Protocol**: The Djed stablecoin protocol is at the heart of Kash's design, enabling the creation of KUSD and KRV, a reserve currency. This integration not only enhances the ecosystem's stability but also its utility. For a detailed understanding of the Djed stablecoin protocol and its role in the Kash ecosystem, refer to [the Djed Protocol](https://eprint.iacr.org/2021/1069.pdf).

- **ASIC Resistance with RandomX**: Kash adopts the RandomX algorithm, an ASIC-resistant proof-of-work mechanism. This choice reflects Kash's commitment to maintaining a decentralized and egalitarian mining landscape.

- **Rapid and Secure Transactions**: Leveraging the Kaspa protocol, Kash inherits its renowned sub-second block times and instant confirmations, ensuring rapid and secure transactions.

- **Adherence to PHANTOM Protocol**: Kash upholds the core attributes of Kaspa, including its reliance on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a sophisticated generalization of Nakamoto consensus.

## Requirements

Go 1.18 or later.

## Installation

### Build from Source

- Install Go according to the installation instructions here:
  [http://golang.org/doc/install](http://golang.org/doc/install)

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Clone the Kashd repository:

```bash
$ git clone https://github.com/Kash-Protocol/kashd
$ cd kashd
```

- Before installing Kashd, ensure that the `randomx` library is properly built. Navigate to the `util/randomx` directory within the project and execute the build process for `randomx`:

```bash
$ cd util/randomx
$ ./build.sh
$ cd ../../
```

- Set the necessary environment variables for CGO and Go build flags:

```bash
$ export CGO_ENABLED=1
$ export CGO_LDFLAGS="-L$(pwd)/util/randomx/lib -lrandomx -lm"
$ export GOFLAGS='-ldflags=-extldflags="-static"'
```

- After setting the environment variables and successfully building `randomx`, install Kashd and all its dependencies:

```bash
$ go install . ./cmd/...
```

- Kashd (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.

## Getting Started

Kashd operates with minimal configuration for basic operations. Advanced users can tweak various settings for optimized performance.

```bash
$ kashd
```

## Community and Support

Join our Discord server for community discussions, support, and updates: https://discord.gg/YNYnNN5Pf2

## Issue Tracker

Report issues and track progress on our [GitHub issue tracker](https://github.com/Kash-Protocol/kashd/issues). View issue priorities at https://github.com/orgs/kaspanet/projects/4

## Documentation

Access our comprehensive [documentation](https://github.com/kaspanet/docs) for detailed information about Kashd. The documentation is continuously updated to reflect the latest developments.

## License

Kashd is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
