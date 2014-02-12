Setup
=====

This section assumes Drone is installed as a service and running on Port 80. 

Admin
-------------

First you need to setup your administrative user. Navigate to http://localhost/install
where you will be prompted to create your account.


GitHub
------

Navigate to https://github.com/settings/applications and click the **Register New Application**
button. Enter the following information into the form:

Application Name
  Drone

Homepage URL
  http://localhost

Authorization Callback URL
  http://localhost/auth/login/github

Click **Register Application**

Copy and paste the **Client ID** and **Client Secret** into the Drone admin / settings screen.


Mail Server
-----------

Setting up an email server is highly recommended. Drone uses email messages
for the following:

* account activation emails
* password reset emails
* team invitation emails
* build result emails

Navigate to http://localhost/account/admin/settings and enter your SMTP server details:
