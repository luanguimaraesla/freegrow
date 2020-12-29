<template>
  <div class="gadgetss container">
    <h2 class="subtitle is-3">
    Registered Gadgets
    </h2>
    <div class="columns is-multiline">
      <div v-for="gadget in gadgets" :gadget="gadget" :key="gadget.id" class="column is-one-quarter">
        <router-link :to="'/gadget/' + gadget.id">
          <GadgetCard :gadget="gadget" />
        </router-link>
      </div>
    </div>
  </div>
</template>


<script>
  import GadgetCard from '@/components/GadgetCard';
  export default {
    name: 'GadgetList',
    components : {
      GadgetCard
    },
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
    }
  }
</script>
<style lang="scss" scoped>
  .gadgets {
    margin-top: 100px;
    text-align: center;
  }
</style>

