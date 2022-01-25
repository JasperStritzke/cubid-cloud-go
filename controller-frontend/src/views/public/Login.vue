<template>
  <v-container fluid fill-height>
    <v-layout justify-center align-center column>
      <v-card max-width="450" :loading="loading">
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
            <totp-input v-model="otp" ref="totp" :disabled="loading" :rules="totpRequired"/>
          </v-card-text>
          <v-card-actions>
            <v-btn type="submit" color="primary" block :loading="loading">Sign in</v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script>
import Vue from 'vue'
import Component from "vue-class-component";
import axios from "axios";
import TotpInput from "@/components/custom_inputs/TotpInput";
import AuthorMention from "@/components/AuthorMention";

@Component({
  name: "Login",
  components: {AuthorMention, TotpInput}
})
export default class extends Vue {
  login() {
    if (this.loading) return;

    if (this.$refs.loginForm.validate()) {
      this.loading = true;

      axios.post(
          "/auth/login",
          {
            a: btoa(this.username),
            b: btoa(this.password),
            c: this.$refs.totp.getRawValue()
          }
      ).then(r => r.data).then(data => {
        this.loading = false

        this.$store.dispatch('add_snackbar', "Authenticated as " + this.username)

        this.$store.commit('set_auth_token', {token: data.token, username: this.username})

        this.$router.push("/")
      }).catch(() => {
        this.loading = false;

        this.$store.dispatch('trigger_alert', {type: "error", text: "Invalid credentials"})
      })
    }
  }

  requiredRules = [
    v => !!v || "Field is required"
  ]

  totpRequired = [
    v => v.length >= 9 || "Field must be filled out"
  ]

  loading = false

  username = ""
  password = ""
  otp = ""
}
</script>
