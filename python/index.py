#!/usr/bin/env /usr/local/bin/pdm run python
# -*- coding=utf-8 -*-

import json
import os
import sys

import requests


def data_handler(item):
    it = {
        "title": item["path_with_namespace"],
        "subtitle": item["description"],
        "icon": "./gl.png",
        "icontype": "filetype",
        "arg": item["web_url"],
        "mods": {
            "cmd": {
                "arg": item["ssh_url_to_repo"],
                "subtitle": "copy ssh to clipboard",
            },
            "alt": {
                "arg": item["http_url_to_repo"],
                "subtitle": "copy http to clipboard",

            },
        },
    }
    return it


#  环境获取
gitlab_host = os.getenv("GITLAB_HOST") or "192.168.2.30"
gitlab_port = os.getenv("GITLAB_PORT") or "9122"
token = os.getenv("GITLAB_ACCESS_TOKEN") or "sMTC-iDELEBY9iCm7qKb"
protocol = os.getenv("PROTOCOL") or "http"

search = sys.argv[1] if len(sys.argv) > 1 else ""

url = protocol + "://" + gitlab_host + ":" + gitlab_port + "/api/v4/projects"

headers = {"PRIVATE-TOKEN": token}
params = {"search": search}

res = requests.get(url, headers=headers, params=params)


data = list(map(data_handler, res.json()))

pac = {"items": data}

print(json.dumps(pac))
