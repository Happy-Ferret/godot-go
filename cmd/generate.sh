#!/bin/bash
export PKG_PATH="$GOPATH/src/github.com/shadowapex/godot-go"
export GODOT_DOC_PATH=$PKG_PATH/doc/doc/classes

# Check if the "doc" folder is present where we can fetch all of the godot documentation
# to include it as Go documentation as part of class generation.
if [ -d $PKG_PATH/doc ]; then
    echo "Godot documentation found. Pulling the latest changes..."
    cd "$PKG_PATH/doc"
    git pull origin master
else
    echo "Godot documentation not found. Pulling the latest changes..."
    mkdir "$PKG_PATH/doc"
    cd "$PKG_PATH/doc"
    git init
    git remote add -f origin https://github.com/godotengine/godot.git
    git config core.sparseCheckout true
    echo "doc/classes/" >> "$PKG_PATH"/doc/.git/info/sparse-checkout
    git pull origin master
fi

# Check if "godot" is present, generate our JSON.
which godot
if [ "$?" == "0" ]; then
    echo "Generating godot_api.json from godot..."
    godot --gdnative-generate-json-api $PKG_PATH/cmd/generate/templates/godot_api.json --no-window
else
    echo "The 'godot' binary was not found in \$PATH. Using previously generated godot_api.json in repository..."
fi

# Cleaning old classes
if [[ "$ONLY" == "classes" || "$ONLY" == "" ]]; then
	echo "Cleaning old generated classes..."
	rm -rf "$PKG_PATH"/godot/*.gen.go
fi

# Run code generation
go run "$PKG_PATH"/cmd/generate/main.go
