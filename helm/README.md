# echoserver Helm Chart

This Helm Chart will create the echoserver Pod to create EBS tags on PVC

## Installation Commands

You can install the chart directly

```
helm upgrade --install echoserver ./echoserver/ --namespace=default --debug --dry-run
helm upgrade --install echoserver ./echoserver/ --namespace=default --debug 
```

If Helm-Chart is in a repo

```
helm upgrade --install echoserver chartmuseum/echoserver/ --namespace=default --debug --dry-run
helm upgrade --install echoserver chartmuseum/echoserver/ --namespace=default --debug 
```
