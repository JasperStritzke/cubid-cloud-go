<template>
  <v-layout column class="alert-holder">
    <v-slide-x-reverse-transition group>
      <v-alert :type="alert.type" v-for="alert in $store.state.alerts" :key="alert.id">
        {{ alert.text }}
      </v-alert>
    </v-slide-x-reverse-transition>
  </v-layout>
</template>

<style scoped>
.alert-holder {
  position: fixed;
  bottom: 15px;
  right: 15px
}
</style>
