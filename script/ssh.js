#!/usr/bin/env osascript -l JavaScript

Application("com.runningwithcrayons.Alfred").runTrigger("ssh", {inWorkflow: "shooterfei", withArgument: "test"})
