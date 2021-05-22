# Detailed Overview


#### Downloading Dockyard for Linux

**Download dockyard**

```$xslt
curl -L -o dockyard https://github.com/adaptive-scale/dockyard/releases/download/v0.1.1/dockyard_linux
chmod +x dockyard
sudo mv dockyard /usr/bin
dockyard -h
```

*Download sample Markdown pages**

```
curl -L -o data.tar.gz curl -L -o dockyard https://github.com/adaptive-scale/dockyard/releases/download/v0.1.1/docs.zip
unzip docs.zip
```

#### Downloading Dockyard for MacOS

**Download dockyard**

```$xslt
curl -L -o dockyard https://github.com/adaptive-scale/dockyard/releases/download/v0.1.1/dockyard_macos
chmod +x dockyard
mv dockyard /usr/local/bin
dockyard -h
```

**Download sample Markdown pages**

```
curl -L -o docs.zip https://github.com/adaptive-scale/dockyard/releases/download/v0.1.1/docs.zip
unzip docs.zip
```

**Run it on .md files**

```$xslt
dockyard --location docs --branding myowncompany --serve
```

Go to http://localhost:10000

#### Downloading Dockyard for Windows

Go to [here](https://github.com/adaptive-scale/dockyard/releases/download/v0.1.1/dockyard.exe)

**Run it on .md files**

Go to directory where you downloaded the binary.

Open your Command prompt. Run - 

```$xslt
dockyard.exe --location docs --branding myowncompany --serve
```
Go to http://localhost:10000

### Building side menu

You can build hierarchical directory structure that automatically detects the files and generates the html files.

***For instance***

You can have a directory called  - `doc`. To add a menu item create a file `<sequence>_<menu_item>.md` under the directory `doc`. 
The menu items are ordered by sequence number and then lexicologically.
A `<sequence>` could be something like `001` or anything that could be used to sort the menu item.

The file could look something as follows - 

```
doc
doc/001_first_item.md
doc/002_second_item.md
```

### Generating static pages


### Options


```$xslt

  -branding string
        branding of documentation (default "Acme")
  -location string
        location of documentation
  -serve
        generate and serve the documentation
  -theme string
        only default supported at this point (default "default")
  -watch
        watch for document change and regenerate when changed

```
