# This topic should be separated into several guides that compose a Saltstack 'track'. 
The goal of the track will be to show how to use SaltStack to configure and manage a fleet of Linodes with WordPress deployed on them.

This will be written from the perspective of a managed WordPress hosting company using Linode as the platform for their business. The company will offer for each of their users access to a pre-installed WordPress setup along with SSH access to that user's document root.

The implementation will have a set number of users installed on the same Linode and a number of Linodes according to how many users the company services. Include a diagram that describes the final implementation.

Digital Ocean has a multi-guide track for SaltStack here:
https://www.digitalocean.com/community/tutorial_series/managing-development-environments-with-saltstack

Their guides describe using Salt to manage a load-balanced (they set up HAProxy) set of web servers. This differs from this guide, as the web servers are not load-balanced here. I'm not aware of a method for managing NodeBalancers via Salt. This guide's central conceit avoids that complication (without telling users to set up their own HAProxy instead of using NodeBalancers) and can be entirely orchestrated through Salt. Reselling Linode is also a popular use of the platform and I don't believe it's described elsewhere in the library yet.

If this guide ends up overly long, we may want to split it into a track composed of multiple guides. This could potentially replace the existing SaltStack guides in the library.

## Install SaltStack

Set up a master and minion and running a test ping between them. The Linodes will be manually created by the user via the Manager or API. This will mostly be the same as the existing Installing Salt guide.

## Salt States

- This guide will initially just create a Salt state for Apache and then apply that to the minion. The test for this will be loading the default Apache page.

- It will then create further states for installing MySQL and PHP and then apply them to the minion. The guide will note when applying them that the declarative nature of the state configuration means that the minion does not also attempt to install Apache again. Include note about how the Salt state describes that the minion *should* be in.

- Set Salt states for WordPress, similar to this: https://gist.github.com/hbokh/944987bb7f34dc38767830b882364e3e

- Talk about how a service can be set to start at boot or reload automatically on certain conditions (e.g. if a virtual host is added):
ninx:
  service.running:
    - enable: True
    - reload: True
    - watch:
      - file: /etc/nginx/sites-enabled/vhost.conf

- Make note of pushing a file to the minion from the master, as in the /var/www/wordpress/wp-config.php state:
  file.managed:
    - source: salt://vm/wp-config.php

    https://docs.saltstack.com/en/getstarted/config/files.html

## Salt Pillar

- Introduce Pillar and explain that one use case for Pillar is storing secrets

- Move WordPress/database creds/password to Pillar and then use Jinja2 to update the password parameters in the WordPress states

## Pillar, States, and Jinja2 continued

- Introduce another use case for Pillar: storing the usernames of the service

- Set up users in Pillar and move WordPress creds/other relevant info (domain name of website) for each into Pillar

- In the salt state file, update the states so that they loop through the users from Pillar via Jinja2's loop syntax and create configurations for each user.

- Create new salt state for Linux user creation on minion (to give SSH access to users)

## More Minions

- Create new minions, add users to them via Pillar, and apply state

- Set up Salt Cloud to handle creation of new minions

## Command Execution

- Update software on minions

- Create execution module or state that resets data or reinstalls software for user (in case user accidentally breaks their WordPress installation)

- Check on disk usage across fleet

## Other administration possibilities

- Set disk quotas for users

- Set chroot jails for SSH users