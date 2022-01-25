<template>
  <div>
    <v-slide-y-transition v-for="dialog in $store.state.dialogs" :key="dialog.id">
      <confirm-dialog :dialog="dialog"/>
    </v-slide-y-transition>
  </div>
</template>

<script>
import Vue from 'vue'
import Component from "vue-class-component";
import ConfirmDialog from "@/components/dialog/ConfirmDialog";

@Component({
  name: "DialogHolder",
  components: {ConfirmDialog}
})
export default class extends Vue {
  t = true
}
</script>

<style scoped>

</style>
