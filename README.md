### Minimum Pods Kubernetes scheduler

This is a custom kubernetes scheduler that can be used for demo.
It's not intended for production usage.

The scheduler watches pods and binds them to nodes that have a minimum number of pods, then emits "Scheduled" events.
