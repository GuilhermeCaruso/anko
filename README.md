# Anko
> Simple Application watcher

![GitHub](https://img.shields.io/badge/golang%20->=1.15.x-blue.svg) [![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-yellow)](https://img.shields.io/badge/License-Apache%202.0-yellow)


<p>
  <img src=".github/assets/anko_logo.png" width=150>
</p>

The Anko project aims to be a command line used to watch files of different extensions, executing actions previously configured, such as, restart, deletion, execution of scripts, among others.

The project so far has the following functionalities:

- Definition of execution language. [List of implemented languages](#implemented-languages)
- Declaration of extensions and files for watching.
- Independent configuration file

# Summary

- [Install](#install)
- [Guide](#guide)
  - [Anko File](#anko-file)
- [Example](#example)
- [Implemented Languages](#implemented-languages)
- [OS compatibility](#os-compatibility)
- [Author](#author)
- [License](#license)


# Install

### **If you have golang installed**

- To get Anko CLI

```sh
go get -u github.com/GuilhermeCaruso/anko
```

- If you prefer to create your own build using the flags you want, just clone this repository and run the golang build command

```sh
git clone git@github.com:GuilhermeCaruso/anko.git && \
cd anko && \
go <your_flags_here> build  
```

Once with the binary, remember to add it to the system PATH

> If you still don't know how to do it, [read here](https://superuser.com/questions/284342/what-are-path-and-other-environment-variables-and-how-can-i-set-or-use-them)!


### **If you don't have golang isntalled**

Soon we will make available the download the binaries...


# Guide

The anko project was initially developed to observe changes in golang projects and reload it on demand. During development, we noticed that the project could be used for the most diverse languages ​​and technologies (With the help of the community), because of that, it was necessary to add a configuration file, the Anko file

## Anko File

- **Localization**

The Anko file is a `.yaml` file written, preferably, at the root of the project you want to observe, but nothing prevents you from saving it wherever you want.

- **Nomenclature**

The project was developed to avoid as much as possible that the user has to adapt the project to suit him, however, the file name is the only information that must be kept as a standard.

The configuration file name must always be: `anko.yaml`

- **Properties**

> Root

| key | type | description |
|-|-|-|
|`application`| Application | Base of the anko file

> Application

|key |type| description|
|-|-|-|
|`root_path`| string | Path to the directory to be observed
|`exec_ath`| string | Path to application entry point
|`language`| string | Execution language. Check [here](#implemented-languages) if your favorite language has already been implemented
|`watch`| Watch | List of files and extensions to be observed. Initially only update events

> Watch

|key |type| description|
|-|-|-|
|`extensions`| []string | List of extensions that must be observed
|`files`| []string | List of files that must be observed




# Implemented languages

anko.yaml

```yaml
application:
  root_path: "."
  exec_path: "example/test.go"
  language: go
  watch:
    extensions: 
      - go
      - yaml
      - mod
    files:
      - README.md
```