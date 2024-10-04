package proxy

type Application struct{}

func (app *Application) HandleRequest(url, method string) (int, string) {
	if url == "user" && method == "GET" {
		return 200, "OK"
	}

	if url == "user" && method == "POST" {
		return 201, "User Created"
	}

	return 404, "Not OK"
}
