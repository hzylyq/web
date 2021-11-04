package framework

type IResponse interface {
	Json(interface{}) IResponse
	JsonP(interface{}) IResponse
	Xml(interface{}) IResponse
	Html(string, interface{}) IResponse
	Text(string, ...interface{}) IResponse
	Redirect(string) IResponse
	SetHeader(string, string) IResponse
	SetCookie(string, string, int, string, string, bool, bool) IResponse
	SetStatus(int) IResponse
	SetOkStatus() IResponse
}
