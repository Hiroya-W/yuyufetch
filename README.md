# yuyufetch

A command-line self-introduction tool.

![image](https://user-images.githubusercontent.com/43127622/202208641-bae0179c-f681-4aff-bd52-f093feb5bcde.png)

Inspired by [Neofetch](https://github.com/dylanaraps/neofetch), [perufetch](https://github.com/TadaTeruki/perufetch).

## How to install yuyufetch

```bash
go install github.com/Hiroya-W/yuyufetch@latest
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
