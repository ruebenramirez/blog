{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.hugo
  ];

  shellHook = ''
    echo "Hugo development environment loaded"
    echo "Hugo version: $(hugo version)"
  '';
}
