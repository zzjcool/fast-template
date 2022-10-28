# Fast Template

Fast rendering with parameters and templates, The rendering engine uses go-template

go-template: https://pkg.go.dev/text/template#hdr-Examples

sprig: http://masterminds.github.io/sprig/

## Example

download:

```bash
curl https://github.com/zzjcool/fast-template/releases/download/latest/ft -LJo ft
chmod +x ft
```

Render the `example/README.template` template with the values of `example/value.yaml` and output to `example/README.md`

```bash
./ft build -t "example/README.template" -v "example/value.yaml" -o "example/README.md"
```
