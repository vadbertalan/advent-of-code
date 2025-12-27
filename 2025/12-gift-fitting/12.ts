// https://adventofcode.com/2025/day/12

import { readLines, getInputFileExt } from "@/utils-ts/index.js";

const aocDay = 12;

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

export function first(lines: string[]): string {
  let result = 0;

  for (const line of lines) {
    if (!line.includes("x")) {
      continue;
    }

    const parts = line.split("x");
    const w = parseInt(parts[0]);

    const rest = parts[1].split(":");
    const h = parseInt(rest[0]);

    const arrayOfCounts = rest[1]
      .trim()
      .split(" ")
      .map((x) => parseInt(x));
    const total = arrayOfCounts.reduce((a, b) => a + b, 0);

    result += total <= Math.floor(w / 3) * Math.floor(h / 3) ? 1 : 0;
  }

  return result.toString();
}

// 254 too low
// 512 correct: needed `<=` instead of `<`

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
  } catch (e) {
    console.error(
      "Error reading input file. Make sure you are in the directory of the day's puzzle.",
      e
    );
    throw e;
  }
}

!module.parent && main();
