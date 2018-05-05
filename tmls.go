package main

func main() {
	sessions := getSessions()
	selectedSession := selectItem(sessions)

	attachSession(selectedSession)
}
