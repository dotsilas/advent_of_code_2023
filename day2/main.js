import { part1 } from "./Part1.js";
const fs = require("fs")

const lines = fs.readFileSync("./input.txt", "utf8")
  .split("\n")
  .filter(line => line !== "")

// answer part 1
part1(lines)


