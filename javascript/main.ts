import mvn_search from "./mvn_search"
import npm_search from "./npm_search"

const flag = process.argv[3]
const q = process.argv[4]


switch (flag) {
    case 'mvn_search':
        mvn_search(q)
        break;
    case 'npm_search':
        npm_search(q)
    default:
        break;
}
