---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
---

You can use extensions to override and customize Drone's default behavior. If Drone does _almost_ everything you want, you can use extensions to enhance Drone to accommodate _all_ of your team's unique requirements.

Extensions are simple HTTP services that receive and respond to requests using a REST protocol. You can write extensions in any language, although we do provide starter projects for the Go programming language.

Here is a list of extension types:

{{< link "/extensions/admission.md" "Admission Extensions" >}}
{{< link "/extensions/conversion.md" "Conversion Extensions" >}}
{{< link "/extensions/conversion.md" "Configuration Extensions" >}}
{{< link "/extensions/environment" "Environment Extensions" >}}
{{< link "/extensions/registry" "Registry Extensions" >}}
{{< link "/extensions/secret.md" "Secret Extensions" >}}
{{< link "/extensions/validation.md" "Validation Extensions" >}}