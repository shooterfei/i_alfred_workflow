#!/usr/bin/env pdm run python

# -*- coding=utf-8 -*-

import requests
import json
import sys
import time


def data_handler(item):
    it = {
        'title': item["package"]["name"],
        'subtitle': "",
        'arg': 123,
        'mods': {
            'cmd': {
                'arg': 123,
                'subtitle': 'copy maven dependency to clipboard'
            },
            'alt': {
                'arg': 123,
                'subtitle': 'copy gradle dependency to clipboard'
            }
        }
    }
    return it


#  环境获取


search = sys.argv[1] if len(sys.argv) > 1 else ''

url = "https://www.npmjs.com/search"

params = {
    'q': search,
}

headers = {
    'x-spiferack': '1'
}
res = requests.get(url, headers=headers, params=params)



data = list(map(data_handler, res.json()['objects']))

# for item in data:
#     print(item['id'])

pac = {
    'items': data
}

# print(pac)
print(json.dumps(pac))


