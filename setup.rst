Setup
=====

This section assumes Drone is installed as a service and running on Port 80. 

Configure
---------

First you need to enable one of the OAuth services using then you will bea able to login as admin
and add other authorized users if you have disabled open registration.


GitHub
------

Navigate to https://github.com/settings/applications and click the **Register New Application**
button. Enter the following information into the form:

Application Name
  Drone

Homepage URL
  http://localhost

Authorization Callback URL
  http://localhost/api/auth/github.com

Click **Register Application**

Copy and paste the **Client ID** and **Client Secret** into the Drone admin / settings screen.

BitBucket
---------

To allow Drone access to your BitBucket projects, you will need to authorize
Drone as an "Integrated Application" within Bitbucket. This will generate OAuth keys
that Drone uses to authenticate itself to Drone.

Navigate to your BitBucket account settings: https://bitbucket.org/account/user/USER_NAME/.
Click on **Integrated applications** and then press the **Add consumer** button.
This will open up a dialog window asking for **Name**, **Description**, and **URL**.
Once you complete and submit this dialog, you will be given two pieces of information:

- Key
- Secret

Now in Drone you can navigate to Drone's settings and scroll to the section 
called **Bitbucket OAuth Consumer Key and Secret** and enter the key value in
the first field, and the secret in the second field.

**Note:** In addition to being able to grant BitBucket user access to drone,
you can also associate Drone to a BitBucket **Team**. To generate keys for a
team, navigate to the team's page, click **Manage team**, and then click on the
**Integrated Applications** item in the left-hand navigation.

The first time you add a BitBucket repository to Drone, you will be asked to
authorize the Drone to BitBucket connection.

Mail Server
-----------

Setting up an email server is highly recommended. Drone uses email messages
for the following:

* account activation emails
* password reset emails
* team invitation emails
* build result emails

Navigate to http://localhost/account/admin/settings and enter your SMTP server details:
