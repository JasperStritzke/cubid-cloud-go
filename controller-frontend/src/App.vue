<template>
  <v-app>
    <v-main>
      <router-view/>
      <alert-holder/>
      <snackbar-holder/>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import AlertHolder from "@/components/alert/AlertHolder.vue";
import SnackbarHolder from "@/components/alert/SnackbarHolder.vue";

export default Vue.extend({
  name: 'App',
  components: {SnackbarHolder, AlertHolder},
  data: () => ({
    //
  }),
});
</script>
