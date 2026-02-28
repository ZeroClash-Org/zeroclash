<!-- markdownlint-disable MD041 -->
![Zeroclash](./Document/assets/zeroclash.png)

# Zeroclash - Different clash kernel

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ZeroClash-Org/zeroclash)
![GitHub contributors](https://img.shields.io/github/contributors/ZeroClash-Org/zeroclash)
![GitHub License](https://img.shields.io/github/license/ZeroClash-Org/zeroclash)
[![codecov](https://codecov.io/gh/ZeroClash-Org/zeroclash/graph/badge.svg?token=1FZD2EAONP)](https://codecov.io/gh/ZeroClash-Org/zeroclash)

## Make Contributions

Learn [Contributing to a project](https://docs.github.com/en/get-started/exploring-projects-on-github/contributing-to-a-project).

This project following the [conventional branch](https://conventional-branch.github.io) and [conventional commit](https://conventionalcommits.org).

### Pre-Commit Hooks

We recommend to use prek as the pre-commit trigger to check new code, to install it, please refer [prek installation](https://prek.j178.dev/installation/).

After install prek, install hooks via `prek install -c .pre-commit-config.yaml`.

## Configuration

```yaml
# [controlAddr] is the address for the control service to listen
controlAddr: <string>

# [proxy] is a list of proxy, refer to [ProxyCfg]
proxy: <[]ProxyCfg>

# [rule] is a list of rule, refer to [RuleCfg]
rule: <[]RuleCfg>
```

## Coverage Trace

![Coverage Trace](https://codecov.io/gh/ZeroClash-Org/zeroclash/graphs/sunburst.svg?token=1FZD2EAONP)

## License

This software is released under the MIT license.
