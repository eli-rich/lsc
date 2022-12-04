# LSC

A basic tool to quickly create a LICENSE for your open source project.

- [LSC](#lsc)
  - [Installation](#installation)
    - [With Go](#with-go)
    - [From Release](#from-release)
    - [From Source](#from-source)
  - [Usage](#usage)
  - [Licenses](#licenses)
  - [Contributing](#contributing)

## Installation

### With Go

```sh
go install github.com/eli-rich/lsc
```

### From Release

- Download the release from the releases page.
- Add executable to your path.

### From Source

- Clone the repository.
  - `git clone https://github.com/eli-rich/lsc.git`
- Build the binary.
  - `cd lsc`
  - `go install .`

## Usage

Generate an MIT license.

```sh
lsc "Eli Richardson" 2022 MIT
```

Generate a GPL-3.0 license.

```sh
lsc "Eli Richardson" 2022 GPL
```

Generate an MPL license.

```sh
lsc "Eli Richardson" 2022 MPL
```

Generate a BSD-3 license.

```sh
lsc "Eli Richardson" 2022 BSD3
```

## Licenses

- MIT
- APACHE
- BSD2
- BSD3
- GPL
- MPL

## Contributing

Pul lrequests are encouraged! Feel free to add any open source licenses that are missing.

- Add the license to the root of the project.
- Make sure — if needed — to include {{.Year}} and {{.Name}} in the license.
- Update `main.go` to include the new license.
- Submit the PR!
