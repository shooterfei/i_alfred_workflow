import fetch from "node-fetch"
import { Alfy, Item } from ".."

interface Result {
    name: string,
    value: string,
    text: string
}

interface Mvn {
    success: boolean,
    results: Array<Result>
}

const runTask = (query: string): void => {
    let url = `https://mvn.coderead.cn/search?keyword=${query}`
    const alfy = new Alfy()
    fetch(url)
        .then(res => res.json() as Promise<Mvn>)
        .then(res => {
            const reg1 = `'text' >(.*)</span>.*description'>(.*)</span>.*description'>(.*)</span>`
            const reg2 = /<.?em>/g
            res.results.forEach(result => {
                let arr = result.name.match(reg1)
                let title = arr![1].replace(reg2, '').trim()
                let time = arr![2].trim()
                let description = arr![3].replace(reg2, '').trim()

                const temp: Item = {
                    title: `${description}:${title}`,
                    subtitle: `update by ${time}`,
                    arg: `${description}:${title}`
                }
                alfy.addItem(temp)
            })

            console.log(JSON.stringify(alfy))
        })
}


export default runTask
