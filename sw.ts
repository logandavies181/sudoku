/// <reference lib="webworker" />
//  https://github.com/denoland/deno/issues/15975
declare const self: ServiceWorkerGlobalScope

const webCache = `webCache`

// TODO: static check here
const urlsToCache = ["index.html", "output.css"]

self.addEventListener("install", (event) => {
  event.waitUntil(
    caches.open(webCache).then((cache) => {
      return cache.addAll(urlsToCache)
    }),
  )
})

// Network-first cache
self.addEventListener('fetch', event => {
  event.respondWith(
    fetch(event.request)
      .then(networkResponse => {
        if (networkResponse.ok) {
          caches.open(webCache).then(cache => {
            cache.put(event.request, networkResponse.clone())
          })
        }
        return networkResponse.clone();
      })
      .catch(() => {
        return caches.match(event.request).then((request) => {
          if (!request) {
            throw "couldn't find event in cache"
          }
          return request
        })
      })
  )
})
