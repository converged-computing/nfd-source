# Node Feature Discovery Source

This is [node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery) source moved out of tree to allow for using the APIs separately from Kubernetes.
To join the discussion, please see [this issue](https://github.com/kubernetes-sigs/node-feature-discovery/issues/1581).

## Notes

These are development notes for the migration. I used [git-filter-repo](https://github.com/newren/git-filter-repo/blob/main/INSTALL.md) to select just source and the specific API directory:

```bash
git-filter-repo  --path source/ --path pkg/apis/nfd/v1alpha1/ --path pkg/utils/ --path pkg/cpuid/ --path LICENSE
```

The above ensures that original author contributions (commit history) is maintained.

### Changes

I made the following changes to cleanly separate the Kubernetes logic from the source logic.

- pkg/apis/nfd/v1alpha1/types.go
  - The need for corev1 in types.go is only because of Taints, which I don't see anywhere in source so I'm removing.
  - metav1 is to add object metadata (also for Kubernetes) so I will remove that for now.
  - The protocol buffers would also be for gRPC, which is a specific use case not necessary for underlying source.
  - It looks like we don't actually need NodeFeature, NodeFeatureSpec, NodeFeatureList, NodeFeatureRule (these are higher level in the logic)
  - The entire namespace of Rule/Match is oriented to custom, which also seems like it belongs in the upstream
  - I'm keeping the same path for now to maintain git history
  - When we remove the above, the interface is very simple!
- pkg/apis/nfd/v1alpha1/register.go
  - Is primarily for K8s schemas, etc. can be removed.
- pkg/apis/nfd/v1alpha1/feature.go
  - MergeInto functions are not used, as aren't the Node specific ones. 
- pkg/apis/nfd/v1alpha1/nodefeaturerule
  - Also does not seem used in source
- pkg/utils/flags.go also not used
- pkg/utils/memory_resources.go (and test) are not used in source, this is again for K8s
- pkg/utils/(klog|kubeconf|grpc_log.go|metrics.go|tls.go) are kubernetes specific, also removing, and kubernetes.go
- klog is replaced with "slog" which is part of go as of 1.21.