package proxy

type IServer interface {
	HandleRequest(string, string) (int, string)
}
