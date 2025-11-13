package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	// TODO

	count := 1

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count
}
