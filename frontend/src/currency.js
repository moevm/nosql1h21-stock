export function currency(amount, name) {
    let factors = ["m", "b", "t"]
    let factor = ""
    for (let i = 0; i < factors.length; i++) {
        if (Math.abs(amount) < 10000)
            break
        amount = Math.round(amount)
        amount /= 1000
        factor = factors[i]
    }
    return amount + factor + " " + name
}