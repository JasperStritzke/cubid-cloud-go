<template>
  <v-container fluid class="pa-10 pt-5 pl-5">
    <v-row>
      <v-col>
        <v-card-title style="word-break: break-word">Scan QR-Code and enter OTP code below</v-card-title>
        <v-card-subtitle>Recommended is <code>Google Authenticator</code> (IOS supports OTPs natively)</v-card-subtitle>
        <v-card-text>
          <v-form ref="form" @submit.prevent="finishSetup">
            <totp-input
                :rules="rules"
                :disabled="!$store.state.setup.totpTokenVisible"
                v-model="totp"
                :loading="loading"
                dense
            />
            <v-btn block color="primary" :disabled="!$store.state.setup.totpTokenVisible" type="submit"
                   :loading="loading">
              Finish Setup
            </v-btn>
          </v-form>
        </v-card-text>
      </v-col>
      <v-col>
        <v-row justify="end" align-content="center">
          <div class="totp-empty-preview">
            <v-img v-if="$store.state.setup.totpTokenVisible" max-width="256"
                   :src="`data:image/jpeg;base64,${$store.state.setup.totpQRCode}`" transition="fab-transition"/>
            <v-icon v-else size="">mdi-help-circle-outline</v-icon>
          </div>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import TotpInput from "@/components/custom_inputs/TotpInput";
import {isTotpNumber} from "@/util/numeric";

export default {
  components: {TotpInput},
  props: ["setupToken"],
  data: () => ({
    rules: [v => !!v || "OTP is required", v => v.toString().length === 9 || "OTP must be exactly 6 digits", v => isTotpNumber(v) || "TOTP must be numeric"],
    totp: "",
    loading: false,
  }),
  methods: {
    finishSetup() {
      if (this.$refs.form.validate()) {
        this.loading = true;

        axios.post(
            "/auth/setup",
            {totp: this.totp.replaceAll(" - ", "").replaceAll(" ", "").replaceAll("-", "")},
            {headers: {"X-Setup-Token": this.setupToken}}
        ).then(r => r.data).then(({username}) => {
          this.loading = false
          this.totp = ""

          this.$store.dispatch('trigger_alert', {type: "success", text: `Account ${username} was created. You can log-in now.`})
          this.$emit("finished")
        }).catch(({response}) => {
          this.loading = false

          const {data} = response

          if (data && data.message === "TOTP was invalid") {
            this.$store.dispatch('trigger_alert', {type: "warning", text: "Code was invalid!"})
            return
          }

          alert("Something went wrong. Reloading page...")
          window.location.reload()
        })
      }
    },
  }
}
</script>

<style scoped>
div.totp-empty-preview {
  --color: #696969;

  width: 256px;
  height: 256px;
  border-radius: 5px;
  border: 2px dashed var(--color);

  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  padding: 5px;
}
</style>
