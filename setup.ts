import fs from 'fs';
import path from 'path';
import axios from 'axios';
import 'dotenv/config';

// Parse arguments
const args = process.argv.slice(2);
const help = args.includes('--help') || args.includes('-h');
if (help) {
    console.log(`
Usage: npx tsx setup.ts [-y YEAR] [-d DAY]

Options:
  -y YEAR    Year to generate (default: current year)
  -d DAY     Day to generate (default: current day)
`);
    process.exit(0);
}

const now = new Date();
// Default to now
let year = now.getFullYear();
let day = now.getDate();

const yIndex = args.indexOf('-y');
if (yIndex !== -1 && args[yIndex + 1]) year = parseInt(args[yIndex + 1]);

const dIndex = args.indexOf('-d');
if (dIndex !== -1 && args[dIndex + 1]) day = parseInt(args[dIndex + 1]);

// Check December rule if relying on defaults
const isDefault = yIndex === -1 && dIndex === -1;
if (isDefault && now.getMonth() !== 11) { // 11 is December
    console.error("Error: Falling back to current year and day is only allowed in December.");
    process.exit(1);
}

// Ensure 2 digits for day
const dayStr = day.toString().padStart(2, '0');
const yearStr = year.toString();

const session = process.env.AOC_SESSION;

console.log(`Hello, AoC warrior! ☀️\nSetting up TypeScript workspace for day ${day}, ${year}. GL! 🤙\n`);

const dayDir = path.join(yearStr, dayStr);

// Create directories
fs.mkdirSync(dayDir, { recursive: true });
console.log(`Created folder for problem ${dayStr}, make sure to navigate into it:\n\ncd ${dayDir}\n`);

const tsFileName = path.join(dayDir, `${day}.ts`);
const exInFileName = path.join(dayDir, `${day}.exin`);
const inFileName = path.join(dayDir, `${day}.in`);

// Generate TS file from Schema
const schemaPath = 'schema.ts';
let schemaContent = fs.readFileSync(schemaPath, 'utf-8');

// Replace standard placeholders
schemaContent = schemaContent.replace(/aocDay = 999/g, `aocDay = ${day}`);
schemaContent = schemaContent.replace(/yyyy\/day\/dd/g, `${year}/day/${day}`);
// In TypeScript schema, we import from '../../utils-ts/index.js'
// If we are deeper or shallower, that might break.
// Structure is YYYY/DD/day.ts -> 2 levels deep. so ../../utils-ts is correct.

fs.writeFileSync(tsFileName, schemaContent);

// Create empty example file
fs.writeFileSync(exInFileName, "\n");

// Fetch Input
async function fetchInput() {
    if (!session) {
        console.warn("AOC_SESSION env var not found. Creating empty input file.");
        fs.writeFileSync(inFileName, "");
        return;
    }
    
    try {
        const url = `https://adventofcode.com/${year}/day/${day}/input`;
        const response = await axios.get(url, {
            headers: {
                Cookie: `session=${session}`,
                'User-Agent': 'github.com/vadbertalan/adventofcode via axios'
            },
            responseType: 'text'
        });
        
        fs.writeFileSync(inFileName, response.data);
    } catch (error: any) {
        console.error(`Error fetching input: ${error.message}`);
        // Create empty on fail
        if (!fs.existsSync(inFileName)) {
             fs.writeFileSync(inFileName, "");
        }
    }
}

fetchInput().then(() => {
    console.log(`Created files:\n- ${tsFileName}\n- ${exInFileName}\n- ${inFileName}\n`);
});
