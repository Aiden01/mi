let
  pkgs = import <nixpkgs> {};
  executable = pkgs.callPackage ./default.nix {};
  mi = pkgs.runCommand "mi" { buildInputs = [pkgs.makeWrapper]; } ''
   mkdir -p $out/bin
   cp ${executable}/bin/mi $out/bin/mi
  '';
in mi
