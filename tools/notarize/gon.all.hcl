# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_all/pftest"]
bundle_id = "com.kyleu.projectforge.pftest"

notarize {
  path = "./build/dist/pftest_0.0.0_darwin_all_desktop.dmg"
  bundle_id = "com.kyleu.projectforge.pftest"
}

apple_id {
  username = "projectforge@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/pftest_0.0.0_darwin_all.dmg"
  volume_name = "Test Project"
}

zip {
  output_path = "./build/dist/pftest_0.0.0_darwin_all_notarized.zip"
}