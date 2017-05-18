Title: Features
Desc: Listing of aah framework features and capabilities.
Keywords: features, feature, aah framework, capabilities
---
### Server
  * HTTP
  * HTTPS [refer to config](https://docs.aahframework.org/app-config.html#section-ssl)
      * HTTP/2, can be disabled via configuration.
      * Certificate File and Key File.
      * [Let's Encrypt CA](https://docs.aahframework.org/app-config.html#section-lets-encrypt) - automatic install and serve certificates.
  * UNIX Socket
  * [Server Extension Points](https://docs.aahframework.org/server-extension.html) and [Request Lifecycle](https://docs.aahframework.org/request-life-cycle.html).
  * `go1.8` Graceful shutdown support.
  * Gzip Compression.

### Configuration
aah framework uses an enhanced version of [forge syntax](https://docs.aahframework.org/configuration.html) for application, route, project, security and i18n config files.

  * Environment profiles are supported. For e.g: dev, qa, prod, etc.
  * Organize your config files as you need, you can always add `include` references.  

### URL Routing and Reverse Route
  * Customized version of High performance [httprouter](https://github.com/julienschmidt/httprouter).
  * Flexible [routes configuration](https://docs.aahframework.org/routes-config.html) for application, [static files](https://docs.aahframework.org/static-files.html) and [namespace/group](https://docs.aahframework.org/routes-config.html#namespace-group-routes).
  * Supports Domains and Sub-domains.
  * Flexible reverse route URL by `route name`.
  * Access root domain and subdomain reverse routes easily from view templates and application codebase.
  * Adding Controllers with or without sub-package names for routes. So `v1`, `v2`, sub-packages are possible.
  * Redirect Trailing Slash, Auto Options, and Method Not Allowed.
  * Custom Not Found options for non found routes.

### i18n Internalization and Localization
  * [Message files](https://docs.aahframework.org/i18n.html) supported with `Language ID + Region ID` or `Language ID`.
      * Language ID follows the two-letter `ISO 639-1` standard.
      * Region ID follows the two-letter `ISO 3166-1` standard.
  * Default fallback `i18n.default` if request `Locale` is not found.
  * Messages are accessible from View template files as well as application codebase.
  * Organize your message files with sub-directories.

### Session Management
  * aah framework provides `stateful` and `stateless` HTTP state management. Default is `stateless`. Perfect fit for Web and API application, refer to [configuration](https://docs.aahframework.org/security-config.html).
  * Session data is Signed using HMAC and Encrypted using AES.
  * Out-of-the-box `cookie` and `file` session store is supported.
  * You can easily add your [own session store](https://docs.aahframework.org/session.html).

### View
  * Go view engine with [partial inheritance support](https://docs.aahframework.org/views.html).
  * Multiple view layouts for your unique use case.
  * Framework provided [template funcs](https://docs.aahframework.org/template-funcs.html), Plus you can add your own easily.
  * Custom template delimiter for templates.
  * You can add your own view engine into the framework.

### Middleware
  * Flexible [Middleware](https://docs.aahframework.org/middleware.html) with [Abort](https://docs.aahframework.org/middleware.html#abort-the-middleware-flow) feature and taking control of [response writing](https://docs.aahframework.org/reply.html#done) within the framework.
  * Bring your `http.Handler`, `http.HandlerFunc` into [aah](https://docs.aahframework.org/middleware.html#bring-go-lang-native-middleware-into-aah).

### Event Publisher/Emitter
Simple and efficient [Event Publisher](https://docs.aahframework.org/event-publisher.html) with Asynchronous and Synchronous publish. aah Server extension points built using event publisher.

### Interceptors
  * aah framework supports [per controller and per action level](https://docs.aahframework.org/interceptors.html) interceptors (`Before`, `After`, `Finally` and `Panic`).

### Reply Builder
  * Simple, efficient and chained [Reply builder](https://docs.aahframework.org/reply.html) to compose your response.
  * Supports rich reply types `HTML`, `JSON`, `JSONP`, `XML`, `Text`, `Binary`, `FileDownload`, `FileInline`, `Redirect`, etc.

### Static File Delivery
aah framework supports flexible and easy to configure [static file](https://docs.aahframework.org/static-files.html) delivery.

  * Serves directory and it's subtree files.
  * Serves individual files.
  * Directory listing.
  * All capabilities of `http.ServeContent`.

### Logger
  * Simple to use log library and it's [configuration](https://docs.aahframework.org/log-config.html).
  * Supported Log `Level`'s are `ERROR`, `WARN`, `INFO`, `DEBUG` and `TRACE`.
  * Multiple log instances are available besides the default one.
  * Out-of-the-box `Console` and `File` receivers are supported, `HTTP/HTTPS` receiver will be available in future releases.
  * Define your custom log message format in the config.

### Easy to use Application Binary
  * Easy to build and deploy [aah application binary](https://docs.aahframework.org/aah-application-binary.html).
  * Cross compile build is supported (aah is only recognizing cross compile build request and setting the appropriate values, Go lang does the rest for you).

### Essentials Library
aah helps to increase your productivity, the framework's [essentials](https://godoc.org/aahframework.org/essentials.v0) library provides a lot of useful helper/util methods in the following areas:

  * filepath
  * GUID (Globally Unique Identifier)
  * Crypto random string, Math random string, random byte generation at fixed length
  * go
  * io
  * os
  * reflect
  * string
  * archive (zip)

Refer to [godoc](https://godoc.org/aahframework.org/essentials.v0).

## Upcoming Features

Please refer to the [Roadmap](https://github.com/go-aah/aah/projects/3) for more details.
