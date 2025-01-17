{

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    # nixpkgs-stable.url = "github:nixos/nixpkgs/release-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }:

    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        goVersion = 23;
      in
      {
        overlays.default = final: prev: {
          go = final."go_1_${toString goVersion}";
        };
        devShells.default =
          with pkgs;
          mkShell {
            buildInputs = [
              go
              gotools
              golangci-lint
              gopls
              delve

              cobra-cli
            ];

            # env = { };

            # shellHook = '''';
          };
      }
    );

}
