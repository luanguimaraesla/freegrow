<template>
  <div class="gadget-single">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            {{ gadget.name }}
          </h1>
          <h2 class="subtitle ">
            {{ gadget.type }}
          </h2>
        </div>
      </div>
    </section>
    <section class="gadget-content">
      <div class="container">
        <p class="is-size-4">State: {{ gadget.state }}</p>
      </div>
    </section>
  </div>
</template>
<script>
  export default {
    name: 'GadgetSingle',
    data () {
      return {
        gadget: {},
        gadgets: [
          {
            id: 1,
            name: 'General Light',
            state: 'on',
            type: 'Light'
          },
          {
            id: 2,
            name: 'Irrigator A',
            state: 'off',
            type: 'Irrigator'
          }
        ]
      }    
    },
    created() {
      const ID = Number(this.$route.params.id);
      let gadget = this.gadgets.find(gadget => gadget.id === ID);
      this.gadget = gadget;
    }
  }
</script>
<style lang="scss" scoped>
  .gadget-single {
    margin-top: 30px;
  }
  .hero {
    margin-bottom: 70px;
  }
  .gadget-images {
    margin-top: 50px;
  }
  .description {
    margin-bottom: 30px;
  }
</style>

