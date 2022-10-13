This repository is inspired from the [Demo plugin](https://github.com/traefik/plugindemo) repository feel free to get back there for reference

The existing plugins can be browsed into the [Plugin Catalog](https://plugins.traefik.io).

# Plugin purpose

This plugin is intended to authorize/un-authorize requests based on a given value inside header


## Configuration & Usage 

To install this plugin on your Traefik instance you will need to do the following steps:

1. Add static configurations inside Traefik deployment (CLI or YAML):
    
    CLI:

    ````
   --experimental.plugins.headerauthentication.modulename=github.com/omar-shrbajy-arive/headerauthentication
    --experimental.plugins.headerauthentication.version=v1.0.3
   ````
   
    YAML: 

    ```(yaml)
   experimental:
    plugins:
     headerauthentication:
      moduleName: "github.com/omar-shrbajy-arive/headerauthentication"
      version: "v1.0.3"

   ```
 
2. Create the middleware and call it inside your router
    ```(yaml)
   http:
     middlewares:
        my-headerauthentication:
            plugin:
                headerauthentication:
                    header:
                        key: 123test
                        name: X-API-KEY

   ```
