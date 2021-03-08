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