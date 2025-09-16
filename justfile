alias c := check
@check:
    deno check **/*.ts

alias b := build
@build: check
    mkdir -p dist/sudoku
    cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" lib/vendor/wasm_exec
    deno bundle index.html --outdir dist/sudoku
    # TODO: enable sw
    #bun build sw.ts --outdir dist/sudoku
    #echo "\n// $(git rev-parse HEAD) $(uuidgen)" >> dist/sudoku/sw.js # trigger reload
    deno run -A npm:@tailwindcss/cli -o dist/sudoku/output.css
    GOOS=js GOARCH=wasm go build -o dist/sudoku/main.wasm ./wasm
    cp public/favicon.svg manifest.json dist/sudoku

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
