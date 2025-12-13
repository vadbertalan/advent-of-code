// https://adventofcode.com/yyyy/day/dd

import { readLines, getInputFileExt } from '@/utils-ts/index.js';

const aocDay = 999;

//  ____            _     _
// |  _ \\ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \\__,_|_|   \\__| |_|

export function first(lines: string[]): string {
    let result = 0;
    

    
    return result.toString();
}

//  ____            _     ____
// |  _ \\ __ _ _ __| |_  |___ \\
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \\__,_|_|   \\__| |_____|

export function second(lines: string[]): string {
    return ""
}

function main() {
    const inputFileExtension = getInputFileExt(aocDay);
    
    try {
        const lines = readLines(`${aocDay}.${inputFileExtension}`);
        
        console.log("--- First ---");
        console.log(first(lines));
        
        console.log("\n--- Second ---");
        console.log(second(lines));
    } catch (e) {
        console.error("Error reading input file. Make sure you are in the directory of the day's puzzle.", e);
    }
}

main();
