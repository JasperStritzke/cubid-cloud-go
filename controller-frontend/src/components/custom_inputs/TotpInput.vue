<template>
  <v-text-field
      label="6 digit OTP" outlined maxlength="9" :dense="dense"
      :rules="rules" :disabled="disabled"
      :value="value"
      @input="input"
      append-icon="mdi-timer"
      @keydown="cancelSpace"
      placeholder="### - ###"
      :loading="loading"
  />
</template>

<script>
import {isTotpNumber} from "@/util/numeric";

export default {
  name: "TotpInput",
  props: ["value", "loading", "disabled", "rules", "dense"],
  methods: {
    input(v) {
      this.$emit("input", v)

      let computed = v.replaceAll(" - ", "").replaceAll(" ", "").replaceAll("-", "")

      if (computed.length > 3) {
        computed = computed.substring(0, 3) + " - " + computed.substring(3)
      }

      this.$nextTick(() => {
        this.$emit("input", computed)
      })
    },
    getRawValue() {
      return this.value.replaceAll(" - ", "").replaceAll(" ", "").replaceAll("-", "")
    },
    cancelSpace(e) {
      if (e.key.length >= 4) return

      if (!isTotpNumber(e.key) || e.key === " ") {
        e.preventDefault()
      }
    },
  }
}
</script>

<style scoped>

</style>
