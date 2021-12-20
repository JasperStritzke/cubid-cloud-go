<template>
  <v-row justify="center">
    <v-dialog
        v-model="value"
        transition="dialog-bottom-transition"
    >
      <v-card>
        <v-toolbar
            dark
            color="primary"
        >
          <v-toolbar-title>Cloud Setup</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-flex align-self-center>
              <v-btn text @click="$emit('input', false)">Close</v-btn>
            </v-flex>
          </v-toolbar-items>
        </v-toolbar>
        <v-card-text class="mt-5">
          <v-stepper v-model="index" vertical flat>
            <!-- Setup key -->
            <v-stepper-step :complete="index > 1" step="1">
              Provide setup key
              <small>The key is provided in the <code>.setupkey</code> file in the cloud's working directory.</small>
            </v-stepper-step>

            <v-stepper-content step="1">
              <v-form ref="setupKeyForm">
                <v-card class="mb-12" flat>
                  <v-card-text>
                    <v-text-field
                        append-icon="mdi-key"
                        outlined label="Setuptoken"
                        v-model="setupKey.model"
                    />
                  </v-card-text>
                </v-card>
                <v-btn color="primary" @click="submitSetupKey">Submit Setupkey</v-btn>
              </v-form>
            </v-stepper-content>

            <!-- Setup key -->
            <v-stepper-step :complete="index > 2" step="2">
              Configure analytics for this app
            </v-stepper-step>

            <v-stepper-content step="2">
              <v-card color="grey lighten-1" class="mb-12" height="200px">

              </v-card>
              <v-btn color="primary" @click="index = 3">Continue</v-btn>
            </v-stepper-content>

            <v-stepper-step
                :complete="index > 3"
                :rules="[() => false]"
                step="3"
            >
              Select an ad format and name ad unit
            </v-stepper-step>

            <v-stepper-content step="3">
              <v-card color="grey lighten-1" class="mb-12" height="200px">

              </v-card>
              <v-btn color="primary" @click="index = 4">Continue</v-btn>
            </v-stepper-content>

            <v-stepper-step step="4">
              View setup instructions
            </v-stepper-step>
            <v-stepper-content step="4">
              <v-card color="grey lighten-1" class="mb-12" height="200px"></v-card>
              <v-btn color="primary" @click="index = 1">Continue</v-btn>
            </v-stepper-content>
          </v-stepper>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  name: "Setup",
  props: ["value"],
  methods: {
    submitSetupKey() {
      if (this.$refs.setupKeyForm.validate()) {
        this.index = 2
      }
    }
  },
  data: () => ({
    index: 1,
    rules: [() => {

    }],
    setupKey: {
      model: "",
      rules: [v => !!v || "Token is required"]
    }
  })
}
</script>

<style scoped>

</style>
