# Go Service with Local Build and Dokku Deploy

Copy the contents of `git-hooks/pre-push.sh` to `.git/hooks/pre-push.sample` and rename to `pre-push`

Modify the script with with your Docker Hub repo name, image name, Droplet domain/address, and Dokku app name.

Ensure your Droplet is configured and accessible via SSH, and you are logged in to Docker Hub.

Domain, proxying and SSL certificates should already be set up, and the application must be created.

This application deploys to the root domain.

```
dokku apps:create <app-name>
dokku domains:clear-global
dokku domains:set <app-name> <domain>
dokku proxy:ports-set <app-name> http:80:<port-exposed-by-docker-container>
```

### SSL Certs

If adding an existing cert/key pair, they must be named `server.crt` and `server.key` and put in a `.tar` file named `cert-key.tar` then uploaded to your Droplet.

```
tar cvf cert-key.tar server.crt server.key
scp cert-key.tar root@<domain.tld>:/root
```

In your Droplet:

```
dokku certs:add <app-name> < cert-key.tar
```

Alternatively, you can use the [dokku-letsencrypt](https://github.com/dokku/dokku-letsencrypt) pluggin.

\*letsencrypt will rate limit to 5 API calls per 7 day period.
