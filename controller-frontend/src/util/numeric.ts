export function isTotpNumber(v: string): boolean {
    let computed = v.replaceAll(" - ", "")
    return !isNaN(Number(computed))
}
