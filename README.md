# Hackattic - [Help me unpack](https://hackattic.com/challenges/help_me_unpack)

Receive bytes and extract some numbers from those bytes.

## Usage

Assuming you have `make` installed in your system and `ACCESS_TOKEN` environment variable is set.

> [!WARNING]
> Not all functionality is tested. Doesn't support `.env`.

### Test

```bash
make test
```

#### Format

Uses [gofmt](https://pkg.go.dev/cmd/gofmt)

```bash
make fmt
```

#### Lint

Uses [golint](https://pkg.go.dev/golang.org/x/lint/golint)

```bash
make lint
```

#### Vet

Uses [go vet](https://pkg.go.dev/cmd/vet) and [shadow](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/shadow)

```bash
make vet
```

#### CI/CD

Uses [golangci-lint](https://golangci-lint.run/)

```bash
make ci
```

#### Build a binary

```bash
make build
```

#### Remove the binary

```bash
make clean
```

</br>

> [!TIP]
> If you don't want to install `make`, after setting the environment variable run:

```bash
go run main.go
# You can alo set the environment variable following way
ACCESS_TOKEN='your_token_here' go run main.go
```

## 🎉

![image](https://github.com/therin/hackattic/assets/86803100/3a00f8ba-2d38-4a58-bd55-03cbf79a21ff)
