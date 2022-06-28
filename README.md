# json2yaml
Convert json configuration to yaml

# Usage

Get help:
```
./json2yaml -h
```

Convert `target.json` to `result.yml`:
```
./json2yaml
```

Convert `foo.json` to `foo.yml`:
```
./json2yaml -f "foo.json"
```

Convert `foo.json` to `bar.yml`:
```
./json2yaml -f "foo.json" -o "bar.yml"
```

**NOTE:**
By default it'll convert and add `header.json` first to the yml. you can modify it yourself, or, you can set `-H ""` to ignore it.
