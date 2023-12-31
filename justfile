build-cli:
    cd cli && go build -o ./build/s

build-packages:
    cd packages/node && npm run build

publish-packages-preview:
    just build-packages
    cp packages/node/package.json packages/node/README.md packages/node/dist
    cd packages/node/dist && npm publish --dry-run

publish-packages:
    just build-packages
    cp packages/node/package.json packages/node/README.md packages/node/dist
    cd packages/node/dist && npm publish --access public
