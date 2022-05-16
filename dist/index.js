import "@jxa/global-type";
import { run } from "@jxa/run";
let argv = process.argv;
run((arg) => {
    // console.log(arg[2])
    // console.log("------------------------")
    // console.log(`'${JSON.stringify(arg[2])}'`)
    Application("com.runningwithcrayons.Alfred")
        .runTrigger("ssh", {
        inWorkflow: "shooterfei",
        withArgument: `${arg[2]}`
    });
}, argv);
//# sourceMappingURL=index.js.map
