<template>
  <v-app>
    <v-main>
      <v-container fluid>
        <div class="mr-4">
          <v-form ref="form" v-model="valid" lazy-validation>
            <v-text-field v-model="name" :counter="10" :rules="nameRules" label="Name" required />
            <v-text-field v-model="email" :rules="emailRules" label="E-mail" required />
            <v-btn :disabled="!valid" color="success" class="mr-4" @click="validate">Validate</v-btn>
          </v-form>
        </div>
      </v-container>
    </v-main>
    <v-footer app />
  </v-app>
</template>

<script>
export default {
  name: 'Login',
  props: {
  },
  data() {
    return {
      input: this.input
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
