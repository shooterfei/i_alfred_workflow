#!/usr/bin/env pdm run python

# -*- coding=utf-8 -*-

import requests
import json
import sys
import time


def data_handler(item):
    # item = col.defaultdict(lambda: '')
    g = item['g']
    a = item['a']
    # v = item['latestVersion']
    v = item.get('v', '') if not item.get('v', '') == '' else item['latestVersion']
    mvn = "<dependency>\n  <groupid>%s</groupid>\n  <artifactid>%s</artifactid>\n  <version>%s</version>\n</dependency>" % (g, a, v)
    gradle = "implementation('%s:%s:%s')" % (g, a, v)
    it = {
        'title': g + ":" + a + ":" + v,
        'subtitle': "update at %s \t count: %s" % (time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(item["timestamp"] / 1000)), item.get("versionCount", "0")),
        'arg': "test",
        'mods': {
            'cmd': {
                'arg': mvn,
                'subtitle': 'copy maven dependency to clipboard'
            },
            'alt': {
                'arg': gradle,
                'subtitle': 'copy gradle dependency to clipboard'
            }
        }
    }
    return it


#  环境获取


search = sys.argv[1] if len(sys.argv) > 1 else ''

url = "https://search.maven.org/solrsearch/select"

params = {
    'q': search,
    'start': 0,
    'rows': 20
}

res = requests.get(url, params=params)


data = list(map(data_handler, res.json()['response']['docs']))

# for item in data:
#     print(item['id'])

pac = {
    'items': data
}

# print(pac)
print(json.dumps(pac))
