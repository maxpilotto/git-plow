#!/usr/bin/env python3

import argparse
import subprocess
import sys
import re

ERR_NO_URL = 1
ERR_MALFORMED_URL = 2
RE_REPO = re.compile(r'(?P<base_url>(?:(?P<protocol>https?)://)?(?:www\.)?(?P<domain>.+?)(?:/(?P<user>[A-Za-z0-9-_]+))(?:/(?P<repo>[A-Za-z0-9-_]+)))(?:/tree/(?P<treeref>[A-Za-z0-9-_.]+))?(?:/(?P<subdir>[A-Za-z0-9-_/]+))?')

def run(commands):
    if not isinstance(commands, list):
        commands = [commands]

    if args.debug:
        print('\nCommands: ')
        print('\n'.join(commands))

    process = subprocess.Popen(' && '.join(commands), shell=True, stdout=subprocess.PIPE)
    process.wait()

args = argparse.ArgumentParser()
args.add_argument('-s', '--subdir', help='Specify a subdirectory. If not defined it will be extracted from the url', action='store', type=str)
args.add_argument('-r', '--ref', help='Specify a branch/tag or any tree ref. If not defined it will be extracted from the url. Defaults to branch "main"', action='store', type=str)
args.add_argument('-k', '--keep', help='Keep originial structure and do not elevate the subdir', action='store_true')
args.add_argument('-d', '--debug', help='Show debug info', action='store_true')
args.add_argument('url', help='The url of the repository to clone')
args = args.parse_args()

url = args.url

if url.endswith('/'):
    url = url[:-1]

repo_match = RE_REPO.match(url)

if not repo_match:
    print('Malformed URL')
    sys.exit(ERR_MALFORMED_URL)

base_url = repo_match.group('base_url')
repo_name = base_url[base_url.rfind('/') + 1:]

if args.subdir:
    subdir = args.subdir
else:
    subdir = repo_match.group('subdir')

if args.ref:
    tree_ref = args.ref
else:
    tree_ref = repo_match.group('treeref')

if not subdir:
    print('No subdir specified')
    print('Doing regular clone')

    cmds = [f'git clone {base_url}']

    if tree_ref and tree_ref:
        cmds.append(f'cd {repo_name}')
        cmds.append(f'git checkout {tree_ref}')
    
    run(cmds)

    sys.exit(0)

subdir = subdir.removeprefix('/').removesuffix('/')
depth = len(subdir.split('/'))
subdir_last_segment = subdir.rfind('/')
dest_folder = subdir[subdir_last_segment + 1:] if subdir_last_segment != -1 else subdir
cmds = []

cmds.append(f'git clone -n --depth={depth} --filter=tree:0 {base_url}')
cmds.append(f'cd {repo_name}')
cmds.append(f'git sparse-checkout set --no-cone {subdir}')
cmds.append('git checkout')

if not args.keep:
    cmds.append(f'cd ..')
    cmds.append(f'cp -R {repo_name}/{subdir} .')
    cmds.append(f'rm -rf {repo_name}')

run(cmds)