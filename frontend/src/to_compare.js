import { reactive } from 'vue'

export let toCompare = []
toCompare.last = function() {
    return this.length > 0 ? this[this.length - 1] : {}
}
toCompare = reactive(toCompare)

export function addToCompare(stock) {
    toCompare.push(stock)
    if (toCompare.length > 2)
        toCompare.shift()
}