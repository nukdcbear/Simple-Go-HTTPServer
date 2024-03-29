package httpserver

import (
	"dagger.io/dagger"
	// "dagger.io/dagger/core"
	// "universe.dagger.io/alpine"
	// "universe.dagger.io/bash"
    "universe.dagger.io/go"
)

dagger.#Plan & {
	client: {
		filesystem: {
			"./": read: {
				contents: dagger.#FS
                include: ["go.mod",
                    "go.sum",
                    "server.test.go",
                    "server.go",
                    "config.yaml"
                ]
			}
			"./output": write: contents: actions.build.output
		}
    }

    actions: {
        test: go.#Test & {
            source: client.filesystem."./".read.contents
            package: "./..."
        }

        build: go.#Build & {
            source: client.filesystem."./".read.contents
            arch: "amd64"
            os: "windows"
        }
    }
}