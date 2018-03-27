# Installation instructions

1. Install Go language
2. Run
```sh
curl https://raw.githubusercontent.com/mewa/wuff/master/install | sh
```
3. (Optional) Link executables

You should have 2 executables at your $HOME/.wuff path:

* `runwuff` is used to start wuff as an independent background process and is just a wrapper to `wuff`.
* `wuff` is the main program. It may be useful to run `wuff` to verify there are no config errors.

# Example config

`wuff` expects a config file under path `$HOME/.wuff/config.hcl` (in HCL format).

An example config can be found below:

```hcl
email = "notifyme@example.com" # email receiving notifications

smtp.server = "smtp.example.com" # smtp server address used
smtp.port = 25 # smtp server port, can be omitted (defaults to 587)
smtp.user = "user@example.com" # user used to authenticate against smtp server
smtp.password = "examplepassword" # password used to authenticate against smtp server
smtp.sender = "from@example.com" # sender email address

# service declaration block
# for each such block checks will be made (independently)
service {
  name = "httpd"
  check = "service httpd status" # command used to check service health, defaults to "service ${service.name} status"
  start = "service httpd start" # command used to start service when it's down, defaults to "service ${service.name} start"
  checkPeriod = 30 # amount of time in seconds to wait between successful checks, defaults to 10
  retries = 3 # number of retries, before the checks will quit and a failure notification will be sent, defaults to 0
  retryPeriod = 10 # amount of time in seconds to wait between failed checks, defaults to 0 (immediate retries)
}
```

# Notes
As `wuff` is managing services, it need appropriate permissions to restart them.
