{inputs}: let
  std = import ./std.nix {inherit inputs;};
  tree = import ./tree.nix {inherit inputs pkgs;};
  inherit (inputs.nixpkgs) lib;
  overlay = import ./packages {inherit inputs tree lib std;};
  systems = import ./systems {inherit inputs tree lib std;};
  shells = import ./shells {inherit inputs tree lib std pkgs;};
  inherit (import ./pkgs.nix {inherit inputs tree overlay;}) pkgs;
  formatter = import ./formatter.nix {inherit inputs pkgs;};
  inherit (std) set;
  checks = set.map (_: deployLib: deployLib.deployChecks inputs.self.deploy) inputs.deploy-rs.lib;
in
  {
    inherit inputs tree std pkgs checks formatter lib;
    legacyPackages = pkgs;
  } // systems // shells
