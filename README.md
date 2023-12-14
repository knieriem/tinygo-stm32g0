This repository contains files adding initial support for STM32G0 microcontrollers to TinyGo.

The implementation is based on existing ones for L0 and L5,
using a recent revision (v0.30.0+) from the dev branch of TinyGo. It is a work in progress;
once remaining problems are resolved, it could be integrated into TinyGo.
So far UART, GPIO, I²C and SPI have been used successfully on an G031K8
on a NUCLEO-32 board, and on a G030K6 on a proprietary board.

[TinyGo]: https://github.com/tinygo-org/tinygo

To add these files to TinyGo, run

```sh
make TINYGO=path-to-tinygo dist
```

which will produce a tar file that may be untared within the root of
a local TinyGo installation.
