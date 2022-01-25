<template>
  <div>

    <v-navigation-drawer app v-model="drawer" :temporary="$vuetify.breakpoint.mobile">
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="text-h6">
            cubid.cloud
          </v-list-item-title>
          <v-list-item-subtitle>
            Logged in as <b>{{ $store.state.auth.userName }}</b>
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
          dense
          nav
      >
        <v-list-item
            v-for="item in items"
            :key="item.title"
            @click="goto(item.link)"
            link
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <template v-slot:append>
        <v-list dense nav>
          <v-list-item
              v-for="item in endingLinks"
              :key="item.title"
              @click="goto(item.link)"
              link
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item link @click="$store.dispatch('logout')">
            <v-list-item-icon>
              <v-icon color="red">mdi-logout</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>Logout</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </template>
    </v-navigation-drawer>

    <v-app-bar dark app color="primary" elevation="0" absolute>
      <v-app-bar-nav-icon @click="drawer = !drawer"/>
      <v-app-bar-title v-text="routeName"/>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  computed: {
    routeName() {
      if (this.$route.name) {
        return this.$route.name
      }

      return "Error"
    }
  },
  methods: {
    goto(link) {
      if (this.$router.currentRoute.path === link) return

      this.$router.push(link)
    }
  },
  data: () => ({
    drawer: true,
    items: [
      {title: "Dashboard", icon: "mdi-view-dashboard", link: "/"},
      {title: "Executors", icon: "mdi-cogs", link: "/executors"},
      {title: "Groups", icon: "mdi-group", link: "/groups"},
      {title: "Servers", icon: "mdi-server-network", link: "/servers"},
      {title: "Proxy", icon: "mdi-arrow-decision", link: "/proxy"},
      {title: "Players", icon: "mdi-account-multiple", link: "/players"},
      {title: "Permissions", icon: "mdi-key", link: "/permissions"},
    ],
    endingLinks: [
      {title: "Account", icon: "mdi-account", link: "/account"},
      {title: "Accounts", icon: "mdi-account-edit", link: "/accounts"},
    ]
  })
}
</script>

<style scoped>

</style>
