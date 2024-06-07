//
//  FileShareApp.swift
//  FileShare
//
//  Created by Hans Halfar on 13.05.24.
//

import MobileServer
import SwiftUI

class GoServer {
  public let server: MobileServer!
  private let port: Int

  public var frontendURL: URL {
    return URL(string: "http://localhost:\(port)/nextjs/")!
  }

  init(port: Int) {
    self.port = port
    guard let server = MobileNewNextJSHandler(":\(port)") else {
      NSLog("failed to launch go server")
      // TODO: crash here
      self.server = nil
      return
    }
    self.server = server

  }

}

class AppDelegate: NSObject, UIApplicationDelegate {
  let goServer: GoServer = GoServer(port: 3002)

  func application(
    _ application: UIApplication,
    didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]? = nil
  ) -> Bool {
    return true
  }
}

struct ContentView: View {
  @Binding var url: URL
  var body: some View {
    Text(url.absoluteString)

      .padding(10)
      .font(.caption2)
      .background(
        .regularMaterial,
        in: RoundedRectangle(cornerRadius: 8, style: .continuous)
      )
  }
}

@main
struct FileShareApp: App {
  @UIApplicationDelegateAdaptor(AppDelegate.self) var appDelegate
  @State var url: URL = URL(string: "http://localhost:3002/nextjs/")!
  @State var isLoading: Bool = true

  var body: some Scene {
    WindowGroup {
      VStack {
        ZStack(alignment: .bottomTrailing) {
          WebView(url: $url, isLoading: $isLoading)
            .ignoresSafeArea()
            .overlay {
              ProgressView().opacity(isLoading ? 1 : 0)
            }
          ContentView(url: $url)
        }
        HStack(alignment: .center) {

          Spacer()
          Button(
            action: {
              self.url = URL(string: "http://localhost:3002/nextjs/")!
              print(self.url)
            },
            label: {
              Label(
                title: { Text("Embedded") },
                icon: { /*@START_MENU_TOKEN@*/
                  Image(systemName: "42.circle") /*@END_MENU_TOKEN@*/
                }
              )
            })
          Spacer()
          Button(
            action: {
              self.url = URL(string: "http://localhost:3003/nextjs/")!
              print(self.url)
            },
            label: {
              Label(
                title: { Text("Remote (DEV)") },
                icon: { /*@START_MENU_TOKEN@*/
                  Image(systemName: "42.circle") /*@END_MENU_TOKEN@*/
                }
              )
            }
          )

          Spacer()
        }.padding()
      }
    }
  }
}
