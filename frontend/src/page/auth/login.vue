<template>
  <v-app>
    <v-main class="d-flex align-center justify-center" style="min-height: 100vh; background-color: #1f1f1f;">
      <v-card class="pa-8" max-width="400" elevation="10" color="surface">
        <v-text-field
            v-model="username"
            label="Username"
            prepend-inner-icon="mdi-account"
            variant="solo"
            class="mb-4"
        />
        <v-text-field
            v-model="password"
            label="Password"
            prepend-inner-icon="mdi-lock"
            :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
            :type="showPassword ? 'text' : 'password'"
            variant="solo"
            class="mb-6"
            @click:append-inner="showPassword = !showPassword"
        />
        <v-btn :disabled="loginDisabled" color="primary" block class="mb-4" @click="onLogin">
          Sign In
        </v-btn>
        <v-btn variant="flat" color="secondary" block>
          Connect with Metamask
        </v-btn>
      </v-card>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "PageLogin",
  data() {
    return {
      loading: false,
      username: "",
      password: "",
      showPassword: false,
    };
  },
  computed: {
    loginDisabled() {
      // if (this.loading) return true;
      return this.username.trim() === "" || this.password.trim() === "";
    },
  },
  methods: {
    onLogin() {
      const username = this.username.trim();
      const password = this.password.trim();
      if (username === "" || password === "") return;
      this.loading = true;
    },
  },
};
</script>
