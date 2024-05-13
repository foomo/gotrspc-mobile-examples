//
//  ContentView.swift
//  FileShare
//
//  Created by Hans Halfar on 13.05.24.
//

import SwiftUI

struct ContentView: View {
    @State var isShowing = false
    var body: some View {
        VStack {
            Image(systemName: "globe")
                .imageScale(.large)
                .foregroundStyle(.tint)
            Text("Hello, world!")
            Button {
                isShowing.toggle()
            } label: {
                Text("select a file")
            }
            .fileImporter(isPresented: $isShowing, allowedContentTypes: [.item], allowsMultipleSelection: true, onCompletion: { results in
                            
                            switch results {
                            case .success(let fileurls):
                                print(fileurls.count)
                                
                                for fileurl in fileurls {
                                    print(fileurl.path)
                                    server?.exposeFile(fileurl.path)
                                }
                                
                            case .failure(let error):
                                print(error)
                            }
                            
                        })
        }
        .padding()
    }
}

#Preview {
    ContentView()
}
