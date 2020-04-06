<template>
  <div class="sw-box">
    <div class="ckbx-style-14">
      <input type="checkbox" id="ckbx-style-14-1" v-model="status" @click="turn">
      <label for="ckbx-style-14-1"></label>
    </div>
  </div>
</template>

<script>
  import request from '@/utils/request'
  import 'css-checkbox-library/dist/css/checkboxes.min.css'

  export default {
    name: 'SwitchButton',
    data (){
      return {
        status: false,
      }
    },
    methods: {
      async turn() {
        this.$store.commit('showLoading')
        if (this.status) {
          this.status = false
          await request.post('/off')
        } else {
          this.status = true
          await request.post('/on')
        }
        this.$store.commit('hideLoading')
      },
      async get() {
        this.$store.commit('showLoading')
        await request.get('/status').then(response => {
          this.status = response.data.light
        })
        this.$store.commit('hideLoading')
      }
    },
    mounted() {
      this.get()
    },
    computed: {
      text() {
        return this.status ? "turn off" : "turn on"
      }
    }
  }
</script>

<style scoped>
  .sw-box {
    width: 102px;
    margin: auto;
    margin-top: 216px 
  }
</style>