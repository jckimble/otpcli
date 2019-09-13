# OtpCli 
[![pipeline status](https://gitlab.com/jckimble/otpcli/badges/master/pipeline.svg)](https://gitlab.com/jckimble/otpcli/commits/master)

A quick tool for totp tokens. Made cause other desktop authenticators bug me.

---
* [Install](#install)
* [Configuration](#configuration)
* [License](#license)

---

## Install
```sh
go install gitlab.com/jckimble/otpcli
```
### Rofi
```sh
rofi -show otp -modi otp:"otpcli code --rofi"
```

## Configuration
This project uses viper so it supports any type viper does including environment variables. It searches for '$HOME/.config/otpcli/config.*'. otpsecrets can be a plain text file in which if '$HOME/.config/otpcli/secrets.txt' exists a config file isn't needed but this is not recommended besides from testing
```yaml
---
otpsecrets: ~/.config/otpcli/secrets.asc
gpgkey: 0xDEADBEEF 
```

## License

Copyright 2019 James Kimble

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
