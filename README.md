## 'local' lyra workflows

This repo contains a user directory for Lyra workflows and hiera data.

It's got some sample workflows that are works-in-progress and demoable, but which shouldn't be checked in to the main codebase (yet, or maybe ever...)

To use it, clone the repo and either:

1. If running from source or homebrew, just `cd lyra-local` as you run lyra. Lyra looks for a `hiera.yaml` and `workflows/` relative to your current working directory, so these will be found.
2. If running from docker, bind-mount the directory into the `local` mountpoint that exists in the container build:

```
docker run -it \
--mount type=bind,src=$HOME/lyra-local,dst=/src/lyra/local \
lyraproj/lyra:latest /bin/ash
```

Please submit pull requests if you get additional workflows working, or just @ me and I'll add you as a collaborator!

--eric0
