# Salt Truck :truck: [![Build Status](https://travis-ci.org/Benzinga/salt-truck.svg?branch=master)](https://travis-ci.org/Benzinga/salt-truck)

Salt Truck is a set of utilities for Salt-based server fleets, particularly for Salt master servers.

You can download it as a single, self-contained binary for your platform [here](https://github.com/Benzinga/salt-truck/releases).

## Commands

### SSH
This command uses Salt to query the IP address of an instance, then SSH to it. It will automatically prefer private IP addresses when possible. If the selector matches multiple minions, you will be presented with an option for which minion to SSH to.

```
truck ssh "selector*"
```
