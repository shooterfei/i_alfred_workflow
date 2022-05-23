import "@jxa/global-type"
import { run } from "@jxa/run"
import path from "path"
import { fileURLToPath } from "url"

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

let argv = process.argv
console.log("dir====", process.chdir("../"))

let config = {
  mode: argv[2],
  pwd: __dirname,
  host: process.env.host,
  port: process.env.port,
  username: process.env.username,
  password: process.env.password,
  rootdir: process.env.rootdir,
  timeout: process.env.timeout
}
run((arg) => {
  switch (arg.mode) {
    case 'temp':
      Application("com.runningwithcrayons.Alfred")
        .runTrigger("ssh", {
          inWorkflow: "shooterfei",
          withArgument: `${arg.mode}`
        })
      break;
    case 'run':
      const iTerm = Application('iTerm')
      iTerm.launch()
      iTerm.activate()
      iTerm.includeStandardAdditions = true;
      
      let script = `
        export host=${arg.host}
        export port=${arg.port}
        export username=${arg.username}
        export password=${arg.password}
        export rootdir=${arg.rootdir}
        export timeout=${arg.timeout}
        '${arg.pwd}/script/template_default.exp'
      `
      const window = iTerm.currentWindow()
      console.log("config", JSON.stringify(arg))
      let tab = window.currentTab()
      if (window.tabs.length < 2) {
        tab = window.createTabWithDefaultProfile()
      } else {
        tab = window.tabs[1]
      }
      tab.currentSession.write({text: script})
      // console.log(tab)
      break;
    default:
      break;
  }

}, config)


