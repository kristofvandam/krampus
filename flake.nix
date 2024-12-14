{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    parts.url = "github:hercules-ci/flake-parts";
    devshell.url = "github:numtide/devshell";
  };

  outputs = { self, parts, ... } @ inputs:
  parts.lib.mkFlake { inherit inputs; }
  {

    imports = [
      inputs.devshell.flakeModule
    ];

    systems = [ "x86_64-linux" ];

    perSystem = { config, pkgs, system, lib, builtins, ... }: {
      devshells.default = {
        packages = with pkgs; [ go buildah passt go-swag kubernetes-helm sqlite sqlitebrowser ];
        motd = ''
          {202}🔨 Project $(basename $PRJ_ROOT){reset}
          $(type -p menu &>/dev/null && menu)
        '';
        commands = [
        {
            name = "watch-dev-build"; help = "Watch for changes and rebuild";
            command = "ag -g 'go$' --ignore docs $PRJ_ROOT/src |entr -r sh -c \"cd $PRJ_ROOT/src; DEBUG=true DB_PATH=/tmp/krampus.db go run *.go\"";
        }
        {
            name = "watch-dev-build-with-swag"; help = "Watch for changes and rebuild";
            command = "ag -g 'go$' --ignore docs $PRJ_ROOT/src |entr -r sh -c \"cd $PRJ_ROOT/src; swag init --parseDependency --parseInternal; DEBUG=true DB_PATH=/tmp/krampus.db go run *.go\"";
        }
        ];

        env = [
        {
            name = "GOROOT";
            eval = "${pkgs.go}/share/go";
        }
        {
            name = "GOPATH";
            eval = "\${PRJ_ROOT}/.go";
        }
        ];
      };
    };
  };
}
