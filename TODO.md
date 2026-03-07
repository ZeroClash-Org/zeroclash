# TODO

## Milestone: Minimum Viable Product

In this stage, zeroclash could do basically network data inbound process,
package forward and rule match.

- [ ] Inbound design and implementation
- [ ] Shadowsocks, hysteria, hysteria2 adaptor implementation
- [ ] Rule match implementation

## Milestone: Refactor

In this stage, the code refactor should be the main point of the entire work,
we will use modern way to make code better.

Code like below should use something like function registry factor to refactor.

```go
switch strings.ToLower(v.Type) {
case "shadowsocks":
        proxies[v.Name] = nil
case "hysteria":
        proxies[v.Name] = nil
case "hysteria2":
        proxies[v.Name] = nil
default:
        logger.Get().Warn("unknown protocol", zap.String("protocol", proxies[v.Name].Name()))
}
```

## Milestone: Performance

In this stage, the performance should be improved via something like
pre-compile, zero-copy and etc.
