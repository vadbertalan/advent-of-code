// https://adventofcode.com/2020/day/19

import { readLines, getInputFileExt } from "@/utils-ts/index.js";
import { exit } from "process";

const aocDay = 19;

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

class LeafNode {
  constructor(public aOrB: "a" | "b") { }
}

class RuleSet {
  constructor(public refIndices: number[]) { }
}

class CombinedRuleSet {
  constructor(public refIndices1: number[], public refIndices2: number[]) { }
}

const ROOT_INDEX = 0;

type RuleMap = Record<number, LeafNode | RuleSet | CombinedRuleSet>

type MatchesResult = {
  matches: false
} | {
  matches: true;
  strOffset: number;
}

function matchesRuleSet(ruleMap: RuleMap, rule: RuleSet, str: string, initialStrOffset: number): MatchesResult {
  let localOffset = 0
  for (let index = 0; index < rule.refIndices.length; index++) {
    const subrule = rule.refIndices[index];
    const matchesResult = matches(ruleMap, subrule, str, initialStrOffset + localOffset)
    if (!matchesResult.matches) {
      return { matches: false }
    }
    localOffset += matchesResult.strOffset
  }
  return {
    matches: true, strOffset: localOffset
  }
}

function matches(ruleMap: RuleMap, ruleIndex: number, str: string, initialStrOffset: number): MatchesResult {
  const rule = ruleMap[ruleIndex]
  if (rule instanceof LeafNode) {
    const matches = str[0 + initialStrOffset] === rule.aOrB

    if (matches) {
      return {
        matches: true, strOffset: 1
      }
    }
    return { matches: false }
  }

  if (rule instanceof RuleSet) {
    const matchResult = matchesRuleSet(ruleMap, rule, str, initialStrOffset);
    return matchResult;
  }

  if (rule instanceof CombinedRuleSet) {
    const result1 = matchesRuleSet(ruleMap, new RuleSet(rule.refIndices1), str, initialStrOffset);


    if (result1.matches) {
      return result1
    }

    const result2 = matchesRuleSet(ruleMap, new RuleSet(rule.refIndices2), str, initialStrOffset);
    return result2
  }

  console.error("Wrong rule", rule)
  exit(1)
}

export function first(lines: string[]): string {
  const ruleMap: RuleMap = {}
  const inputStrs: string[] = []
  let switched = false;
  for (let index = 0; index < lines.length; index++) {
    const line = lines[index];

    if (line === "") {
      switched = true
      continue
    }

    if (!switched) {
      const left = Number(line.split(": ")[0])
      const right = line.split(": ")[1]

      if (right.includes("|")) {
        const rules = right.split(" | ");
        const leftRules = rules[0].split(" ").map(Number);
        const rightRules = rules[1].split(" ").map(Number);
        ruleMap[left] = new CombinedRuleSet(leftRules, rightRules);
      } else if (right.includes("\"")) {
        const aOrB = right.includes("a") ? "a" : "b"
        ruleMap[left] = new LeafNode(aOrB)
      } else {
        const nums = right.split(" ").map(Number)
        ruleMap[left] = new RuleSet(nums)
      }

      continue
    }

    inputStrs.push(line);
  }

  //// logic -----

  let result = 0;

  for (let index = 0; index < inputStrs.length; index++) {
    const str = inputStrs[index];
    const matchResult = matches(ruleMap, ROOT_INDEX, str, 0)
    if (matchResult.matches && matchResult.strOffset === str.length) {
      result++
    }
  }

  return result.toString();
}

// Your puzzle answer was 230.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

// function matchesRuleSet2(ruleMap: RuleMap, rule: RuleSet, str: string, initialStrOffset: number): MatchesResult {
//   let localOffset = 0
//   for (let index = 0; index < rule.refIndices.length; index++) {
//     const subrule = rule.refIndices[index];
//     const matchesResult = matches2(ruleMap, subrule, str, initialStrOffset + localOffset)
//     if (!matchesResult.matches) {
//       return { matches: false }
//     }
//     localOffset += matchesResult.strOffset
//   }
//   return {
//     matches: true, strOffset: localOffset
//   }
// }

