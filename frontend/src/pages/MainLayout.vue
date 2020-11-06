<template>
  <v-app>
    <v-navigation-drawer app>
      <v-list
        dense
        nav
      >
        <v-list-item
          v-for="item in routes[1].children"
          :key="item.name"
          :to="item.path"
          link
        >
          <v-list-item-icon>
            <v-icon>{{ item.meta.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.meta.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <span>{{ matched }}</span>
      <span style="right: 0; position: absolute;">
        <v-menu offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              icon
              v-bind="attrs"
              v-on="on"
            >
              <v-icon>mdi-dialpad</v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item
              v-for="(item, index) in menu"
              :key="index"
              link
            >
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>

      </span>

    </v-app-bar>

    <!-- Sizes your content based upon application components -->
    <v-main>
      <!-- Provides the application the proper gutter -->
      <v-container>
        <!-- If using vue-router -->
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'Login',
  props: {
  },
  data() {
    return {
      menu: [
        {
          title: 'Logout'
        }
      ]
    }
  },
  computed: {
    routes() {
      return this.$router.options.routes
    },
    matched() {
      const matched = this.$route.matched.filter(item => item.meta && item.meta.title)
      return matched[1].name
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
  }
}
</script>
