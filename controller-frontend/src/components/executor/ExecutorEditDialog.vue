<template>
  <v-row justify="center">
    <v-dialog
        v-model="visible"
        max-width="500" transition="scroll-y-transition"
    >
      <v-form ref="form" @submit.prevent="submit">
        <v-card>
          <v-toolbar color="primary" flat dark>
            <v-toolbar-title>{{ isEdit ? "Edit" : "Create" }} executor {{ localExecutor.name }}</v-toolbar-title>
          </v-toolbar>
          <v-card-text class="mt-5">
            <v-text-field
                v-model="localExecutor.name" label="Name" outlined
                placeholder="Executor-EU-WEST-1" :rules="[requiredRule, nameRule]"
            />
            <v-slider
                thumb-label min="0" max="100" v-model="localExecutor.maxCpuUsage" label="Max CPU Usage"
                step="5" persistent-hint class="mb-2 pr-1" :rules="[requiredRule, cpuRule]"
                hint="If CPU Usage on the executors server reaches this level, it won't start new services">
              <template v-slot:thumb-label="{value}">
                {{ value }}%
              </template>
            </v-slider>
            <v-text-field
                type="number" min="0" step="512" label="Allocatable Memory in MB"
                v-model="localExecutor.maxMemory" :rules="[memoryRule, requiredRule]"
            />
          </v-card-text>
          <v-card-actions>
            <v-spacer/>
            <v-btn type="submit" color="success" elevation="0">{{ isEdit ? "Save" : "Create" }}</v-btn>
          </v-card-actions>
        </v-card>
      </v-form>
    </v-dialog>
  </v-row>
</template>

<script>
const emptyExecutor = {
  name: "",
  maxMemory: 2048,
  maxCpuUsage: 75,
}

export default {
  name: "ExecutorEditDialog",
  data: () => ({
    visible: false,
    isEdit: false,
    localExecutor: {
      name: "error",
      maxCpuUsage: -1,
      maxMemory: -1
    },
    requiredRule: v => !!v || "Field is required",
    nameRule: v => !v.includes("?") || "Name can't contain question marks.",
    cpuRule: v => v > 0 && v <= 100 || "CPU-Usage must be between 0 and 100%",
    memoryRule: v => v >= 128 || "Memory must be at least 128MB",
  }),
  methods: {
    createNew() {
      this.localExecutor = Object.assign({}, emptyExecutor)

      this.isEdit = false

      this.visible = true
    },
    edit(executor) {
      this.localExecutor = executor
      this.isEdit = true

      this.visible = true
    },
    close() {
      this.visible = false
    },
    submit() {
      if (!this.$refs.form.validate()) return

      this.localExecutor.maxMemory = +this.localExecutor.maxMemory

      this.$emit("submit", {
        createNew: !this.isEdit,
        executor: this.localExecutor
      })
    }
  }
}
</script>

<style scoped>

</style>
