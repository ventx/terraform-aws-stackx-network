## Examples - non-k8s

This example deploys the Terraform module with `var.k8s` set to `false`.

Now you will get evenly sized /20 CIDR network ranges for all your subnets.

```hcl
k8s = false
public = true // 10.1.48.0/20, 10.1.64.0/20, 10.1.80.0/20
private = true // 10.1.0.0/20, 10.1.16.0/20, 10.1.32.0/20
db = true // 10.1.144.0/20, 10.1.160.0/20, 10.1.176.0/20
```
