<template>
  <v-container fluid fill-height>
    <v-layout justify-center align-center column>
      <v-card max-width="450" min-width="350" :loading="loading">
        <v-card-title>Login</v-card-title>
        <v-card-subtitle>Sign in using your credentials provided by cubid.</v-card-subtitle>
        <v-form @submit.prevent="login" ref="loginForm">
          <v-card-text>
            <v-text-field
                outlined append-icon="mdi-account"
                label="Username" :rules="requiredRules"
                v-model="username" :disabled="loading"
            />
            <v-text-field
                outlined append-icon="mdi-key"
                label="Password" type="password"
                :rules="requiredRules"
                v-model="password" :disabled="loading"
            />
          </v-card-text>
          <v-card-actions>
            <v-btn type="submit" color="primary" block :loading="loading">Sign in</v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
      <div class="author-mention">cubid.cloud &copy; {{ new Date().getFullYear() }} - made by Jasper S.</div>
      <setup v-model="setup" class="setup-modal"/>
    </v-layout>
  </v-container>
</template>

<script>
import Vue from 'vue'
import Component from "vue-class-component";
import axios from "axios";
import Setup from "@/views/public/Setup";

@Component({
  name: "Login",
  components: {Setup}
})
export default class extends Vue {
  login() {
    if (this.loading) return;

    if (this.$refs.loginForm.validate()) {
      this.loading = true;

      axios.post(
          "/login",
          {
            auth: btoa(`${this.username}:${this.password}`),
          },
      ).then(result => {
        this.username = ""
        this.password = ""
        this.loading = false
      }).catch(() => {
        this.loading = false;
        this.$store.dispatch('trigger_alert', {type: "error", text: "An error occured while logging in."})
      })
    }
  }

  setup = false

  requiredRules = [
    v => !!v || "Field is required"
  ]

  loading = false

  username = ""
  password = ""
}
</script>
<style>
div.author-mention {
  position: fixed;
  bottom: 30px;
  opacity: .5;
  user-select: none;
}

/* setup modal should be ignored in layout */
.setup-modal {
  position: absolute;
}
</style>
