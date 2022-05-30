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

# 更新此目录下info.plist 文件为 当前使用的plist 文件
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
        shutil.copy(info_plist, "../")
