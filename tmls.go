package main

func main() {
	sessions := getSessions()
	selectedSession := selectItem(sessions)

	if selectedSession != nil {
		attachSession(selectedSession)
	}
}
