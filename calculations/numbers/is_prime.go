package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	// TODO

	if CountDivisors(n) == 2 {
		return true
	}

	return false
}
