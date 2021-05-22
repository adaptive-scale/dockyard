# Quickstart

Dockyard is a static site generator that takes *Markdown* and uses predefined layouts to generate the static sites.
It is valuable for generating documentation, information sites, simple landing pages etc.

### Prerequisites

- Dockyard binary
- Understanding of Markdown language

### Instructions

Here we download the dockyard binary, then download some sample markdown files, run the generator on those files and serve them. 

*This quickstart is for Mac*

##### Download dockyard

```$xslt
curl -L -o dockyard https://github.com/adaptive-scale/dockyard/releases/download/v0.1.1/dockyard_macos
chmod +x dockyard
mv dockyard /usr/local/bin
dockyard -h
```

#### Download sample Markdown pages

```
curl -L -o data.tar.gz curl -L -o dockyard https://github.com/adaptive-scale/dockyard/releases/download/v0.1.1/docs.zip
unzip docs.zip
```

#### Run it on .md files

```$xslt
dockyard --location docs --branding myowncompany --serve
```

To go, `http://localhost:10000`