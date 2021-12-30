---
date: 2000-01-01T00:00:00+00:00
title: What are drone cards
author: eoinmcafee
weight: 130
header: true
aliases:
- /user-guide/cards
---

Introducing drone cards. This is a powerful feature that allows the user to display any kind of plugin output during the drone execution stage.

*Requirements:*

- *Drone Server - v2.6.0*
- *Drone Runner Docker - v1.8.0 / Drone Kubernetes Runner - latest*

For demo purpose I will be showing how we use drone cards to display output during the execution of the docker-plugin.

To get started you first want to design the card template. The easiest way to do this is through the designer:
- https://adaptivecards.io/designer/

The template describes how the card should look & what data should be displayed on the card. When the template has been created you can publish it to the following repo:
- https://github.com/harness/card-templates

Example card template:
```
{
    "type": "AdaptiveCard",
    "body": [
        {
            "type": "Image",
            "url": "https://raw.githubusercontent.com/zricethezav/gifs/master/gitleakslogo.png"
        },
        {
            "type": "ColumnSet",
            "columns": [
                {
                    "type": "Column",
                    "items": [
                        {
                            "type": "FactSet",
                            "facts": [
                                {
                                    "title": "Line:",
                                    "value": "${formatNumber(lineNumber, 0)}"
                                },
                                {
                                    "title": "Commit:",
                                    "value": "${commit}"
                                }
                            ],
                            "separator": true,
                            "$data": "${issues}"
                        }
                    ],
                    "width": "auto",
                    "verticalContentAlignment": "Center",
                    "style": "attention",
                    "bleed": true
                }
            ],
            "spacing": "Small",
            "style": "attention"
        }
    ],
    "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
    "version": "1.5"
}
```
Example card data:
```
{
	"Issues": [{
		"line": "*****",
		"lineNumber": 30,
		"offender": "*****",
		"offenderEntropy": -1,
		"commit": "e64f915755c8b2e6d50bc0c01efd34ad0920cb61",
		"repo": ".",
		"repoURL": "",
		"leakURL": "",
		"rule": "AWS Access Key",
		"commitMessage": "add leak\n",
		"author": "name",
		"email": "email",
		"file": ".drone.yml",
		"date": "2021-11-17T13:25:48Z",
		"tags": "key, AWS"
	}]
}
```

Full details on adaptive cards can be found here: https://docs.microsoft.com/en-us/adaptive-cards/