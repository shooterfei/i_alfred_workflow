#!/usr/bin/env python3

# -*- coding=utf-8 -*-

import plistlib
import os
import glob
import shutil


reset_fields = [
        "GITHUB_ACCESS_TOKEN",
        "GITLAB_ACCESS_TOKEN",
        "GITLAB_HOST",
        "GITLAB_PORT",
        "PROTOCOL"
    ]

# 处理plist文件, 如果存在工作流则继续使用否则重置为默认配置
if __name__ == "__main__":
    info_plist = ""
    home = os.environ['HOME']
    files = glob.glob(f"{home}/Library/Application Support/Alfred" +
        "/Alfred.alfredpreferences/workflows/*/info.plist")

    for path in files:
        with open(path, "rb") as fp:
            pl = plistlib.load(fp)
        if pl["bundleid"] == "shooterfei":
            info_plist = path
            break
    if info_plist != "":
        print("copying")
        shutil.copy(info_plist, "../bin/")
    else:
        print("reset")
        with open("../bin/info.plist", "rb") as fp:
            pl = plistlib.load(fp)
        for (key, value) in pl["variables"].items():
            if key in reset_fields:
                pl["variables"][key] = ""
        with open("../bin/info.plist", "wb") as fp:
            plistlib.dump(pl, fp, fmt=plistlib.FMT_XML, sort_keys=False)
