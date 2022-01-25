<template>
  <v-card outlined @click.right.prevent="contextMenu = !contextMenu">
    <v-card-title>
      {{ executor.name }}
      <v-spacer/>
      <v-menu offset-overflow v-model="contextMenu">
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon small v-on="on" v-bind="attrs">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>
        <v-list>
          <v-list-item link @click="$emit('edit')">
            <v-list-item-title>Edit</v-list-item-title>
          </v-list-item>
          <v-list-item link @click="$emit('delete')">
            <v-list-item-title>Delete</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-card-title>

    <v-card-text>
      <div>Max Memory</div>
      <p class="text-h6 text--primary">
        {{ executor.maxMemory }} MB
      </p>

      <div>Max CPU</div>
      <p class="text-h6 text--primary">{{ executor.maxCpuUsage }}%</p>
    </v-card-text>
  </v-card>
</template>

<script>
import axios from "axios";

export default {
  name: "ExecutorCard",
  props: ["executor"],
  data: () => ({
    contextMenu: false
  }),
  methods: {
    deleteExecutor() {
      axios.delete("/executor/" + this.executor.name + "/")
          .then(() => {

          })
          .catch(() => {
          })
    }
  }
}
</script>
