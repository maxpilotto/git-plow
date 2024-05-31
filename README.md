# git-plow

This script lets you download a sub directory from a git repository without downloading the whole thing.

## Requirements

+ Git 2.25
+ Python 3

## Basic Usage

Let's say you want to download the sub directory `/android/vision-quickstart` from the repo `https://github.com/googlesamples/mlkit`.

You can navigate to the directory from your browser, copy the url and just paste it like this:

```bash
$ git-plow https://github.com/googlesamples/mlkit/android/vision-quickstart

# This is also a valid url
$ git-plow https://github.com/googlesamples/mlkit/tree/main/android/vision-quickstart
```

This will download the folder `vision-quickstart` and put it in your working directory.

You can also keep the original folder structure by using the `-k` flag, that will create the directories `mlkit/android/vision-quickstart` but only `vision-quickstart` will contain any file.

You can specify a sub directory and tree reference using the arguments as well:

```bash
$ git-plow https://github.com/googlesamples/mlkit -s android/vision-quickstart -t dev
```