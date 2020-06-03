{ sources ? import ./nix/sources.nix,
  pkgs ? import sources.nixpkgs {} }:

pkgs.buildGo113Module {
  pname = "vaportext";
  version = "1.0.2";

  modSha256 = "1q1lx8fwhr26gis0fl76h48xk04kvvrzadv3arnih9qgc31l6rjp";

  src = pkgs.lib.cleanSource ./.;

  meta = with pkgs.lib; {
    description = "make text more ａｅｓｔｈｅｔｉｃ";
    platforms = platforms.linux ++ platforms.darwin;
  };
}
