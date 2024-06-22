Rust dev container 설정

```json title:rust_devContainer
// For format details, see https://aka.ms/devcontainer.json. For config options, see the

// README at: https://github.com/devcontainers/templates/tree/main/src/rust

{

    "name": "Rust",

    // Or use a Dockerfile or Docker Compose file. More info: https://containers.dev/guide/dockerfile

    //"image": "mcr.microsoft.com/devcontainers/rust:1-1-bullseye",

    "build":{

        "dockerfile": "Dockerfile", // Dockerfile path

        "context": ".."

    },

    "customizations": {

        "vscode": {

            "extensions": [

                "amazonwebservices.aws-toolkit-vscode",

                "ms-azuretools.vscode-docker"

            ],

            "settings": {

                "rust-analyzer.cargo.autoreload": true

            }

        }

    }

  

    // Use 'mounts' to make the cargo cache persistent in a Docker Volume.

    // "mounts": [

    //  {

    //      "source": "devcontainer-cargo-cache-${devcontainerId}",

    //      "target": "/usr/local/cargo",

    //      "type": "vo
    lume"

    //  }

    // ]

  

    // Features to add to the dev container. More info: https://containers.dev/features.

    // "features": {},

  

    // Use 'forwardPorts' to make a list of ports inside the container available locally.

    // "forwardPorts": [],

  

    // Use 'postCreateCommand' to run commands after the container is created.

    // "postCreateCommand": "rustc --version",

  

    // Configure tool-specific properties.

    // "customizations": {},

  

    // Uncomment to connect as root instead. More info: https://aka.ms/dev-containers-non-root.

    // "remoteUser": "root"

}
```
도커파일 설정
``` Dokerfile
FROM mcr.microsoft.com/devcontainers/rust:1-1-bullseye

RUN apt-get update
```