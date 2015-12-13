+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "FAQ"
weight = 6
menu = "main"
toc = true
+++

# Why don't I see my repository?

Here are some tips for troubleshooting:

* Is your GitHub repository public?
* Do you have Admin access to the repository?
* Did you grant access to your GitHub organization?
* Did you recently create the repository? There may be up to a 15 minute delay due to internal caching.
* Does your organization restrict integrations? [Learn more](https://github.com/blog/1941-organization-approved-applications).

# Can I use with private repositories?

Not with the free hosted service at this time.

In order to use with private repositories you will need to install and run your
own instance, which is really really simple, we promise. See the [the documentation](../install).

# Can I use with GitHub Enterprise?

In order to use with GitHub Enterprise you will need to install and run your
own instance. See the [the documentation](../install).

# Can I install on my own server?

Yes, we have a Docker image available. See the [the documentation](../install).

# Why isn't this thing working?!

Verify that hooks are being sent correctly. You can see an audit log of
all hooks in the **Webhooks & Services** section or your repository settings.

Verify the message is being processed successful. An unsuccessful message
will be flagged accordingly in GitHub. Error messages from the service are written
to the response body.

![hook error](/docs/images/hook_error.png)

Verify the response from a successful hook. The approval settings, approval status,
and list of individuals that approved the pull request are included in the payload
for auditing purposes.

![hook success](/docs/images/hook_success.png)

If the payload indicates the pull request was approved, and this is not reflected
in the status you can click the re-deliver button to re-deliver the payload.
