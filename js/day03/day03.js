const fs = require('fs');
const path = require('path');

function readFileContents(filePath) {
  try {
    const absolutePath = path.resolve(__dirname, filePath);
    const data = fs.readFileSync(absolutePath, 'utf8');
    return data;
  } catch (err) {
    console.error(`Error reading file from disk: ${err}`);
  }
}

function _getMulMatches(input) {
  // Regular expression for the pattern mul(x, y)
  const regex = /mul\((\d+),(\d+)\)/g;

  // Use matchAll to find all occurrences
  const matches = [...input.matchAll(regex)];

  // Map the matches to extract the components as arrays of numbers
  // Map matches to include match details and their index
  return matches.map((match) => ({
    x: parseInt(match[1], 10), // The first number
    y: parseInt(match[2], 10), // The second number
    index: match.index, // The starting index of the match in the input
  }));
}

function part1(input) {
  // Map the matches to extract the components as arrays of numbers
  const numbers = _getMulMatches(input);

  // Multiply the numbers and sum the results
  return numbers.flatMap((number) => number.x * number.y).reduce((acc, curr) => acc + curr, 0);
}

function part2(input) {
  const numbers = _getMulMatches(input);

  // Regular expression to match "do()" and "don't()"
  const regexFunctions = /do\(\)|don't\(\)/g;

  // Use matchAll to find all matches
  const matchesFunctions = [...input.matchAll(regexFunctions)];
  // Map matches to include the match, index, and whether it's a "do()"
  const dos = matchesFunctions.map((match) => ({
    index: match.index, // Starting index of the match
    isDo: match[0] === 'do()', // True if the match is "do()", false otherwise
  }));

  let result = 0;
  // go through numbers, use the index to find a do() or don't() that is before the number
  numbers.map((number) => {
    const relevantFunction = dos.filter((func) => func.index < number.index).slice(-1)[0]; // Get the last one that appears before
    if (relevantFunction === undefined || relevantFunction.isDo) {
      result += number.x * number.y;
    }
  });
  // Multiply the numbers and sum the results
  return result;
}

// Usage
const fileContents = readFileContents('data.txt');

console.log(part1(fileContents));

console.log(part2(fileContents));
