<template>
  <v-app>
    <v-main>
      <v-container>
        <div class="form-container-outer">
          <div class="form-container">
            <div class="form-container-inner">
              <div class="logo"> Login </div>
              <v-form ref="form" class="form">
                <v-text-field
                  v-model="form.email"
                  :rules="[
                    v => !!v || 'E-mail is required',
                    v => /.+@.+\..+/.test(v) || 'E-mail must be valid'
                  ]"
                  label="E-mail"
                  required
                />
                <v-text-field
                  v-model="form.password"
                  :type="'password'"
                  :rules="[
                    v => !!v || 'Password is required',
                  ]"
                  label="Password"
                  required
                />
              </v-form>
              <v-btn
                style="float: right;"
                color="primary"
                :loading="loading"
                @click="login()"
              >
                Login
              </v-btn>
            </div>
          </div>
        </div>
        <!-- Alert -->
        <v-alert
          v-model="errAlertVisible"
          dismissible
          text
          dark
          type="error"
        >
          Login Error
        </v-alert>
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
      loading: false,
      errAlertVisible: false,
      form: {
        email: '',
        password: ''
      }
    }
  },
  methods: {
    async login() {
      this.loading = true
      if (this.$refs.form.validate()) {
        this.$store.dispatch('user/login', this.form).then(() => {
          setTimeout(() => {
            this.$router.push({ path: this.redirect || '/' })
          }, 1000)
          this.loading = false
        }).catch(() => {
          this.errAlertVisible = true
          this.loading = false
          setTimeout(() => {
            this.errAlertVisible = false
          }, 10000)
        })
      } else {
        this.loading = false
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.form-container-outer {
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  height: 100%;
  margin:auto;
  .form-container{
      display:flex;
      align-items:center;
      justify-content:center;
      width: 30%;
      height: 100%;
      // border:1px solid #FFF;
      margin: auto;
    .form-container-inner{
      width: 100%;
      .logo{
        text-align: center;
        font-size: 24px;
      }
      .form{
        // padding: 16px;
        margin: 6px 0px;
        // border: 1px solid white;
        // box-shadow:0px 0px 5px rgb(100 ,100, 100);
      }
    }
  }
}
</style>
