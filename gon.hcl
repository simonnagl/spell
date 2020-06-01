# The path follows a pattern
# ./dist/BUILD-ID_TARGET/BINARY-NAME
source = ["./dist/spell-apple_darwin_amd64/spell"]
bundle_id = "com.github.simonnagl.spell.cmd.spell"

apple_id {
  username = "simonnagl@aim.com"
  password = "@env:AC_PASSWORD"
}
sign {
    application_identity = "Apple Development: simonnagl@aim.com"
}