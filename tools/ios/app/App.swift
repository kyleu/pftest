// Content managed by Project Forge, see [projectforge.md] for details.
import SwiftUI
import PftestServer

@main
struct Project: App {
    init() {
        print("starting Test Project...")
        let path = NSSearchPathForDirectoriesInDomains(.libraryDirectory, .userDomainMask, true)
        let port = PftestServer.CmdLib(path[0])
        print("Test Project started on port [\(port)]")
        let url = URL.init(string: "http://localhost:\(port)/")!
        self.cv = ContentView(url: URLRequest(url: url))
    }

    var cv: ContentView

    var body: some Scene {
        WindowGroup {
            cv
        }
    }
}
