package piscine

func UltimateDivMod(a *int, b *int) {
	dividend := *a
	divisor := *b

	quotient := dividend / divisor
	remander := dividend % divisor

	*a = quotient
	*b = remander
}
