export class Alfy {
    items: Array<Item>;
    constructor() {
        this.items = []
    };
    addItem(item: Item) {
        this.items.push(item)
    }
}

export interface Item {
    uid?: string,
    type?: string,
    title: string,
    subtitle: string,
    arg: string,
    autocomplete?: string,
    icon?: Icon,
    valid?: boolean,
    match?: string,
    mods?: Map<string, Mod>,
    quicklookurl?: string,
    text?: Map<string, string>
}

interface Icon {
    type: string,
    path: string
}

interface Mod {
    arg: string,
    subtitle: string,
    title?: string
}

