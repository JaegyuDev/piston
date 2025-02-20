# TODO

## Add modded jar support
- [ ] Fabric
- [ ] (Neo)Forge
- [ ] Spigot
- [ ] Paper
- [ ] ... (probably won't do any forks past paper.)

## Stretch Goals
- [ ] Managing/Migrating JREs
- [ ] Modpack/Mod Support

## High Priority / Bugs
- [ ] Currently, there's an assumption that `--allow-snapshots` will always return a snapshot even if the latest version is a release. This leads to some loaders returning a snapshot from before the current release with this flag set.