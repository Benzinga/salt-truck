# :truck: Salt Truck

Salt Truck is a set of utilities for Salt-based server fleets, particularly for Salt master servers.

## Commands

### SSH
This command uses Salt to query the IP address of an instance, then SSH to it. It will automatically prefer private IP addresses when possible. If the selector matches multiple minions, you will be presented with an option for which minion to SSH to.

```
truck ssh "selector*"
```
