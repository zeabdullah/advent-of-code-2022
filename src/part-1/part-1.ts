import fs from 'node:fs/promises'

const elvesCaloriesArraySeparatedByZeroes = (
    await getElvesCaloriesFromFile('src/elves-calories.txt')
)
    .split('\n')
    .map(str => Number(str))

const elves: number[][] = []
let currentZero = 0
elvesCaloriesArraySeparatedByZeroes.forEach((num, index) => {
    if (num === 0) {
        elves.push(
            elvesCaloriesArraySeparatedByZeroes
                .slice(currentZero, index)
                .filter(i => i !== 0) as number[],
        )
        currentZero = index
    }
})

let maxCalories = 0
elves.forEach(elf => {
    const currentElfCalories = elf.reduce(
        (acc, currentCal) => currentCal + acc,
        0,
    )
    if (currentElfCalories > maxCalories) maxCalories = currentElfCalories
}, 0)
console.log(maxCalories)

async function getElvesCaloriesFromFile(filename: string) {
    return fs.readFile(filename, 'utf-8')
}
