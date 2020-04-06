<template>
  <div id="app">
    <SwitchButton msg="Welcome to Your Vue.js App"/>
    <Loading v-if='LOADING'/>
  </div>
</template>

<script>
import SwitchButton from './components/Switch.vue'
import Loading from './components/Loading.vue'
import {mapState} from 'vuex'
export default {
  name: 'app',
  components: {
    SwitchButton,
    Loading
  },
  computed:{
    ...mapState([
      'LOADING'
    ])
  }
}
</script>

<style>
</style>
