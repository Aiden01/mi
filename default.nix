{ buildGoPackage }:
buildGoPackage {
  name = "mi";
  goPackagePath = "github.com/eagle453/mi";
  src = ./.;
  goDeps = ./deps.nix;
}
