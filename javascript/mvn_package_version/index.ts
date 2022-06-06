// import { Alfy, Item } from "..";
import * as  cheerio  from "cheerio";
import { first } from "cheerio/lib/api/traversing";

const runTask = (groupId: string, artifactId: string): void => {
    let url = `https://mvn.coderead.cn/version?groupId=${groupId}&artifactId=${artifactId}`;
    fetch(url)
        .then((res) => res.text())
        .then((res) => {
            // const alfy = new Alfy()
            // const domPasrser = new DOMParser()
            // const html = parseDocument(res);
            const $ = cheerio.load(res)
            $('tr[onclick="doFold($(this))"]').each((i, elem) => {
                console.log($(elem).html())
            })
            // console.log($.html())

            // const vHtmls = html.sourceCodeLocation('tr[onclick="dofold($(this))"]')
            // vHtmls.forEach(vHtml => {
            //     const version = vHtml.querySelector('td:nth-child(1)')!.textContent || ""
            //     const downloadCount = vHtml.querySelector('td:nth-child(3)')!.textContent || ""
            //     const time = vHtml.querySelector('td:nth-child(4)')!.textContent

            //     const temp: Item = {
            //         title: version,
            //         arg: version,
            //         subtitle: `update by ${time} \t downloadCount: ${downloadCount}`
            //     }
            //     alfy.addItem(temp)
            // })
            // console.log(res);
        });
    // let url = `https://mvn.coderead.cn/search?keyword=${query}`
    // const alfy = new Alfy()

    // fetch(url)
    //     .then(res => res.json() as Promise<Mvn>)
    //     .then(res => {
    //         const reg1 = `'text' >(.*)</span>.*description'>(.*)</span>.*description'>(.*)</span>`
    //         const reg2 = /<.?em>/g
    //         res.results.forEach(result => {
    //             let arr = result.name.match(reg1)
    //             let title = arr![1].replace(reg2, '').trim()
    //             let time = arr![2].trim()
    //             let description = arr![3].replace(reg2, '').trim()

    //             const temp: Item = {
    //                 title: `${description}:${title}`,
    //                 subtitle: `update by ${time}`,
    //                 arg: `${description}:${title}`
    //             }
    //             alfy.addItem(temp)
    //         })

    //         console.log(JSON.stringify(alfy))
    //     })
};

export default runTask;
