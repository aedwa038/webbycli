[![CircleCI](https://dl.circleci.com/status-badge/img/gh/aedwa038/webbycli/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/aedwa038/webbycli/tree/main)
[![codecov](https://codecov.io/gh/aedwa038/webbycli/branch/main/graph/badge.svg?token=4FUOFXU93Y)](https://codecov.io/gh/aedwa038/webbycli)

# Webby CLI

webby cli is aa command line tool that will allow the user to query the [Merriam-Webster](https://dictionaryapi.com/products/json#sec-2.fl) Dictionary and receive a formatted response with the word definition


## Build
Webby comes with both a Makefile and a bazel script for building. Building with the Makefile is prefered.

After cloning the project you can build it with the following command:
```bash
make build
```

For packaging up webby for distribution run the following commands for tars
```
 make package
```

This will build webby for linux, windows and mac and packge the files up as TARS. 
TODO: add packing for deb files.

## Installation
After packaging, installing webby is simple as decompressing the tar file.
The tar file should contain two files

    * `webby`: the binary for execution needed
    * `config.yaml`: config file used to storing your Webster api key
Before using webby be sure to obtain a key from the [Merriam-Webster](https://dictionaryapi.com/products/json#sec-2.fl) and add to the configuration file.

## Usage
Invoke the tool for a definition
```bash
    webby -- -term=adapt --config=/path/to/configfile
```

**NB: This application is not tested in a real aws environment**
