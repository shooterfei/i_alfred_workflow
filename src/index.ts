import "@jxa/global-type"
import { run } from "@jxa/run"


let argv = process.argv
run((arg) => {

  // const cwd = arg[1]

  switch (arg[2]) {
    case 'temp':
      Application("com.runningwithcrayons.Alfred")
        .runTrigger("ssh", {
          inWorkflow: "shooterfei",
          withArgument: `${arg[2]}`
        })
      break;
    case 'run':
      const iTerm = Application('iTerm')
      iTerm.launch()
      iTerm.activate()
      iTerm.includeStandardAdditions = true;
      const window = iTerm.currentWindow()
      console.log(window.tabs.length)
      let tab = window.currentTab()
      if (window.tabs.length < 2) {
        tab = window.createTabWithDefaultProfile()
      } else {
        tab = window.tabs[1]
      }
      // console.log(tab)
      break;
    default:
      break;
  }

}, argv)


