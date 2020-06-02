{ sources ? import ./nix/sources.nix,
  pkgs ? import sources.nixpkgs {} }:

pkgs.buildGo113Module {
  pname = "vaportext";
  version = "1.0.1";

  modSha256 = "1aa3d2x6yk3ik9mwlkzia0mi7cq339884zz4vhqc0fr5vil2r54n";

  src = pkgs.lib.cleanSource ./.;

  meta = with pkgs.lib; {
    description = "make text more ａｅｓｔｈｅｔｉｃ";
    platforms = platforms.linux ++ platforms.darwin;
  };
}
