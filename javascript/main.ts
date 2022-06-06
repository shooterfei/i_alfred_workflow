import mvn_search from "./mvn_search"
import npm_search from "./npm_search"
import versions from "./mvn_package_version"

const flag = process.argv[3]
const q = process.argv[4]


switch (flag) {
    case 'mvn_search':
        mvn_search(q)
        break;
    case 'npm_search':
        npm_search(q)
    case 'versions':
       const groupId = process.argv[5] 
       const artifactId = process.argv[7] 
        versions(groupId, artifactId)
    default:
        break;
}
