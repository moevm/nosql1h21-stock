export let toCompare = []

export function addToCompare(stock) {
    toCompare.push(stock)
    toCompare = toCompare.slice(-2)
}