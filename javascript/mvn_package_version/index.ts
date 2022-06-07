import { Alfy, Item } from "..";
import * as cheerio from "cheerio";

const runTask = (groupId: string, artifactId: string): void => {
    let url = `https://mvn.coderead.cn/version?groupId=${groupId}&artifactId=${artifactId}`;
    fetch(url)
        .then((res) => res.text())
        .then((res) => {
            const alfy = new Alfy();
            const $ = cheerio.load(res);
            $('tr[onclick="doFold($(this))"]').each((i, elem) => {
                const vHtml = $(elem).children();

                const version = $(vHtml.get(0)).html();
                const downloadCount = $(vHtml.get(2)).html();
                const time = $(vHtml.get(3)).html();
                const temp: Item = {
                    title: version!,
                    arg: i.toString(),
                    subtitle: `update by ${time} \t downloadCount: ${downloadCount}`,
                    mods: {
                        ctrl: {
                            arg: `<dependency>\n\t<groupid>${groupId}</groupid>\n\t<ardifactid>${artifactId}</artifactid>\n\t<version>${version}</version></dependency>`,
                            subtitle: "copy for Maven",
                        },
                        cmd: {
                            arg: `implementation '${groupId}:${artifactId}:${version}'`,
                            subtitle: "copy for Gradle",
                        },
                        alt: {
                            arg: "implementation('${groupId}:${artifactId}:${version}')",
                            subtitle: "copy for Kotlin",
                        },
                    },
                };
                alfy.addItem(temp);
            });
            console.log(JSON.stringify(alfy));
        });
};

export default runTask;
