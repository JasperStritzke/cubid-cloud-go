<template>
  <v-snackbar v-model="$store.state.snackbar.value" @input="onHide">{{ $store.state.snackbar.text }}</v-snackbar>
</template>

<script>
export default {
  name: "SnackbarHolder",
  methods: {
    onHide() {
      this.$store.dispatch('next_snackbar')
    }
  }
}
</script>

<style scoped>

</style>
