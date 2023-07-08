# multirun

Run multiple terminal commands at the same time

## Usage

### Commandline

```bash
mun \
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

| Argument | Description                               | Type                             |
| -------- | ----------------------------------------- | -------------------------------- |
| `color`  | Select when the command should use colors | `none`, `mun`, `command`, `both` |