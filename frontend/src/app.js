import AuthOnchain from "app.vue"
import routes from "app/routes"
import { createRouter, createWebHistory } from "vue-router"
import { createApp } from "vue"
import { createVuetify } from "vuetify"
import { aliases, mdi } from "vuetify/iconsets/mdi"
import "vuetify/styles";
import "@mdi/font/css/materialdesignicons.css"

let app = createApp(AuthOnchain)

// Create Vuetify 3 instance.
const vuetify = createVuetify({
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: { mdi },
  },
});

// Use Vuetify 3.
app.use(vuetify)

// Configure client-side routing.
const router = createRouter({
  history: createWebHistory("/"),
  routes: routes,
});

// Use router.
app.use(router)

// Mount to #app.
app.mount("#app");
