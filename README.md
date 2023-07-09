# multirun
> (under development)

Run multiple terminal commands at the same time.

## Usage

### Commandline

```bash
mun --color both \
    sass-preprocessor:"npx sass -w --no-source-map src/styles:static/styles" \
    ts-compiller:"npx tsc -w"
```

### Configuration file

```yaml
# multirun.yaml

commands:
    sass-preprocessor: npx sass -w --no-source-map src/styles:static/styles
    ts-compiller: npx tsc -w
settings:
    color: both
```

### Output

```
ts-compiller      | [4:03:07 PM] Starting compilation in watch mode...
sass-preprocessor | Sass is watching for changes. Press Ctrl-C to stop.
sass-preprocessor | 
ts-compiller      | [4:03:09 PM] Found 0 errors. Watching for file changes.
```

## Settings

### `color`
> Optional. Default: `both`

Select when the command should use colors in the terminal.

Possible values are:
 - `"none"`: the output will contain no colors at all;
 - `"mun"`: the output will contain colors only from the mun;
 - `"command"`: the output will contain colors only from the running command;
 - `"both"`: the output will contain colors from both mun and running command.