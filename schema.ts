// https://adventofcode.com/yyyy/day/dd

import { readLines, getInputFileExt } from "@/utils-ts/index.js";

const aocDay = 999;

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

export function first(lines: string[]): string {
  let result = 0;

  return result.toString();
}

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

export function second(lines: string[]): string {
  return "";
}

function main() {
  const useRealInput = process.argv.includes("-r");
  const inputFileExtension = getInputFileExt(useRealInput);

  try {
    const lines = readLines(`${__dirname}/${aocDay}.${inputFileExtension}`);

    console.log("--- First ---");
    const start1 = performance.now();
    console.log(first(lines));
    const end1 = performance.now();
    console.log(`\n✨ Done in ${((end1 - start1) / 1000).toFixed(2)} s`);

    console.log("\n--- Second ---");
    const start2 = performance.now();
    console.log(second(lines));
    const end2 = performance.now();
    console.log(`\n✨ Done in ${((end2 - start2) / 1000).toFixed(2)} s`);
  } catch (e) {
    console.error(
      "Error reading input file. Make sure you are in the directory of the day's puzzle.",
      e
    );
    throw e;
  }
}

!module.parent && main();
