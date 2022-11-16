# yuyufetch

A command-line my information tool

## How to install yuyufetch

```bash
go install github.com/Hiroya-W/yuyufetch/cmd/yuyufetch@latest
```

## How to build yuyufetch

`yuyufetch` can be compiled with the following command.

```
make build
```

Run `yuyufetch`.

```
./bin/yuyufetch
```

## Convert image to ASCII art

Using [jp2a](https://github.com/Talinx/jp2a):

```bash
docker run -t --rm -e COLORTERM="$COLORTERM" -v "$(pwd)":/app talinx/jp2a --colors --size=36x18 --chars="..,;McM0MNWM" yuyu.png
```
