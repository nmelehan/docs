---
author:
  name: Linode
  email: docs@linode.com
description: 'Use the Linode Manager to check if your Linodes have any pending free upgrades and then take those upgrades.'
keywords: ["Linode Manager", "upgrade", "pending upgrades", "free upgrades"]
license: '[CC BY-ND 4.0](http://creativecommons.org/licenses/by-nd/4.0/)'
modified: 2019-01-02
modified_by:
  name: Linode
published: 2019-01-02
title: How to Upgrade your Linode to the Newest Plans
---

Background on how Linode periodically makes free upgrades across each of the hardware plan tiers, like increases to storage space and memory. Explain how your new hardware plan will have the same monthly and hourly cost as your current plan.

{{< note >}}
Linode announces new hardware upgrades on the [Linode Blog](https://blog.linode.com)

## How to Tell if your Linode is Eligible for an Upgrade

New manager: On Linodes page, each Linode will show `(pending upgrade)` in the plan description column. On a Linode's dashboard, a `pending upgrade` banner will appear.

Old manager: a `free upgrade` banner will appear on the Linode's dashboard

Include screenshots? Two separate guides for old manager and new manager?

## How to Initiate a Pending Upgrade

-    Visit the Linode's dashboard

-    Click through the upgrade banner

-    Review which hardware resources will change

-    Enter the upgrade queue:

     -   Upgrading will gracefully power your Linode down (if it is not already shut down)

     -   The migration takes approximately 3-5 minutes per GB of disk space your Linode has

     -   The Linode will be powered on at the end of the upgrade if it was running prior to the upgrade.

### Initiate a Pending Upgrade with the API

How to use the `mutate` endpoint:

https://developers.linode.com/api/v4#operation/mutateLinodeInstance