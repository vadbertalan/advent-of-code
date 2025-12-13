import fs from 'fs';
import path from 'path';

/**
 * Reads a file and returns its content as a list of strings, split by newline.
 * Trims the last empty line if it exists (common behavior for input files).
 */
export function readLines(filePath: string): string[] {
    const content = fs.readFileSync(filePath, 'utf-8');
    // split by newline, handle \r\n
    const lines = content.split(/\r?\n/);
    
    // Remove last empty line if present (often mostly empty files end with newline)
    if (lines.length > 0 && lines[lines.length - 1] === '') {
        lines.pop();
    }
    
    return lines;
}

/**
 * Gets the input file extension based on some logic, defaulting to 'in'.
 * In the Go code this seemed to just default to "in" mostly, but we can make it flexible.
 */
export function getInputFileExt(day: number): string {
    return "in";
}
