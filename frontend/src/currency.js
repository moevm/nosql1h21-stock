export function currency(amount, name) {
    let factors = ["K", "M", "B", "T"]
    let factor = ""
    for (let i = 0; i < factors.length; i++) {
        if (Math.abs(amount) < 1000)
            break
        amount /= 100
        amount = Math.round(amount)
        amount /= 10
        factor = factors[i]
    }
    return amount + factor + " " + name
}