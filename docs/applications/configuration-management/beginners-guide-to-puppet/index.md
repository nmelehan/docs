---
author:
  name: Linode
  email: docs@linode.com
description: 'A look into Puppet''s primary components, features, and configurations for the new Puppet user'
keywords: ["puppet", "automation", "puppet enterprise", "configuration management"]
license: '[CC BY-ND 4.0](https://creativecommons.org/licenses/by-nd/4.0)'
modified: 2019-01-02
modified_by:
  name: Linode
published: 2018-10-16
title: A Beginner's Guide to Puppet
external_resources:
 - '[SaltStack Documentation](https://docs.saltstack.com/)'
---

EDITOR'S NOTE: This is an outline for an A Beginner's Guide to Puppet guide. It uses A Beginner's Guide to Salt as a model for the guide.

INTRODUCTION: This section should describe what Puppet is, what jobs is it used for, and what this guide will explain with regard to Puppet.

[Puppet](https://puppet.com/) is a configuration management system. Puppet uses a master/agent model in which one or more dedicated Puppet *master* servers manages one or more *nodes* running the Puppet *agent* software.

Puppet's primary job is to manage the installed software, configuration files, and other resources on your infrastructure's servers. You describe these resources using Puppet's proprietary declarative language. This process of mapping declared resources to a system configuration is generally referred to as configuration management.

This guide will introduce the core concepts that Puppet employs to fulfill this functionality.

## Masters and Nodes

Describe the relationship between the Puppet Master and your nodes. Explain how the nodes run the Puppet agent software and Masters run the server software.

Describe the two jobs of the agent software:

-   Phoning home to the master server to find out what aspects of the system need to be managed (installed software, configuration, files, etc)

-   Applying those changes to the system

### Authorization

Explain authentication between masters and nodes.

## Resources and Resource Types

Define a *resource* as any aspect of the system that will be managed by Puppet.

Show example of using the `puppet resource` command to interactively create a file on a node:

    sudo puppet resource file /tmp/test content='Hello Puppet!'
{{< output >}}
Notice: /File[/tmp/test]/ensure: defined content as '{md5}6fa6ce40c758db2a5b07537feb98f8e1'

file { '/tmp/test':
  ensure  => 'file',
  content => '{md5}6fa6ce40c758db2a5b07537feb98f8e1',
}
{{< /output >}}

Explain how the `file {}` block is an example of a resource declaration in the Puppet Language, which will be explored [further in the guide](#the-puppet-language).

    cat /tmp/test

{{< output >}}
Hello Puppet!
{{< /output >}}

Give examples of other resource types and link to resource types list in Puppet docs.

Explain how instead of using commands like `puppet resource` to manually manage resources, you instead should record your resources in files written in [the Puppet language](#the-puppet-language). Puppet then interprets these files and applies the changes they declare.

## The Puppet Language

Explain how the Puppet Language is a domain-specific language (instead of being a more general language, like YAML or Ruby). Explore benefits and trade-offs of a DSL:

https://puppet.com/blog/why-puppet-has-its-own-configuration-language

{{< note >}}
The following sections illustrate language concepts by writing code directly on the node and applying it. These demonstrations do not reflect where your code will be kept how it will be applied when actually using Puppet. Instead, in your actual usage your code will be stored on the Master and delivered to the node during an [agent run](#the-agent-run).
{{< /note >}}

### Manifests

A manifest is any Puppet code file, and it ends in the `.pp` extension.

Show example of creating a `hello.pp` manifest directly on a node and invoking it with the agent `apply` command:

{{< file "/tmp/hello.pp" >}}
notify { 'Hello Puppet!': }
{{< /file >}}

    sudo puppet apply /tmp/hello.pp

{{< output >}}
Hello Puppet!
{{< /output >}}

### Resource Declarations

Describe resource declaration syntax in more detail: titles, namevars, relationship metaparameters, etc

### Classes

Classes are named blocks of Puppet code that group together resource declarations. Classes are *defined*, and then they are *declared*.

Give example of how to define and declare a class and directly apply it on a node using the `puppet apply` command:

{{< file "cowsay.pp" >}}
class cowsay {
  package { 'cowsay':
    ensure   => present,
    provider => 'gem',
  }
}

include cowsay
{{< /file >}}

The `class cowsay {}` block defines the class, and the `include` line declares it for use.

    sudo puppet apply cowsay.pp

{{< output >}}
Notice: Compiled catalog for cowsay.puppet.vm in environment production in 0.07 seconds
Notice: /Stage[main]/Cowsay/Package[cowsay]/ensure: created
Notice: Applied catalog in 0.63 seconds
{{< /output >}}

    cowsay 'Hello Puppet!'

{{< output >}}
 _______________
| Hello Puppet! |
 ---------------
      \   ^__^
       \  (oo)\_______
          (__)\       )\/\
              ||----w |
              ||     ||
{{< /output >}}

### Variables

Explain and demonstrate variable syntax in a class. Explain behavior of undefined variables.

Show how to parameterize a class with variables. Show how to declare a class with class-like syntax so that a variable can be set.

### Defined Resource Types

Define what a Defined Resource Type is and how it differs from a class. Talk about how Puppet classes are always singletons, but DRTs can be declared more than once. Talk about why you would use a DRT (example use case: virtual host configs, creating system users). Explain the uniqueness constraint for titles and namevars between DRT declarations.

## Modules

Define what a module is: A module is a [specific directory structure](https://puppet.com/docs/puppet/5.3/modules_fundamentals.html) that organizes your Puppet code.

Explain how a module is structured:

-   A module can have an `init.pp`. This contains a class definition, and when it is present, it is considered the 'main' class for the module. The name of the class should be the same as the name of the module that contains it.

-   All other manifests in your module should define one class. If the name of your module is `my_module`, the name of your other class should be `my_module::other_class`, and the manifest that contains it should be called `other_class.pp`.

-   `init.pp` and your other class manifests should be stored inside a `manifests/` directory inside the module directory.

-   [Files](#files) are stored in a `files/` directory within the module directory.

-   [Templates](#variables-and-templates) are stored in a `files/` directory within the module directory.

### Files

Explain and show the syntax for sourcing a file in a resource (e.g. a default config for apache or something).

### Templates

Show how to use variables in templates, where to store templates in a module, how templates are included in your resources.

## The Agent Run

Explain the cycle by which an agent requests a catalog from the Master:

- The agent sends a *catalog request* to the master. A catalog is the list of resources the agent should configure.
-- Bundled with this catalog request, the agent includes a set of [facts](#facts), or intrinsic attributes, about the node system. These facts will be used by the Master to determine which resources should be applied.
- The Master compiles the catalog using the Puppet codebase you've written and the facts supplied by the node.
- The Master sends the catalog to the node, and the agent software applies the changes the catalog describes.

### site.pp

Specify where the site.pp file lives, when it is used (link to #the-agent-run)/what it does, how node classification works, including info on regex name matching and other custom classifier matching schemes for nodes.

### The modulepath

The directory that holds your modules on the Master is the *modulepath*. The modulepath can be revealed by running the `config` command:

    sudo puppet config print modulepath

{{< output >}}
/etc/puppetlabs/code/environments/production/modules:/etc/puppetlabs/code/modules:/usr/share/puppet/modules
{{< /output >}}

{{< note >}}
The modulepath is specified by the Puppet environment. [Environments](https://puppet.com/docs/puppet/4.10/environments_about.html) let you group agent nodes according to the different aspects of your development workflow, e.g. production, testing, and so on.
{{< /note >}}

Puppet will search for manifests within the directories in the modulepath in the order that they appear.

### Facts

Define what facts are, how they are gathered by the agent, how they can be accessed and used in manifests (demo the `$facts['fact_variable_name']` syntax]).

Demo conditional logic with facts (example of Apache vs HTTPD package name)

### Hiera

Explain and demo storing data with Hiera, define what kind of data you'd store in it.

Show how to store data in YAML, link to the other options for storing data.

Show using the `lookup()` function in a manifest.

## The Forge

Introduce the Forge. Explain how different tags are applied to different modules in the Forge according to the level of approval or support they have received.

### The Puppetfile

Introduce how to include modules from the Forge in the Puppetfile.

## Puppet Enterprise

Explore Differences between Puppet open source and Puppet Enterprise.