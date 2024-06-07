import SwiftUI
import WebKit

class WkWebViewViewController: UIViewController {

  var webView: WKWebView {
    return self.view as! WKWebView
  }
  var config: WKWebViewConfiguration

  init(config: WKWebViewConfiguration) {
    self.config = config
    super.init(nibName: nil, bundle: nil)
    self.view = WKWebView(frame: .zero, configuration: config)
  }

  required init?(coder: NSCoder) {
    fatalError("init(coder:) has not been implemented")
  }

  override func viewDidLoad() {
    super.viewDidLoad()
  }
}

extension WkWebViewViewController: WKNavigationDelegate {
}

struct WebView: UIViewControllerRepresentable {
  @Binding var url: URL
  @Binding var isLoading: Bool

  init(url: Binding<URL>, isLoading: Binding<Bool>) {

    self._url = url
    self._isLoading = isLoading

  }

  func makeUIViewController(context: Context) -> WkWebViewViewController {
    let config = WKWebViewConfiguration()
    config.ignoresViewportScaleLimits = false
    config.allowsInlineMediaPlayback = true
    // TODO: move events to the web view controller
    config.userContentController.add(context.coordinator.messageHandler, name: WebMessages.boot.rawValue)

    let vc = WkWebViewViewController(config: config)
    let view = vc.webView
    view.isInspectable = true
    view.navigationDelegate = context.coordinator
    return vc
  }

  func makeCoordinator() -> Coordinator {
    return Coordinator(self, url: self.$url)
  }

  class Coordinator: NSObject, WKNavigationDelegate {

    let messageHandler: WebKitMessageHandler
    let parent: WebView

    init(_ parent: WebView, url: Binding<URL>) {
      self.messageHandler = WebKitMessageHandler()
      self.parent = parent
    }

    func webView(_ webView: WKWebView, didCommit navigation: WKNavigation!) {
      parent.isLoading = true
      print("loading page")
    }

    func webView(_ webView: WKWebView, didFinish navigation: WKNavigation!) {
      parent.isLoading = false
      print("loaded page")
    }

    func webView(
      _ webView: WKWebView, didReceive challenge: URLAuthenticationChallenge,
      completionHandler: @escaping (URLSession.AuthChallengeDisposition, URLCredential?) -> Void
    ) {
      // Handle SSL challenges here
      handleAuthenticationChallenge(challenge: challenge, completionHandler: completionHandler)
    }

    private func handleAuthenticationChallenge(
      challenge: URLAuthenticationChallenge,
      completionHandler: @escaping (URLSession.AuthChallengeDisposition, URLCredential?) -> Void
    ) {
      guard let serverTrust = challenge.protectionSpace.serverTrust else {
        completionHandler(.performDefaultHandling, nil)
        return
      }

      let credential = URLCredential(trust: serverTrust)
      completionHandler(.useCredential, credential)
    }

  }

  func updateUIViewController(_ uiViewController: WkWebViewViewController, context: Context) {
      print("loading: \(url) was \(uiViewController.webView.url)")
      let old =  uiViewController.webView.url?.absoluteString.trimmingCharacters(in:   CharacterSet(charactersIn: "/"))
      let new =  url.absoluteString.trimmingCharacters(in:   CharacterSet(charactersIn: "/"))
      if old != new {
          uiViewController.webView.load(URLRequest(url: url))
      }

      print("user content controller", uiViewController.webView.configuration.userContentController)
    }
}
