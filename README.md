# mackerel-plugin-sample

Sample plugin for mackerel.io agent.

## Synopsis

```shell
mackerel-plugin-sample [-metric-key-prefix=<prefix>]
```

## Example of mackerel-agent.conf

```
[plugin.metrics.sample]
command = "/path/to/mackerel-plugin-sample"
```

## How to release

Release by TravisCI

1. `git tag vx.y.z`
2. git push --tags
3. Wait to build at https://travis-ci.org/shibayu36/mackerel-plugin-sample
4. See https://github.com/shibayu36/mackerel-plugin-sample/releases

Release by manually

1. `git tag vx.y.z`
2. GITHUB_TOKEN=... goreleaser --rm-dist
3. See https://github.com/shibayu36/mackerel-plugin-sample/releases
