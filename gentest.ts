import fs from 'fs';
import path from 'path';
import { execSync } from 'child_process';

const args = process.argv.slice(2);

if (args.length !== 6) {
    console.log("Usage: npx tsx gentest.ts <year> <day> <first_example_answer> <second_example_answer> <first_answer> <second_answer>");
    process.exit(1);
}

const [year, day, firstExampleAnswer, secondExampleAnswer, firstAnswer, secondAnswer] = args;

const dayDir = path.join(year, day.padStart(2, '0'));
if (!fs.existsSync(dayDir)) {
    console.error(`Directory ${dayDir} does not exist. Please run setup first.`);
    process.exit(1);
}

const testFilePath = path.join(dayDir, `${parseInt(day)}.test.ts`);

const templateAdjusted = `
import { first, second } from './${parseInt(day)}';
import { readLines } from '@/utils-ts/index';

describe('Day ${day}', () => {
    
    test('First part example', () => {
        const lines = readLines('${dayDir}/${parseInt(day)}.exin');
        expect(first(lines)).toBe('${firstExampleAnswer}');
    });

    test('Second part example', () => {
         const lines = readLines('${dayDir}/${parseInt(day)}.exin');
         expect(second(lines)).toBe('${secondExampleAnswer}');
    });
    
    test('First part real', () => {
        const lines = readLines('${dayDir}/${parseInt(day)}.in');
        expect(first(lines)).toBe('${firstAnswer}');
    });

    test('Second part real', () => {
        const lines = readLines('${dayDir}/${parseInt(day)}.in');
        expect(second(lines)).toBe('${secondAnswer}');
    });
});
`;

fs.writeFileSync(testFilePath, templateAdjusted);
console.log(`Generated test file: ${testFilePath}`);

try {
    console.log("Running generated test...");
    execSync(`npx jest ${testFilePath}`, { stdio: 'inherit' });
} catch (e) {
    console.error("Test execution failed.");
    process.exit(1);
}
