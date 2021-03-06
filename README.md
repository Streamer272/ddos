# DDOS
[![License](https://img.shields.io/github/license/Streamer272/ddos)](https://github.com/Streamer272/ddos/blob/master/LICENSE)
[![Stars](https://img.shields.io/github/stars/Streamer272/ddos)](https://img.shields.io/github/stars/Streamer272/ddos)

Simple DDOS application to bring down your nemesis' website

## Usage
```bash
ddos --address www.mynemisis.com:443 --output mynemesis-ddos.log --message "HELLO SUCKER" --max-retry-count 100
```

| Option (short) |    Option (long)    |                        Description                        |
|:--------------:|:-------------------:|:---------------------------------------------------------:|
|      `-d`      |      `--delay`      |                       Request delay                       |
|      `-r`      | `--max-retry-count` |                    Maximum retry count                    |
|      `-R`      |  `--request-count`  |                       Request count                       |
|      `-a`      |     `--address`     |               Your nemesis' website address               |
|      `-m`      |     `--message`     |                  Custom message to send                   |
|      `-o`      |     `--output`      |                   Output log file path                    |
|      `-l`      |    `--log-level`    |          Log level (NONE / ERROR / WARN / INFO)           |
|      `-H`      |      `--http`       | Use HTTP message (only when `--message` is not specified) |
|      `-i`      |  `--ignore-error`   |       Ignore errors; not terminate program on error       |
|      `-N`      |    `--no-color`     |                Display non-colored output                 |

## Requirements
- Computer running Linux
- Go compiler
- Git

## Installation
- Clone repository using git - `git clone https://github.com/Streamer272/ddos.git`
- Move into repository folder - `cd ddos`
- Build Go - `go build`
- Add to `/usr/bin` [Optional] - `cp ./ddos /usr/bin/ddos`

## Run in docker
```bash
docker run -it --rm -v $(pwd):/home/logs streamer272/ddos --address www.mynemesis.com:443 --output /home/logs/ddos.log --message "HELLO SUCKET" --max-retry-count 100
```

## License
This project is licensed under [MIT](https://github.com/Streamer272/ddos/blob/master/LICENSE) license.
