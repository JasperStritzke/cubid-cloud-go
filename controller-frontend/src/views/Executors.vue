<template>
  <cubid-view>
    <p>Executors ({{ executors.length }})</p>
    <v-container fluid>
      <v-row>
        <v-col v-for="executor in executors" :key="executor.name" cols="12" xl="2" lg="3" md="4" sm="6">
          <executor-card
              :executor="executor"
              @edit="editExecutor(executor)"
              @delete="deleteExecutor(executor)"
          />
        </v-col>
        <v-col cols="12" xl="2" lg="3" md="4" sm="6">
          <v-btn elevation="0" @click="create" color="secondary" block>+</v-btn>
        </v-col>
      </v-row>
    </v-container>

    <executor-edit-dialog
        ref="executorEditDialog"
        @submit="handleSubmit"
    />
  </cubid-view>
</template>

<script>
import Vue from "vue"
import Component from "vue-class-component";
import ExecutorCard from "@/components/executor/ExecutorCard";
import {alertErrors, alertErrorsWithCustomMessage} from "@/util/alertErrors";
import ExecutorEditDialog from "@/components/executor/ExecutorEditDialog";
import Setup from "@/components/dialog/Setup";
import ConfirmDialog from "@/components/dialog/ConfirmDialog";
import {createExecutor, deleteExecutor, getExecutors} from "@/data/managers/executor_manager";
import CubidView from "@/views/CubidView";

@Component({
  name: "Executors",
  components: {CubidView, ConfirmDialog, Setup, ExecutorEditDialog, ExecutorCard},
})
export default class extends Vue {
  executors = []

  mounted() {
    this.fetchExecutors();
  }

  fetchExecutors() {
    getExecutors()
        .then(executors => {
          if (executors !== null) this.executors = executors.sort((a, b) => a.name.localeCompare(b.name))
        }).catch(alertErrors)
  }

  handleSubmit({createNew, executor}) {
    this.$refs.executorEditDialog.close()

    if (createNew) {
      createExecutor(executor)
          .then(this.fetchExecutors)
          .catch(alertErrorsWithCustomMessage(`Executor ${executor.name} already exists`))
    }
  }

  create() {
    this.$refs.executorEditDialog.createNew()
  }

  editExecutor(executor) {
    this.$refs.executorEditDialog.edit(executor)
  }

  deleteExecutor(executor) {
    this.$store.dispatch('show_dialog', {
      title: `Delete Executor ${executor.name}?`,
      text: "This can't be undone!",
      type: "warning",
      danger: true,
      cancelText: "Cancel",
      waitTime: 2000
    }).then(state => {
      if (!state) return

      deleteExecutor(executor.name)
          .then(() => {
            this.fetchExecutors()
            this.$store.dispatch('trigger_alert', {
              type: "success",
              text: `Executor ${executor.name} was deleted.`
            })
          }).catch(alertErrorsWithCustomMessage(`Executor ${executor.name} could not be deleted.`))
    });
  }
}
</script>
