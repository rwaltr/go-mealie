# go-mealie

A dirt simple go CLI client and badly written go library

Enjoy this alpha release


# Config

`mealie` can be configured using a yaml file or by using the `MEALIE_URL` and `MEALIE_TOKEN` env

```
---
url: https://mymealie.com
token: mysupertoken
```
The config file can be in the following locations

```
$XDG_CONFIG_DIR/mealie/config
$HOME/.config/mealie/config
.
```

```
Usage:
  go-mealie [command]

Available Commands:
  completion      generate the autocompletion script for the specified shell
  deleteRecipe    A brief description of your command
  help            Help about any command
  listAllRecipies A brief description of your command
  scrape          A brief description of your command
  view            A brief description of your command

Flags:
  -h, --help   help for go-mealie
```
