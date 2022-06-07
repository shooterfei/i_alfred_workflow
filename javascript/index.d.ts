export class Alfy {
    items: Array<Item>;
    addItem(item: Item);
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
    mods?: Mods,
    quicklookurl?: string,
    text?: Map<string, string>
}

interface Icon {
    type: string,
    path: string
}

interface Mods {
    [key: string]: Mod
}
interface Mod {
    arg: string,
    subtitle: string,
    title?: string
}

