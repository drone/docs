+++
weight = 5
date = "2015-01-13T21:07:00+03:00"
title = "Bintray"

[menu.docs]
parent = "Publish"
+++

Before drone send your files into bintray, you need create appropriate package within the repository from bintray control panel.

![Bintray Create Package](/img/bintray_create_package.png)


```coffeescript
publish:
  bintray:
    username: user
    api_key: 970d28f4dd477bc184fbd10b376de753
    packages:
      - file: path/to/deb/file.deb
        owner: my_organization_or_username
        type: deb
        repository: repository_for_debian_packages
        package: awesome_package
        version: 0.4.2
        target: file-0.4.2.deb
        distr: ubuntu
        component: main
        arch:
          - amd64
        publish: true
        override: true
      - file: path/to/rpm/file.rpm
        owner: my_organization_or_username
        type: rpm
        repository: repository_for_rpm_packages
        package: awesome_package
        version: 0.4.2
        target: file-0.4.2.rpm
        publish: true
        override: true
      - file: path/to/rpm/file.jar
        owner: my_organization_or_username
        type: maven
        repository: repository_for_maven_packages
        package: awesome_package
        version: 0.4.2
        target: file-0.4.2.jar
        publish: true
```

Packages signing (GPG)
----------------

For example you can generate not expire ( but you can set your own expire time: days, weeks, months, years ) at all key, <u>also you need empty passphrase if you want bintray sign your packages automatically</u>

```bash
gpg --gen-key
gpg (GnuPG) 1.4.18; Copyright (C) 2014 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Please select what kind of key you want:
   (1) RSA and RSA (default)
   (2) DSA and Elgamal
   (3) DSA (sign only)
   (4) RSA (sign only)
Your selection? 4
RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (2048)
Requested keysize is 2048 bits
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 0
Key does not expire at all
Is this correct? (y/N) y

Real name: MyOrganizationOrUsername
Email address: gpg@example.com
Comment:
You selected this USER-ID:
    "MyOrganizationOrUsername <gpg@example.com>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? O
You need a Passphrase to protect your secret key.

You dont want a passphrase - this is probably a *bad* idea!
I will do it anyway.  You can change your passphrase at any time,
using this program with the option "--edit-key".

We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
...+++++
+++++
gpg: key 14F4CA71 marked as ultimately trusted
public and secret key created and signed.

gpg: checking the trustdb
gpg: 3 marginal(s) needed, 1 complete(s) needed, PGP trust model
gpg: depth: 0  valid:   3  signed:   0  trust: 0-, 0q, 0n, 0m, 0f, 3u
pub   2048R/14F4CA71 2015-01-13
      Key fingerprint = 61C8 BABD EA98 B8B4 A6D9  B6AA 442A A50E 14F4 CA71
uid                  MyOrganizationOrUsername <gpg@example.com>
```

Good job, your GPG key has been created, now you need extract your private and public keys into the files
```bash
gpg --export-secret-key -a "MyOrganizationOrUsername" > private.key
gpg --export –a "MyOrganizationOrUsername" > public.key
```

Then you need go to your user or organization ( if you publish packages in organization repository ) profile on bintray and press edit. Go to **GPG Signing**, and insert your keys from `private.key` and `public.key` in the appropriate fields, after press update button. Well done, now when drone send a new package into your bintray repository, package will be automatically signed.