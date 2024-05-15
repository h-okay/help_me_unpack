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

```bash
make fmt
```

#### Lint

```bash
make lint
```

#### Vet

```bash
make vet
```

#### CI/CD

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
```

> You can alo set the environment variable following way:

```bash
ACCESS_TOKEN='your_token_here' go run main.go
```

## 🎉

![image](https://github.com/therin/hackattic/assets/86803100/3a00f8ba-2d38-4a58-bd55-03cbf79a21ff)
