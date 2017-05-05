Title: Features
Desc: Listing of aah framework features and capabilities.
Keywords: features, feature, aah framework, capabilities
---
# Features

List of feature sets provided by aah framework.

### Server
  * HTTP
  * HTTPS [config reference](https://docs.aahframework.org/app-config.html#section-ssl)
      * HTTP/2, optionally you can disable using configuration.
      * Certificate File and Key File.
      * [Let's Encrypt CA](https://docs.aahframework.org/app-config.html#section-lets-encrypt) - automatic install and serve certificates.
  * UNIX Socket
  * [Server Extension Points](https://docs.aahframework.org/server-extension.html) and [Request Lifecycle](https://docs.aahframework.org/request-life-cycle.html).
  * `go1.8` Graceful shutdown support.

### Configuration
aah framework uses [forge syntax](https://docs.aahframework.org/configuration.html) developed by [@brettlangdon](https://github.com/brettlangdon") similar to HOCON syntax, not 100%. Enhanced by author of aah framework [@jeevatkm](https://github.com/jeevatkm) for application, routes, project, security, i18n config files.

  * Environment profile is supported. For e.g: dev, qa, prod, etc.
  * Feel free organize your config files as per need and use case with `include` reference.  

### URL Routing and Reverse Route
  * Customized version of High performance [httprouter](https://github.com/julienschmidt/httprouter) by [@julienschmidt](https://github.com/julienschmidt).
  * Flexible [routes configuration](https://docs.aahframework.org/routes-config.html) for application, [static files](https://docs.aahframework.org/static-files.html) and [namespace/group](https://docs.aahframework.org/routes-config.html#namespace-group-routes).
  * Supports Domain and Sub-domain.
  * Flexible reverse route URL by `route name`.
  * Access root domain and subdomain reverse routes easily from view templates and application codebase.
  * Adding Controller with or without sub-package name for routes. So `v1`, `v2`, packages is possible.
  * Redirect Trailing Slash, Auto Options, and Method Not Allowed.
  * Custom Not Found option for not found routes.

### i18n Internalization and Localization
  * [Message files](https://docs.aahframework.org/i18n.html) supported with `Language ID + Region ID` or `Language ID`.
      * Language ID is as per two-letter `ISO 639-1` standard.
      * Region ID is as per two-letter `ISO 3166-1` standard.
  * Default fallback `i18n.default` if request `Locale` is not found.
  * Message is accessible from View template files as well as application codebase.
  * Feel free organize your message file with sub-directories.

### Session Management
  * aah framework provides `stateful` and `stateless` HTTP state management. Default is `stateless`. Perfect fit for Web and API application, refer [configuration](https://docs.aahframework.org/security-config.html).
  * Session data is Signed using HMAC and Encrypted using AES.
  * Out-of-the-box `cookie` and `file` session store is supported.
  * You can easily add your [own session store](https://docs.aahframework.org/session.html).

### View
  * Go view engine with [partial inheritance support](https://docs.aahframework.org/views.html).
  * Multiple view layouts for your unique use case.
  * Framework provided [template funcs](https://docs.aahframework.org/template-funcs.html), Plus you can add your own easily.
  * Custom template delimiter for templates.
  * You can add your own view engine into framework.

### Middleware
  * Flexible [Middleware](https://docs.aahframework.org/middleware.html) with [Abort](https://docs.aahframework.org/middleware.html#abort-the-middleware-flow) feature and taking control of [response writing](https://docs.aahframework.org/reply.html#done) from framework.
  * Bring your `http.Handler`, `http.HandlerFunc` into [aah](https://docs.aahframework.org/middleware.html#bring-go-lang-native-middleware-into-aah).

### Event Publisher/Emitter
Simple and efficient [Event Publisher](https://docs.aahframework.org/event-publisher.html) with Asynchronous and Synchronous publish. aah Server extension points built using event publisher.

### Interceptors
  * aah framework supports [per controller and per action level](https://docs.aahframework.org/interceptors.html) interceptors (`Before`, `After`, `Finally` and `Panic`).

### Reply Builder
  * Simple, efficient and chained [Reply builder](https://docs.aahframework.org/reply.html) to compose your response.
  * Supported reply types `HTML`, `JSON`, `JSONP`, `XML`, `Text`, `Bytes`, `File`, `FileInline`, `Redirect`, etc.

### Static File Delivery
aah framework supports flexible easy to [configure static files](https://docs.aahframework.org/static-files.html) delivery.

  * Serve directory and it's subtree files.
  * Serve individual file.
  * Directory listing.
  * All capabilities of `http.ServeContent`.

### Logger
  * Simple to use [log library and it's configuration](https://docs.aahframework.org/log-config.html).
  * Supported Log `Level`'s are `ERROR`, `WARN`, `INFO`, `DEBUG` and `TRACE`.
  * Multiple log instance if you want use beside the default one.
  * Out-of-the-box `Console` and `File` receiver is supported, `HTTP/HTTPS` receiver will be in upcoming release.
  * Define your own log message format in the config.

### Easy to use Application Binary
  * Easy to build and deploy [aah application binary](https://docs.aahframework.org/aah-application-binary.html).
  * Cross compile build is supported (I'm not doing much expect recognizing cross compile build requested and setting values appropriately, rest Go does it for you).

### Essentials Library
aah framework provides go library with lot of useful helper/util methods on following areas. It helps to increase your productivity instead of re-inventing a wheel. Refer [godoc](https://godoc.org/aahframework.org/essentials.v0).

  * filepath
  * GUID (Globally Unique Identifier)
  * Crypto random string, Math random string, random byte generation at fixed length
  * go
  * io
  * os
  * reflect
  * string
  * archive (zip)

## Upcoming Features

Please refer [Roadmap](https://github.com/go-aah/aah/projects/3) for more details.
