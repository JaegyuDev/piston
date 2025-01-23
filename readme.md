# Piston CLI Tool

`piston` is a command-line utility for interacting with Mojang's Piston API. It allows you to list available Minecraft versions, download specific versions of the server jar, and manage snapshots.

## Features

- List available Minecraft versions
- Download specific Minecraft server versions by version ID
- Support for snapshots

## Installation

To install the `piston` CLI tool, use the following command:

```
go install github.com/JaegyuDev/piston@latest
```

This will fetch and install the tool in your `$GOPATH/bin` or `$HOME/go/bin` directory. These will have to be added to your $PATH.

## Usage

### List all available versions

```
piston list
```

This command will list all available versions from Mojang's Piston API. You can filter the list using flags like `--allow-snapshots` to include snapshot versions.

### Download a specific version

```
piston get <version>
```

Download the server jar for the specified version. By default, the jar will be downloaded to the current directory.

### Download a version to a specific path

```
piston get <version> -o <path>
```

You can specify a custom path using the `-o` flag. If no path is provided, the jar will be downloaded to the current directory.

## Example

To download the latest version:

```
piston get latest
```

To download the latest snapshot:

```
piston get -s latest
```

To list all available versions:

```
piston list
```

To download a specific version (e.g., `1.21.4`):

```
piston get 1.21.4
```

To download the same version but save it to a specific folder:

```
piston get 1.21.4 -o /path/to/folder
```
