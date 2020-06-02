{ sources ? import ./nix/sources.nix,
  pkgs ? import sources.nixpkgs {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go_1_13
    pkgs.goimports
    pkgs.mysql-client
    pkgs.cloud-sql-proxy
    pkgs.golangci-lint
  ];
}
