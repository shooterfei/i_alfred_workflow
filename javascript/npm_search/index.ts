// import fetch from "node-fetch"

import { Alfy, Item } from "..";
import dateFormat from "dateformat";
// 简化的结构体
interface Npm {
  objects: Array<NObject>;
}

interface NObject {
  score: {
    detail: {
      popularity: number;
      quality: number;
      maintenance: number;
    };
  };
  package: {
    name: string;
    date: {
      ts: number;
    };
    links: {
      npm: string;
    };
  };
}

const runTask = (query: string): void => {
  const url = `https://www.npmjs.com/search?q=${query}`;
  let requestInit: RequestInit = {
    headers: {
      "x-spiferack": "1",
    },
  };

  const alfy = new Alfy();
  fetch(url, requestInit)
    .then((res) => res.json() as Promise<Npm>)
    .then((res) => {
      res.objects.forEach((item) => {
        const score = item.score.detail;
        const temp: Item = {
          title: item.package.name,
          subtitle: `update by ${dateFormat(
            item.package.date.ts,
            "yyyy-mm-dd HH:MM:ss"
          )} \t \
p: ${score.popularity.toFixed(2)} \
q: ${score.quality.toFixed(2)} \
m: ${score.maintenance.toFixed(2)}`,
          arg: item.package.links.npm,
        };
        alfy.addItem(temp);
      });
      console.log(JSON.stringify(alfy));
    });
};

export default runTask;
