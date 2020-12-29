<template>
<div class="Signup">
  <div class="hero-body">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-5-tablet is-4-desktop is-3-widescreen">
          <form class="box">
            <div class="field">
              <label for="email" class="label">Email</label>
              <div class="control has-icons-left">
                <input type="email" v-model="user.email" :class="{ 'is-invalid': submitted && errors.has('email') }" placeholder="e.g. bobsmith@gmail.com" class="input" required>
                <span class="icon is-small is-left">
                  <i class="fa fa-envelope"></i>
                </span>
              </div>
            </div>
            <div class="field">
              <label for="" class="label">Username</label>
              <div class="control has-icons-left">
                <input type="text" v-model="user.username" :class="{ 'is-invalid': submitted && errors.has('username') }" placeholder="e.g. bobsmith" class="input" required>
                <span class="icon is-small is-left">
                  <i class="fa fa-person"></i>
                </span>
              </div>
            </div>
            <div class="field">
              <label for="" class="label">Password</label>
              <div class="control has-icons-left">
                <input type="password" v-model="user.password" :class="{ 'is-invalid': submitted && errors.has('password') }" placeholder="********" class="input" required>
                <span class="icon is-small is-left">
                  <i class="fa fa-lock"></i>
                </span>
              </div>
            </div>
            <div class="field">
              <button class="button is-success" @click="handleSubmit" >
                Signup
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
var userService = require('../_services/user');

export default {
  name: 'Signup',
  components: {
  },
  data () {
    return {
      user: {
        email: '',
        username: '',
        password: ''
      },
      submitted: false
    }
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault()
      this.submitted = true
      userService.signup(this.user);
    }
  }
}
</script>

<style lang="scss" scoped>
</style>

