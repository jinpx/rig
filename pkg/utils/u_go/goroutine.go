package u_go

// Go goroutine
func Go(fn func()) {
	go safe_try(fn, nil)
}
