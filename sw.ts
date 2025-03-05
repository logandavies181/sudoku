/// <reference lib="webworker" />
//  https://github.com/denoland/deno/issues/15975
declare const self: ServiceWorkerGlobalScope

const sv = "0.1.0"

// stored html/js/css
const webCache = `web-${sv}`

// TODO: static check here
const urlsToCache = ["index.html", "hex.svg", "main.js", "output.css"]

self.addEventListener("install", (event) => {
  event.waitUntil(
    caches.open(webCache).then((cache) => {
      return cache.addAll(urlsToCache)
    }),
  )
})

self.addEventListener("fetch", (event) => {
  event.respondWith(
    caches.match(event.request).then((request) => {
      return request || fetch(event.request)
    }),
  )
})
