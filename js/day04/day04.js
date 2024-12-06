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

function addPadding(grid, times = 3) {
  let newGrid = grid.map((row) => [...row]);

  for (let i = 0; i < times; i++) {
    newGrid = newGrid.map((row) => {
      row.unshift('.');
      row.push('.');
      return row;
    });

    const padding = Array(newGrid[0].length).fill('.');
    newGrid.unshift(padding);
    newGrid.push(padding);
  }

  return newGrid;
}

function isKeyWord(word) {
  return word === 'XMAS' || word === 'SAMX';
}

function part1(input, padding = 3) {
  //iterate over the grid
  let finds = 0;
  for (let x = 0 + padding; x < input.length - padding; x++) {
    for (let y = 0 + padding; y < input[x].length - padding; y++) {
      //search horizontal
      let word = input[x][y] + input[x][y + 1] + input[x][y + 2] + input[x][y + 3];
      if (isKeyWord(word)) {
        finds++;
      }
      //search vertical
      word = input[x][y] + input[x + 1][y] + input[x + 2][y] + input[x + 3][y];
      if (isKeyWord(word)) {
        finds++;
      }
      //search diagonally up
      word = input[x][y] + input[x - 1][y + 1] + input[x - 2][y + 2] + input[x - 3][y + 3];
      if (isKeyWord(word)) {
        finds++;
      }
      //search diagonally down
      word = input[x][y] + input[x + 1][y + 1] + input[x + 2][y + 2] + input[x + 3][y + 3];
      if (isKeyWord(word)) {
        finds++;
      }
    }
  }
  return finds;
}

function part2(input) {}

const fileContents = readFileContents('ex.txt');

const lines = fileContents
  .trim()
  .split('\n')
  .map((line) => line.split(''));

const paddedGrid = addPadding(lines);

console.log(part1(paddedGrid));
