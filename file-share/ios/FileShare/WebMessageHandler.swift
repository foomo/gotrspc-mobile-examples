import WebKit

enum WebMessages: String {
  case boot = "bootCompleted"
}

class WebKitMessageHandler: NSObject, WKScriptMessageHandler {
  var webView: WKWebView?

  override init() {
    super.init()
    self.webView = nil
  }

  func userContentController(
    _ userContentController: WKUserContentController, didReceive message: WKScriptMessage
  ) {
    if message.name == WebMessages.boot.rawValue {
      guard let wv: WKWebView = self.webView else {
        return
      }
    } else {
      print("unknown message", message.name)
    }
  }

}
