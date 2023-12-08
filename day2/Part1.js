function part1(lines) {
  const cubesInBag = {"red": 12, "green": 13, "blue": 14}

  const games = {}

  lines.forEach(line => {
    const [id, content] = line.split(":");
    const sets = content.split(";");

    let setsMap = []

    sets.forEach(s => {
      const f = s.split(",");

      let setMap = {}

      f.forEach(e => {
        if (e.includes("red")) {
          const number = e.trim().split(" ")[1]
          setMap["red"] = number
        } else if (e.includes("green")) {
          const number = e.trim().split(" ")[1]
          setMap["green"] = number
        } else if (e.includes("green")) {
          const number = e.trim().split(" ")[1]
          setMap["green"] = number
        }})
      }

    games[id] = setsMap
  })

  console.log(games)
}

export { part1 }
