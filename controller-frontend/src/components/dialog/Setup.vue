<template>
  <v-row justify="center">
    <v-dialog
        v-model="$store.state.setup.setupMode"
        transition="scroll-y-transition"
        hide-overlay
        persistent
        max-width="750"
    >
      <v-card>
        <v-toolbar
            dark
            color="primary"
        >
          <v-toolbar-title>Cloud Setup</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-flex align-self-center>
              <v-tooltip left transition="slide-x-transition">
                <template v-slot:activator="{on, attrs}">
                  <v-btn text icon v-on="on" v-bind="attrs">
                    <v-icon>mdi-help-circle</v-icon>
                  </v-btn>
                </template>
                You need to create a super-user to start using the cloud
              </v-tooltip>
            </v-flex>
          </v-toolbar-items>
        </v-toolbar>
        <v-card-text class="mt-5">
          <v-card-title>
            Provide setup key
          </v-card-title>
          <v-card-subtitle>
            You can find it in the <code>setupkey</code> file
          </v-card-subtitle>
          <v-form ref="superUserForm" @submit.prevent="createSuperUser">
            <v-card-text>
              <v-text-field
                  append-icon="mdi-key" label="Setuptoken" counter type="password"
                  v-model="setupKey.model" :rules="setupKey.rules"
                  :disabled="loading || disabled"
              />
            </v-card-text>
            <v-card-title>
              Create an administrator
            </v-card-title>
            <v-card-text>
              <v-text-field
                  :disabled="loading || disabled"

                  label="Username" v-model="user.value"
                  :rules="user.rules"
              />
              <v-row>
                <v-col>
                  <v-text-field
                      :disabled="loading || disabled"
                      label="Password" type="password" v-model="passwordOne"
                      :rules="[requiredRule, atLeastSixChars]"
                      @input="$refs.superUserForm.validate()" clearable
                  />
                </v-col>
                <v-col>
                  <v-text-field
                      :disabled="loading || disabled"
                      label="Repeat Password" type="password" v-model="passwordTwo"
                      :rules="[requiredRule, () => passwordsDoMatch || 'Passwords don\'t match.']"
                      clearable
                  />
                </v-col>
              </v-row>
            </v-card-text>
            <v-card-actions>
              <v-btn color="primary" type="submit" :loading="loading" :disabled="disabled">Start Verification</v-btn>
            </v-card-actions>
          </v-form>
        </v-card-text>

        <totp-preview :setup-token="setupKey.model" @finished="finishHandler"/>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import axios from "axios";
import {alertErrorsWithCustomMessage} from "@/util/alertErrors";
import TotpPreview from "@/components/dialog/TotpPreview";

export default {
  name: "Setup",
  components: {TotpPreview},
  methods: {
    createSuperUser() {
      if (this.$refs.superUserForm.validate()) {
        this.loading = true

        axios.post("/auth/pre-setup",
            {username: this.user.value, password: btoa(this.passwordOne)},
            {headers: {"X-Setup-Token": this.setupKey.model}}
        ).then(r => r.data)
            .then(data => {
              const totpBase64 = data['totp']

              this.$store.commit('view_totp', totpBase64)

              this.loading = false;
              this.disabled = true
            })
            .catch(() => {
              alertErrorsWithCustomMessage("Setup token was invalid!")()

              this.loading = false
              this.setupKey.model = ""
            });
      }
    },
    finishHandler() {
      this.$store.commit('view_totp')
    }
  },
  computed: {
    passwordsDoMatch() {
      return this.passwordOne === this.passwordTwo
    }
  },
  data: () => ({
    setupKey: {
      model: "",
      rules: [v => !!v || "Token is required"]
    },
    disabled: false,
    loading: false,
    requiredRule: v => {
      if (v === "123123" || v === "123456") {
        return "Just don't!"
      }

      return !!v || "Password is required."
    },
    atLeastSixChars: v => v && v.length >= 6 || "Passwords must be at least 6 chars",
    user: {
      rules: [v => v && v.length >= 3 || "Username must be at least 3 characters"],
      value: "",
    },
    passwordOne: "",
    passwordTwo: "",
  })
}
</script>

<style scoped>

</style>
