alias c := check
@check:
    deno check **/*.ts

alias b := build
@build: check
    mkdir -p dist
    cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" lib/vendor/wasm_exec
    deno bundle index.html --outdir dist
    deno bundle sw.ts --outdir dist
    echo "\n// $(git rev-parse HEAD) $(date)" >> dist/sw.js # trigger reload
    deno run -A npm:@tailwindcss/cli -o dist/output.css
    GOOS=js GOARCH=wasm go build -o dist/main.wasm ./wasm
    cp public/favicon.svg manifest.json dist

alias s := serve
@serve: build
    #!/usr/bin/env bash
    cd dist
    python3 -m http.server 8080

alias f := fmt
@fmt:
    deno run -A npm:prettier -w **/*.ts
    go fmt ./...

alias t := test
@test:
    echo "Running go tests"
    go test ./...
