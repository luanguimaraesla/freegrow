<template>
  <div id="app">
    <Nav v-if='showNavbar'/>
    <router-view/>
  </div>
</template>

<script>
import Nav from '@/components/partials/Nav.vue';
export default {
  name: 'App',
  components: {
    Nav
  },
  computed: {
    showNavbar() {
      console.log(this.$route)
      return this.$route.path !== '/signin' && this.$route.path !== '/signup'; 
    }
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
