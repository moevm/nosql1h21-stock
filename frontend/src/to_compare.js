import { reactive } from 'vue'

export let toCompare = reactive([])

export function addToCompare(stock) {
    toCompare.unshift(stock)
    while (toCompare.length > 2)
        toCompare.pop()
}