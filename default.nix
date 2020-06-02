{ sources ? import ./nix/sources.nix,
  pkgs ? import sources.nixpkgs {} }:

pkgs.buildGo113Module {
  pname = "vaportext";
  version = "1.0.1";

  modSha256 = "092q2swhg1815g0y31z8dq108s2vki2shv3gp6zja02icygz9fq5";

  src = pkgs.lib.cleanSource ./.;

  meta = with pkgs.lib; {
    description = "make text more ａｅｓｔｈｅｔｉｃ";
    platforms = platforms.linux ++ platforms.darwin;
  };
}
