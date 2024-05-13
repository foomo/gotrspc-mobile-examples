//
//  FileShareApp.swift
//  FileShare
//
//  Created by Hans Halfar on 13.05.24.
//

import SwiftUI
import MobileServer

public let server = MobileNewNextJSHandler("localhost:3000")

@main
struct FileShareApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
    init() {
        
    }
}