// function matches2(ruleMap: RuleMap, ruleIndex: number, str: string, initialStrOffset: number): MatchesResult {
//   console.log(' * matches:', ruleIndex, str, initialStrOffset)
//   const rule = ruleMap[ruleIndex]

//   if (rule instanceof LeafNode) {
//     const matches = str[0 + initialStrOffset] === rule.aOrB
//     console.log(".  checking for out of bound", initialStrOffset, initialStrOffset >= str.length, str[initialStrOffset])
//     console.log(ruleIndex, 'leaf node', matches, str, rule.aOrB, initialStrOffset)

//     if (matches) {
//       return {
//         matches: true, strOffset: 1
//       }
//     }
//     return { matches: false }
//   }

//   if (rule instanceof RuleSet) {
//     console.log(ruleIndex, "ruleset", rule, str, initialStrOffset)
//     const matchResult = matchesRuleSet2(ruleMap, rule, str, initialStrOffset);
//     console.log(ruleIndex, "ruleset result", matchResult, str)
//     return matchResult;
//   }

//   if (rule instanceof CombinedRuleSet) {
//     console.log(ruleIndex, "comb1", rule.refIndices1, str)
//     const result1 = matchesRuleSet2(ruleMap, new RuleSet(rule.refIndices1), str, initialStrOffset);
//     console.log(ruleIndex, "comb1 result", rule.refIndices1, result1, str)

//     if (result1.matches) {
//       // if (ruleIndex === 8 || ruleIndex === 11) {
//       //   console.log(ruleIndex, "comb2", rule.refIndices2, str)
//       //   const result2 = matchesRuleSet2(ruleMap, new RuleSet(rule.refIndices2), str, initialStrOffset);
//       //   console.log(ruleIndex, "comb2 result", rule.refIndices2, result2, str)
//       //   return result2
//       // }

//       return result1
//     }

//     console.log(ruleIndex, "comb2", rule.refIndices2, str)
//     const result2 = matchesRuleSet2(ruleMap, new RuleSet(rule.refIndices2), str, initialStrOffset);
//     console.log(ruleIndex, "comb2 result", rule.refIndices2, result2, str)
//     return result2
//   }

//   console.error("Wrong rule", rule)
//   exit(1)
// }

export function second(lines: string[]): string {
  // const ruleMap: RuleMap = {}
  // const inputStrs: string[] = []
  // let switched = false;
  // for (let index = 0; index < lines.length; index++) {
  //   const line = lines[index];

  //   if (line === "") {
  //     switched = true
  //     continue
  //   }

  //   if (!switched) {
  //     const left = Number(line.split(": ")[0])
  //     const right = left === 8 ? "42 | 42 8" : left === 11 ? "42 31 | 42 11 31" : line.split(": ")[1]

  //     if (right.includes("|")) {
  //       const rules = right.split(" | ");
  //       const leftRules = rules[0].split(" ").map(Number);
  //       const rightRules = rules[1].split(" ").map(Number);
  //       ruleMap[left] = new CombinedRuleSet(leftRules, rightRules);
  //     } else if (right.includes("\"")) {
  //       const aOrB = right.includes("a") ? "a" : "b"
  //       ruleMap[left] = new LeafNode(aOrB)
  //     } else {
  //       const nums = right.split(" ").map(Number)
  //       ruleMap[left] = new RuleSet(nums)
  //     }

  //     continue
  //   }

  //   inputStrs.push(line);
  // }

  // //// logic -----
  // console.log(ruleMap)

  let result = 0;

  // for (let index = 0; index < inputStrs.length; index++) {
  //   const str = inputStrs[index];
  //   const matchResult = matches2(ruleMap, ROOT_INDEX, str, 0)
  //   if (matchResult.matches && matchResult.strOffset === str.length) {
  //     result++
  //   }
  // }

  return result.toString();
}

// 246 too low
// 247 too low
// 251 too low

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
