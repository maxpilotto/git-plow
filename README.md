# git-plow

This script lets you download a sub directory from a git repository without downloading the whole thing.

- [Requirements](#requirements)
- [Installing](#installing)
- [Basic Usage](#basic-usage)
- [Building](#building)


## Requirements

+ Git 2.25
+ Go 1.23.3+ (for building/installing)

## Installing

```
go install github.com/maxpilotto/git-plow
```

## Basic Usage

**Fetching a sub directory**

```bash
$ git-plow <repo_url> <subdir_path>
```

E.g.
```bash
$ git-plow https://github.com/googlesamples/mlkit android/vision-quickstart
```

This will copy the content of the folder `vision-quickstart` in your working directory, without the need to download the entire repository.

You can also keep the original folder structure by using the `-k` flag, that will create the directories `mlkit/android/vision-quickstart` but only `vision-quickstart` will contain files.

## Building

```
./build.sh
```