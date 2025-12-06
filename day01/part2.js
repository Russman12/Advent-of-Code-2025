const fs = require("fs");

const data = fs.readFileSync("./day1/input.txt").toString();

let zeroCnt = 0;
let currentNum = 50;

data
  .split("\n")
  .filter((s) => s != "")
  .forEach(actualLogic);

function actualLogic(s) {
  const val = parseInt(s.substring(1));

  for (i = 0; i < val; i++) {
    s.startsWith("L") ? currentNum-- : currentNum++;

    if (currentNum % 100 === 0) zeroCnt++;
  }
}

console.log(`count: ${zeroCnt}`);
