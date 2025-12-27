// https://adventofcode.com/2025/day/10

import { readLines, getInputFileExt } from "@/utils-ts/index.js";
const { init } = require("z3-solver");

const aocDay = 10;

function findIndex(s: string, ch: string): number {
  return s.indexOf(ch);
}

function findLastIndex(s: string, ch: string): number {
  return s.lastIndexOf(ch);
}

function parseButtons(s: string): number[][] {
  const ret: number[][] = [];
  const buttonStrs = s.match(/\([^)]+\)/g) || [];
  for (const buttonStr of buttonStrs) {
    const nums = buttonStr
      .slice(1, -1)
      .split(",")
      .map((x) => parseInt(x));
    ret.push(nums);
  }
  return ret;
}

// First part is solved in `2025/10/10_1.go`

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

export async function second(lines: string[]): Promise<string> {
  const { Context } = await init();
  // @ts-ignore
  const ctx = new Context("main");

  let result = 0;

  for (const line of lines) {
    // Parse lights: [.##.]
    const lightsEnd = findIndex(line, "]");

    // Parse buttons: (3) (1,3) (2) (2,3) (0,2) (0,1)
    const buttonsStart = lightsEnd + 2;
    const buttonsEnd = findLastIndex(line, ")");
    const buttonsStr = line.substring(buttonsStart, buttonsEnd + 1);
    const buttons = parseButtons(buttonsStr);

    // Parse joltages: {3,5,4,7}
    const joltagesStart = buttonsEnd + 2;
    const joltagesEnd = findLastIndex(line, "}");
    const joltagesStr = line.substring(joltagesStart + 1, joltagesEnd);
    const joltages = joltagesStr.split(",").map((x) => parseInt(x));

    // const s = new Optimize();
    const opt = new ctx.Optimize();

    const zero = ctx.Int.val(0);

    const vars: any[] = [];

    for (const btn of buttons) {
      const btnVar = ctx.Int.const(`btn_${btn}`);
      opt.add(btnVar.ge(zero));
      vars.push(btnVar);
    }

    // For every joltage, sum the buttons that contribute to it
    for (let i = 0; i < joltages.length; i++) {
      const joltage = joltages[i];
      const sumExprs: any[] = [];

      for (let j = 0; j < buttons.length; j++) {
        if (buttons[j].includes(i)) {
          sumExprs.push(vars[j]);
        }
      }

      if (sumExprs.length > 0) {
        opt.add(ctx.Int.val(joltage).eq(sumExprs.reduce((a, b) => a.add(b))));
      } else {
        // If no buttons contribute but joltage is non-zero, unsatisfiable
        if (joltage !== 0) {
          opt.add(ctx.Bool.val(false));
        }
      }
    }

    const sum = vars.reduce((acc, v) => acc.add(v), ctx.Int.val(0));

    opt.minimize(sum);

    await opt.check();
    const model = opt.model();

    let modelSum = 0;
    for (const v of vars) {
      modelSum += Number(model.eval(v).toString());
    }
    result += modelSum;
  }

  return result.toString();
}

// Your puzzle answer was 16463

function main() {
  const useRealInput = process.argv.includes("-r");
  const inputFileExtension = getInputFileExt(useRealInput);

  try {
    const lines = readLines(`${__dirname}/${aocDay}.${inputFileExtension}`);

    console.log("\n--- Second ---");
    const startTime = performance.now();
    second(lines).then((result) => {
      const endTime = performance.now();
      console.log(result);
      console.log(
        `\n✨ Done in ${((endTime - startTime) / 1000).toFixed(2)} s`
      );
    });
  } catch (e) {
    console.error(
      "Error reading input file. Make sure you are in the directory of the day's puzzle.",
      e
    );
    throw e;
  }
}

!module.parent && main();
