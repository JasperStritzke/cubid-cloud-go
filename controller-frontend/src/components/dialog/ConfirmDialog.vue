<template>
  <v-row justify="center" style="position: absolute">
    <v-dialog
        v-model="dialogState" @input="cancel"
        :persistent="!!dialog.cancelText"
        max-width="500" transition="slide-y-transition"
    >
      <v-card class="py-5">
        <v-card-title class="headline" v-if="dialog.title">{{ dialog.title }}</v-card-title>
        <v-card-text v-if="dialog.text">
          {{ dialog.text }}
        </v-card-text>
        <v-card-actions>
          <v-spacer/>
          <v-btn v-if="dialog.cancelText" outlined :color="color" @click.stop="cancel">{{ dialog.cancelText }}</v-btn>
          <v-btn
              elevation="0" :loading="loading" :color="dialog.danger ? 'error' : 'success'"
              @click.stop="confirm"
          >
            {{ dialog.confirmButton || "Confirm" }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  name: "ConfirmDialog",
  props: ["dialog"],
  data: () => ({
    dialogState: true,
    loading: false,
  }),
  mounted() {
    if (this.dialog.waitTime) {
      this.loading = true

      setTimeout(() => {
        this.loading = false;
      }, this.dialog.waitTime)
    }
  },
  computed: {
    icon() {
      switch (this.dialog.type) {
        case "success":
          return "mdi-check-circle-outline";
        case "info":
          return "mdi-information-outline";
        case "warning":
          return "mdi-alert-decagram";
        case "error":
          return "mdi-alert";
        default:
          return "";
      }
    },
    color() {
      switch (this.dialog.type) {
        case "success":
          return "success";
        case "info":
          return "info";
        case "warning":
          return "warning";
        case "error":
          return "error";
        default:
          return "";
      }
    }
  },
  methods: {
    cancel() {
      this.dialogState = false

      this.dialog.resolve(false)
      this.$store.commit('hide_dialog', this.dialog.id)
    },
    confirm() {
      this.dialogState = false

      this.dialog.resolve(true)

      this.$store.commit('hide_dialog', this.dialog.id)
    }
  }
}
</script>

<style scoped>

</style>
