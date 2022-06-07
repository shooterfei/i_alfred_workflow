#!/usr/bin/env python3

# -*- coding=utf-8 -*-

import plistlib
import os
import glob
# import shutil

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
    plist_files = glob.glob(f"{home}/Library/Application Support/Alfred" +
        "/Alfred.alfredpreferences/workflows/*/info.plist")

    for path in plist_files:
        with open(path, "rb") as fp:
            pl = plistlib.load(fp)
        if pl["bundleid"] == "shooterfei":
            info_plist = path
            break

    if info_plist != "":
        print("moving")
        target = os.path.dirname(info_plist)
        # 涉及到文件删除, 不要在不理解的情况下进行修改导致误删重要文件, 甚至可以事先配置好时间机器
        # shutil.rmtree(target)
        # os.mkdir(target)
        os.system(f"rm -rf '{target}'/*")
        # 会丢失权限,改为命令解压
        # zFile = zipfile.ZipFile("../bin/tools.alfredworkflow", "r")
        # for fileM in zFile.namelist():
        #     zFile.extract(fileM, target)
        os.system(f"rsync -avzP ../bin/ '{target}/'")
        # os.system(f"unzip ../bin/tools.alfredworkflow -d '{target}'")
    else:
        print("install")
        os.system("make package")
        os.system("open ../bin/tools.alfredworkflow")
