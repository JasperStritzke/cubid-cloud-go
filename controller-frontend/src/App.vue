<template>
  <v-app>
    <sidebar-and-app-bar v-if="$store.getters.isAuthenticated"/>

    <v-main>
      <v-breadcrumbs :items="computedRoutes" v-if="$store.getters.isAuthenticated && computedRoutes.length > 1"/>
      <div class="ma-5 mt-0">
        <v-divider class="mb-5" v-if="$store.getters.isAuthenticated"/>
        <author-mention/>

        <setup v-model="$store.state.setupMode" class="setup-modal"/>
      </div>

      <router-view class="fade-in"/>

      <div class="important_notifications">
        <alert-holder/>
        <snackbar-holder/>
        <dialog-holder/>
      </div>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import AlertHolder from "@/components/alert/AlertHolder.vue";
import SnackbarHolder from "@/components/alert/SnackbarHolder.vue";
import Setup from "@/components/dialog/Setup.vue";
import SidebarAndAppBar from "@/components/layout/SidebarAndAppBar.vue";
import AuthorMention from "@/components/AuthorMention.vue";
import DialogHolder from "@/components/alert/DialogHolder.vue";

export default Vue.extend({
  name: 'App',
  components: {DialogHolder, AuthorMention, SidebarAndAppBar, SnackbarHolder, AlertHolder, Setup},
  computed: {
    computedRoutes() {
      let pathArray = this.$route.path.split("/")
      pathArray.shift()

      let index = 0

      let to = "/"
      return pathArray.map(path => {
        const alt = index === 0 ? this.$route.name : path.charAt(0).toUpperCase() + path.substr(1)
        const name = this.$route.matched[index]?.meta?.breadcrumb || alt
        index++

        to = to + path + "/"

        return {
          text: name,
          to: index > 1 ? to : undefined,
        }
      })
    }
  }
});
</script>
<style>
/* setup modal should be ignored in layout */
.setup-modal {
  position: absolute;
}

.important_notifications {
  z-index: 9999999999999999999999 !important;
}

.fade-in {
  animation: fadeIn .3s forwards;
  opacity: 0.2;
  transform: translateY(-15px);
}

@keyframes fadeIn {
  from {
    opacity: 0.2;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
