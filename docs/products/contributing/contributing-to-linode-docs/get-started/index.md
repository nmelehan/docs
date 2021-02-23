---
title: Get Started
description: "Get started contributing to the Linode Docs."
tab_group_main:
    weight: 20
---

{{< content "README" >}}

1. Fork and clone this repository.

1. Download and install Hugo version v0.80 or newer. Installation instructions for different operating systems are available in the Hugo documentation library.

1. In your terminal, navigate into the cloned docs repository.

1. Use the Node Version Manager (NVM) to install and use version 13.14.0 of Node:

        nvm install 13.14.0 nvm use 13.14.0

1. Install the Node dependencies:

        npm install

1. Start the local Hugo web server:

        hugo server

1. In a web browser, navigate to localhost:1313/docs/.